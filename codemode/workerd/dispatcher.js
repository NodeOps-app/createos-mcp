import { loadSpec } from "./spec-loader.js";
import { buildHostModule, wrapMainModule } from "./host.js";
import { buildSdkModule } from "./api-sdk-builder.js";

let specState = { status: "uninitialized", error: null, data: null };
let specPromise = null;

async function ensureSpec(env) {
  if (specState.status === "ready") return specState.data;
  if (!specPromise) {
    specPromise = (async () => {
      try {
        const data = await loadSpec(env.API);
        specState = { status: "ready", error: null, data };
        return data;
      } catch (e) {
        specState = { status: "failed", error: String(e), data: null };
        specPromise = null;
        throw e;
      }
    })();
  }
  return specPromise;
}

const SEARCH_TIMEOUT_MS = 5_000;
const EXECUTE_SYNC_TIMEOUT_MS = 90_000;

function withTimeout(promise, ms) {
  let to;
  const timer = new Promise((_, rej) => {
    to = setTimeout(() => rej(new Error(`timeout after ${ms}ms`)), ms);
  });
  return Promise.race([promise, timer]).finally(() => clearTimeout(to));
}

function makeApiCallback(apiBinding, authCtx) {
  return async function apiCall({ method, path, query, headers, body }) {
    const r = await apiBinding.fetch("https://internal/proxy", {
      method: "POST",
      headers: { "content-type": "application/json" },
      body: JSON.stringify({
        mode: "authenticated",
        method, path, query, headers, body, authCtx,
      }),
    });
    if (!r.ok) {
      const text = await r.text();
      const err = new Error(`api-proxy ${r.status}: ${text}`);
      err.name = "ProxyError";
      throw err;
    }
    return r.json();
  };
}

async function runIsolate(env, { mode, code, authCtx }) {
  const specData = await ensureSpec(env);
  const isolateId = `iso-${crypto.randomUUID()}`;

  const sdkSrc  = buildSdkModule(specData.spec);
  const hostSrc = buildHostModule({ specRef: specData.spec });
  const mainSrc = wrapMainModule(code);

  const worker = env.LOADER.get(isolateId, () => ({
    compatibilityDate: "2024-09-09",
    mainModule: "main.js",
    modules: {
      "main.js":    mainSrc,
      "host.js":    hostSrc,
      "api-sdk.js": sdkSrc,
    },
  }));

  const entry = worker.getEntrypoint();
  const cb = mode === "execute" ? makeApiCallback(env.API, authCtx) : null;
  return entry.run(cb);
}

export default {
  async fetch(req, env, ctx) {
    const url = new URL(req.url);

    if (url.pathname === "/health") {
      try {
        const data = await ensureSpec(env);
        return Response.json({
          ok: true,
          specVersion: data.version,
          endpointCount: data.endpointCount,
          loadedAt: data.loadedAt,
        });
      } catch (e) {
        return Response.json({ ok: false, error: String(e) }, { status: 503 });
      }
    }

    if (url.pathname === "/run" && req.method === "POST") {
      let body;
      try { body = await req.json(); }
      catch { return Response.json({ status: "error", error: "invalid json" }, { status: 400 }); }

      const { mode, code, authCtx } = body;
      if (mode !== "search" && mode !== "execute") {
        return Response.json({ status: "error", error: "mode must be search or execute" }, { status: 400 });
      }
      if (typeof code !== "string" || code.length === 0) {
        return Response.json({ status: "error", error: "code required" }, { status: 400 });
      }
      if (code.length > 1024 * 1024) {
        return Response.json({ status: "error", error: "code exceeds 1MB" }, { status: 413 });
      }

      const timeout = mode === "search" ? SEARCH_TIMEOUT_MS : EXECUTE_SYNC_TIMEOUT_MS;
      try {
        const result = await withTimeout(runIsolate(env, { mode, code, authCtx }), timeout);
        if (result.ok) {
          return Response.json({ status: "done", result: result.result, logs: result.logs });
        }
        return Response.json({
          status: "error",
          errorKind: result.kind === "ProxyError" ? "infra" : "userCode",
          error: result.error,
          stack: result.stack,
          logs: result.logs,
        });
      } catch (e) {
        const msg = String(e?.message ?? e);
        const errorKind = msg.startsWith("timeout") ? "timeout" : "infra";
        return Response.json({ status: "error", errorKind, error: msg });
      }
    }

    return new Response("not found", { status: 404 });
  },
};
