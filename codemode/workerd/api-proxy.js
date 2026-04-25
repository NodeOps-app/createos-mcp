const BOOTSTRAP_PATH = "/api-docs/openapi.yaml";

export default {
  async fetch(req, env) {
    const body = await req.json();
    const { method, path, headers = {}, query, body: reqBody, mode } = body;

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

    return Response.json(
      { error: "authenticated mode not implemented in phase 1" },
      { status: 501 },
    );
  },
};
