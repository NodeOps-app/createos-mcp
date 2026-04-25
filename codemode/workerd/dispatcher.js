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
const SYNC_WINDOW_MS    = 90_000;
const WALL_CAP_MS       = 600_000;
const JOB_TTL_MS        = 30 * 60_000;
const RESULT_CAP_BYTES  = 64 * 1024;

// DoS guards. Bound the number of in-flight isolates and the size of the
// jobStore so a flood of long-running executes (or of stuck CPU-bound code
// that workerd cannot synchronously terminate) cannot exhaust memory or
// pin every worker. /run returns 429 when a limit is hit.
const MAX_RUNNING_JOBS    = 50;   // concurrent execute jobs
const MAX_JOB_STORE_TOTAL = 500;  // running + finished-but-not-yet-evicted
const MAX_RUNNING_SEARCHES = 50;  // concurrent search isolates

const jobStore = new Map();
let runningSearches = 0;

function nowMs() { return Date.now(); }

function evictOldJobs() {
  const t = nowMs();
  for (const [id, j] of jobStore.entries()) {
    if (j.finishedAt && t - j.finishedAt > JOB_TTL_MS) jobStore.delete(id);
  }
}

function countRunningJobs() {
  let n = 0;
  for (const j of jobStore.values()) if (j.status === "running") n++;
  return n;
}

// Evict the oldest finished job to make room. Returns true if anything was
// evicted; false if every job is currently running (caller should 429).
function evictOldestFinished() {
  let oldestId = null;
  let oldestFinished = Infinity;
  for (const [id, j] of jobStore.entries()) {
    if (j.finishedAt && j.finishedAt < oldestFinished) {
      oldestFinished = j.finishedAt;
      oldestId = id;
    }
  }
  if (oldestId !== null) {
    jobStore.delete(oldestId);
    return true;
  }
  return false;
}

function notifyWaiters(job) {
  for (const w of job.waiters) w();
  job.waiters.clear();
}

