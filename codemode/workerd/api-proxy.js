const BOOTSTRAP_PATH = "/api-docs/openapi.yaml";
const PROXY_TIMEOUT_MS = 12_000;

// Backend-bound headers the sandbox is allowed to set. Anything else
// (auth, host, cookie, internal x-* headers) is rejected so user code
// cannot override credentials or smuggle internal state.
const SAFE_REQUEST_HEADERS = new Set([
  "accept",
  "accept-encoding",
  "accept-language",
  "content-type",
  "if-match",
  "if-none-match",
  "if-modified-since",
  "if-unmodified-since",
  "x-request-id",
]);
// Note: x-api-key, x-access-token, authorization, cookie are intentionally
// NOT in the request allowlist. They are credential headers populated only
// by buildAuthHeaders from the trusted authCtx — user code must never be
// able to inject them via the request `headers` field.

// Response headers we relay back to user code. Strips Set-Cookie and any
// x-internal-*, x-vault-*, x-sentinel-* style headers so backend internals
// never leak into the sandbox.
const SAFE_RESPONSE_HEADERS = new Set([
  "content-type",
  "content-length",
  "content-encoding",
  "etag",
  "last-modified",
  "cache-control",
  "date",
  "x-request-id",
]);

// Auth header names match the native Go client (helpers/httpclient.go):
//   api-key    -> "X-Api-Key"
//   bearer     -> "X-Access-Token"  (NOT "Authorization: Bearer" — backend
//                   accepts the X-Access-Token header used by the native
//                   handlers; using Authorization here would silently fail
//                   on Bearer-authenticated callers.)
function buildAuthHeaders(authCtx) {
  const headers = {};
  if (!authCtx) return headers;
  if (authCtx.apiKey) headers["X-Api-Key"] = authCtx.apiKey;
  if (authCtx.bearer) headers["X-Access-Token"] = authCtx.bearer;
  return headers;
}

function sanitizeRequestHeaders(headers) {
  const out = {};
  if (!headers || typeof headers !== "object") return out;
  for (const [k, v] of Object.entries(headers)) {
    if (typeof k !== "string" || typeof v !== "string") continue;
    if (SAFE_REQUEST_HEADERS.has(k.toLowerCase())) out[k] = v;
  }
  return out;
}

function sanitizeResponseHeaders(h) {
  const out = {};
  for (const [k, v] of h) {
    if (SAFE_RESPONSE_HEADERS.has(k.toLowerCase())) out[k] = v;
  }
  return out;
}

// Path validation + base path preservation.
//   - request `path` MUST start with single "/" (rejects "https://..." SSRF
//     and "//attacker" protocol-relative).
//   - if BACKEND_URL has a path prefix (e.g. "https://host/api"), the prefix
//     is preserved and the request path is appended: prefix + path.
//   - bootstrap path ("/api-docs/openapi.yaml") in api-proxy bootstrap mode
//     uses the same prefix-aware concatenation via raw `env.BACKEND_URL +
//     path`, which is fine when BACKEND_URL has no trailing slash.
function buildUrl(base, path, query) {
  if (typeof path !== "string" || !path.startsWith("/") || path.startsWith("//")) {
    throw new Error("path must be a server-relative path starting with '/'");
  }
  const baseUrl = new URL(base);
  // Strip any trailing slash on the base path so we don't end up with "//".
  const basePath = baseUrl.pathname.replace(/\/+$/, "");
  const [reqPath, reqQuery] = path.split("?", 2);
  baseUrl.pathname = basePath + reqPath;
  if (reqQuery !== undefined) baseUrl.search = "?" + reqQuery;
  if (query && typeof query === "object") {
    for (const [k, v] of Object.entries(query)) {
      if (v == null) continue;
      if (Array.isArray(v)) {
        for (const item of v) baseUrl.searchParams.append(k, String(item));
      } else {
        baseUrl.searchParams.set(k, String(v));
      }
    }
  }
  return baseUrl.toString();
}

export default {
  async fetch(req, env) {
    let body;
    try {
      body = await req.json();
    } catch {
      return Response.json({ error: "invalid json body" }, { status: 400 });
    }
    const { mode, method, path, headers = {}, query, body: reqBody, authCtx } = body;

    if (mode === "bootstrap") {
      if (method !== "GET" || path !== BOOTSTRAP_PATH) {
        return Response.json(
          { error: "bootstrap mode only allows GET " + BOOTSTRAP_PATH },
          { status: 400 },
        );
      }
      // Use buildUrl so a BACKEND_URL with a path prefix is honoured.
      let bootUrl;
      try {
        bootUrl = buildUrl(env.BACKEND_URL, path, null);
      } catch (e) {
        return Response.json({ error: String(e?.message ?? e) }, { status: 400 });
      }
      const r = await fetch(bootUrl, { method: "GET" });
      const text = await r.text();
      return Response.json({
        status: r.status,
        headers: Object.fromEntries(r.headers),
        body: text,
      });
    }

    if (mode !== "authenticated") {
      return Response.json({ error: "unknown mode: " + mode }, { status: 400 });
    }

    let url;
    try {
      url = buildUrl(env.BACKEND_URL, path, query);
    } catch (e) {
      return Response.json({ error: String(e?.message ?? e) }, { status: 400 });
    }

    // Auth headers MUST be applied last so sanitised user headers cannot
    // override X-Api-Key / Authorization. SAFE_REQUEST_HEADERS already
    // excludes those, but the spread order is the second line of defence.
    const finalHeaders = {
      ...sanitizeRequestHeaders(headers),
      ...buildAuthHeaders(authCtx),
    };
    if (reqBody != null && !finalHeaders["content-type"] && !finalHeaders["Content-Type"]) {
      finalHeaders["Content-Type"] = "application/json";
    }
    const init = {
      method: method ?? "GET",
      headers: finalHeaders,
    };
    if (reqBody != null) {
      init.body = typeof reqBody === "string" ? reqBody : JSON.stringify(reqBody);
    }

    const ac = new AbortController();
    const t = setTimeout(() => ac.abort(), PROXY_TIMEOUT_MS);
    init.signal = ac.signal;

    let resp;
    try {
      resp = await fetch(url, init);
    } catch (e) {
      clearTimeout(t);
      const isAbort = e?.name === "AbortError";
      return Response.json({
        kind: isAbort ? "timeout" : "network",
        error: String(e?.message ?? e),
      }, { status: 502 });
    }
    clearTimeout(t);

    const contentType = resp.headers.get("content-type") ?? "";
    let dataBody;
    if (contentType.includes("application/json")) {
      try {
        dataBody = await resp.json();
      } catch {
        dataBody = await resp.text();
      }
    } else {
      dataBody = await resp.text();
    }
    return Response.json({
      status: resp.status,
      headers: sanitizeResponseHeaders(resp.headers),
      body: dataBody,
    });
  },
};
