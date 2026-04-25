import { loadSpec } from "./spec-loader.js";
import { buildHostModule, wrapMainModule } from "./host.js";

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

const SEARCH_TIMEOUT_MS = 5000;

async function runIsolate(env, { mode, code }) {
  const specData = await ensureSpec(env);
  const isolateId = `iso-${crypto.randomUUID()}`;
  const hostSrc = buildHostModule({
    specRef: specData.spec,
    mode,
    apiCallId: null,
  });
  const mainSrc = wrapMainModule(code);

  const worker = env.LOADER.get(isolateId, () => ({
    compatibilityDate: "2024-09-09",
    mainModule: "main.js",
    modules: {
      "main.js": mainSrc,
      "host.js": hostSrc,
    },
  }));
  const entry = worker.getEntrypoint();
  return entry.run(null);
}

function withTimeout(promise, ms) {
  return Promise.race([
    promise,
    new Promise((_, rej) =>
      setTimeout(() => rej(new Error(`timeout after ${ms}ms`)), ms),
    ),
  ]);
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
      try {
        body = await req.json();
      } catch {
        return Response.json({ status: "error", error: "invalid json" }, { status: 400 });
      }
      const { mode, code } = body;
      if (mode !== "search") {
        return Response.json(
          { status: "error", error: "phase 1 supports mode=search only" },
          { status: 501 },
        );
      }
      if (typeof code !== "string" || code.length === 0) {
        return Response.json(
          { status: "error", error: "code must be a non-empty string" },
          { status: 400 },
        );
      }
      if (code.length > 1024 * 1024) {
        return Response.json(
          { status: "error", error: "code exceeds 1MB" },
          { status: 413 },
        );
      }
      try {
        const result = await withTimeout(
          runIsolate(env, { mode, code }),
          SEARCH_TIMEOUT_MS,
        );
        if (result.ok) {
          return Response.json({
            status: "done",
            result: result.result,
            logs: result.logs,
          });
        }
        return Response.json({
          status: "error",
          errorKind: "userCode",
          error: result.error,
          stack: result.stack,
          logs: result.logs,
        });
      } catch (e) {
        return Response.json({
          status: "error",
          errorKind: "infra",
          error: String(e?.message ?? e),
        });
      }
    }

    return new Response("not found", { status: 404 });
  },
};