function maybeTruncateResult(result) {
  try {
    const s = JSON.stringify(result);
    if (s.length > RESULT_CAP_BYTES) {
      return { __truncated: true, preview: s.slice(0, RESULT_CAP_BYTES) };
    }
    return result;
  } catch (e) {
    return { __unserialisable: String(e) };
  }
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

// Attempt best-effort isolate teardown on wall-timeout. workerd does not
// expose a synchronous isolate.terminate() API, so the V8 isolate continues
// to run user code until the Worker Loader's LRU evicts it. Disposing the
// entry stub releases the RPC connection so subsequent api-proxy calls from
// the runaway code surface as ProxyError instead of being silently delivered.
function bestEffortDispose(stub) {
  try {
    if (stub && typeof stub[Symbol.dispose] === "function") stub[Symbol.dispose]();
  } catch { /* noop */ }
}

async function startIsolateRun(env, { mode, code, authCtx, timeoutMs }) {
  const specData = await ensureSpec(env);
  const isolateId = `iso-${crypto.randomUUID()}`;

  const sdkSrc  = buildSdkModule(specData.spec);
  const hostSrc = buildHostModule({ specRef: specData.spec });
  const mainSrc = wrapMainModule(code);

  const worker = env.LOADER.get(isolateId, () => ({
    compatibilityDate: "2024-09-09",
    mainModule: "main.js",
    modules: { "main.js": mainSrc, "host.js": hostSrc, "api-sdk.js": sdkSrc },
  }));

  const entry = worker.getEntrypoint();
  const cb = mode === "execute" ? makeApiCallback(env.API, authCtx) : null;

  const cap = typeof timeoutMs === "number" ? timeoutMs : WALL_CAP_MS;
  let timerHandle;
  const timer = new Promise((_, rej) => {
    timerHandle = setTimeout(
      () => rej(new Error(mode === "search" ? "search_timeout" : "wall_timeout")),
      cap,
    );
  });

  try {
    return await Promise.race([entry.run(cb), timer]);
  } catch (e) {
    // SyntaxError thrown by the dynamic ESM compile (because userCode
    // failed to parse) is reclassified as a userCode error so the model
    // sees actionable feedback rather than an opaque infra failure.
    if (e && (e.name === "SyntaxError" || /SyntaxError/.test(String(e?.message ?? "")))) {
      return {
        ok: false,
        error: String(e?.message ?? e),
        kind: "SyntaxError",
        stack: e?.stack ?? null,
        logs: [],
      };
    }
    throw e;
  } finally {
    clearTimeout(timerHandle);
    // Always dispose entry+worker. On success this releases the cached
    // worker promptly so /run does not accumulate isolates faster than
    // the Worker Loader's LRU can evict them. On timeout this also
    // severs the api-proxy callback so runaway code's further calls
    // surface as ProxyError instead of being silently delivered.
    bestEffortDispose(entry);
    bestEffortDispose(worker);
  }
}

function newJob() {
  return {
    status: "running",
    result: null,
    error: null,
    errorKind: null,
    stack: null,
    logs: [],
    createdAt: nowMs(),
    finishedAt: null,
    waiters: new Set(),
  };
}

function finaliseJob(job, isolateOutcome) {
  if (isolateOutcome.kind === "done") {
    job.status = "done";
    job.result = maybeTruncateResult(isolateOutcome.result);
    job.logs = isolateOutcome.logs ?? [];
  } else if (isolateOutcome.kind === "userError") {
    job.status = "error";
    job.errorKind = "userCode";
    job.error = isolateOutcome.error;
    job.stack = isolateOutcome.stack;
    job.logs = isolateOutcome.logs ?? [];
  } else if (isolateOutcome.kind === "wallTimeout") {
    job.status = "error";
    job.errorKind = "timeout";
    job.error = "wall timeout (600s)";
  } else {
    job.status = "error";
    job.errorKind = "infra";
    job.error = isolateOutcome.error ?? "unknown";
  }
  job.finishedAt = nowMs();
  notifyWaiters(job);
}

function jobResponse(jobId, job) {
  if (job.status === "running") {
    return { status: "running", jobId, logsSoFar: job.logs.slice(-50) };
  }
  if (job.status === "done") {
    return { status: "done", result: job.result, logs: job.logs };
  }
  return {
    status: "error",
    errorKind: job.errorKind,
    error: job.error,
    stack: job.stack,
    logs: job.logs,
  };
}

async function awaitJobUntil(jobId, deadlineMs) {
  const job = jobStore.get(jobId);
  if (!job) return null;
  if (job.status !== "running") return job;
  return new Promise((resolve) => {
    let to;
    const cleanup = () => { job.waiters.delete(wake); clearTimeout(to); };
    const wake = () => { cleanup(); resolve(job); };
    job.waiters.add(wake);
    const remaining = Math.max(0, deadlineMs - nowMs());
    to = setTimeout(() => { cleanup(); resolve(job); }, remaining);
  });
}

export default {
  async fetch(req, env, ctx) {
    evictOldJobs();
    const url = new URL(req.url);

    if (url.pathname === "/health") {
      try {
        const data = await ensureSpec(env);
        return Response.json({
          ok: true,
          specVersion: data.version,
          endpointCount: data.endpointCount,
          loadedAt: data.loadedAt,
          jobsRunning: [...jobStore.values()].filter((j) => j.status === "running").length,
          jobStoreSize: jobStore.size,
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

      // Search: simple sync timeout (5s). Timeout disposes the isolate.
      if (mode === "search") {
        if (runningSearches >= MAX_RUNNING_SEARCHES) {
          return Response.json(
            { status: "error", errorKind: "capacity", error: "code mode at search capacity" },
            { status: 429 },
          );
        }
        runningSearches++;
        try {
          const out = await startIsolateRun(env, { mode, code, authCtx, timeoutMs: SEARCH_TIMEOUT_MS });
          if (out.ok) {
            return Response.json({
              status: "done",
              result: maybeTruncateResult(out.result),
              logs: out.logs,
            });
          }
          return Response.json({
            status: "error",
            errorKind: "userCode",
            error: out.error,
            stack: out.stack,
            logs: out.logs,
          });
        } catch (e) {
          const msg = String(e?.message ?? e);
          return Response.json({
            status: "error",
            errorKind: msg.includes("timeout") ? "timeout" : "infra",
            error: msg,
          });
        } finally {
          runningSearches--;
        }
      }

      // Execute: enforce concurrency + total caps before registering a job.
      // Running cap is hard. Total cap evicts the oldest finished job
      // (LRU on finishedAt) before falling back to 429 — this prevents a
      // burst of fast successful executes from filling the store and
      // 429-ing subsequent requests for the next 30 min.
      const runningJobs = countRunningJobs();
      if (runningJobs >= MAX_RUNNING_JOBS) {
        return Response.json(
          { status: "error", errorKind: "capacity", error: "too many running jobs" },
          { status: 429 },
        );
      }
      if (jobStore.size >= MAX_JOB_STORE_TOTAL) {
        if (!evictOldestFinished()) {
          return Response.json(
            { status: "error", errorKind: "capacity", error: "job store full (all running)" },
            { status: 429 },
          );
        }
      }

      const jobId = `j_${crypto.randomUUID()}`;
      const job = newJob();
      jobStore.set(jobId, job);

      const runPromise = (async () => {
        try {
          const out = await startIsolateRun(env, { mode, code, authCtx, timeoutMs: WALL_CAP_MS });
          if (out.ok) finaliseJob(job, { kind: "done", result: out.result, logs: out.logs });
          else finaliseJob(job, { kind: "userError", error: out.error, stack: out.stack, logs: out.logs });
        } catch (e) {
          if (String(e?.message) === "wall_timeout") {
            finaliseJob(job, { kind: "wallTimeout" });
          } else {
            finaliseJob(job, { kind: "infra", error: String(e?.message ?? e) });
          }
        }
      })();

      if (ctx?.waitUntil) ctx.waitUntil(runPromise);

      const deadline = nowMs() + SYNC_WINDOW_MS;
      const finished = await awaitJobUntil(jobId, deadline);
      return Response.json(jobResponse(jobId, finished ?? job));
    }

    if (url.pathname === "/poll" && req.method === "POST") {
      let body;
      try { body = await req.json(); }
      catch { return Response.json({ status: "error", error: "invalid json" }, { status: 400 }); }
      const jobId = body.jobId;
      if (!jobId || typeof jobId !== "string") {
        return Response.json({ status: "error", errorKind: "userCode", error: "jobId required" }, { status: 400 });
      }
      const job = jobStore.get(jobId);
      if (!job) {
        return Response.json({ status: "error", errorKind: "jobMissing", error: "unknown jobId" });
      }
      const deadline = nowMs() + SYNC_WINDOW_MS;
      const finished = await awaitJobUntil(jobId, deadline);
      return Response.json(jobResponse(jobId, finished ?? job));
    }

    return new Response("not found", { status: 404 });
  },
};
