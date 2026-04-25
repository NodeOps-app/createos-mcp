const BOOTSTRAP_PATH = "/api-docs/openapi.yaml";
const PROXY_TIMEOUT_MS = 12_000;

function buildAuthHeaders(authCtx) {
  const headers = {};
  if (!authCtx) return headers;
  if (authCtx.apiKey) headers["X-Api-Key"] = authCtx.apiKey;
  if (authCtx.bearer) headers["Authorization"] = `Bearer ${authCtx.bearer}`;
  return headers;
}

function buildUrl(base, path, query) {
  const u = new URL(path, base);
  if (query && typeof query === "object") {
    for (const [k, v] of Object.entries(query)) {
      if (v == null) continue;
      if (Array.isArray(v)) {
        for (const item of v) u.searchParams.append(k, String(item));
      } else {
        u.searchParams.set(k, String(v));
      }
    }
  }
  return u.toString();
}

export default {
  async fetch(req, env) {
    const body = await req.json();
    const { mode, method, path, headers = {}, query, body: reqBody, authCtx } = body;

    if (mode === "bootstrap") {
      if (method !== "GET" || path !== BOOTSTRAP_PATH) {
        return Response.json(
          { error: "bootstrap mode only allows GET " + BOOTSTRAP_PATH },
          { status: 400 },
        );
      }
      const url = env.BACKEND_URL + path;
      const r = await fetch(url, { method: "GET" });
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

    const url = buildUrl(env.BACKEND_URL, path, query);
    const finalHeaders = {
      ...buildAuthHeaders(authCtx),
      ...headers,
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
      headers: Object.fromEntries(resp.headers),
      body: dataBody,
    });
  },
};
