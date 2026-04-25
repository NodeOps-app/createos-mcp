import { describe, expect, test, mock } from "bun:test";
import apiProxy from "../api-proxy.js";

const env = { BACKEND_URL: "https://backend.test" };

function call(payload) {
  return apiProxy.fetch(
    new Request("https://internal/proxy", {
      method: "POST",
      headers: { "content-type": "application/json" },
      body: JSON.stringify(payload),
    }),
    env,
  );
}

describe("api-proxy", () => {
  test("rejects unknown mode", async () => {
    const r = await call({ mode: "weird", method: "GET", path: "/x" });
    expect(r.status).toBe(400);
  });

  test("authenticated mode injects X-Api-Key", async () => {
    let captured;
    globalThis.fetch = mock(async (url, init) => {
      captured = { url, init };
      return new Response(JSON.stringify({ items: [] }), {
        status: 200,
        headers: { "content-type": "application/json" },
      });
    });
    const r = await call({
      mode: "authenticated",
      method: "GET",
      path: "/v1/projects",
      authCtx: { apiKey: "k1" },
    });
    expect(r.status).toBe(200);
    const body = await r.json();
    expect(body.status).toBe(200);
    expect(captured.init.headers["X-Api-Key"]).toBe("k1");
  });

  test("authenticated mode injects Bearer", async () => {
    let captured;
    globalThis.fetch = mock(async (url, init) => {
      captured = { url, init };
      return new Response("[]", {
        status: 200,
        headers: { "content-type": "application/json" },
      });
    });
    await call({
      mode: "authenticated",
      method: "GET",
      path: "/v1/projects",
      authCtx: { bearer: "tok" },
    });
    expect(captured.init.headers["Authorization"]).toBe("Bearer tok");
  });

  test("query string assembled", async () => {
    let captured;
    globalThis.fetch = mock(async (url, init) => {
      captured = { url, init };
      return new Response("[]", {
        status: 200,
        headers: { "content-type": "application/json" },
      });
    });
    await call({
      mode: "authenticated",
      method: "GET",
      path: "/v1/projects",
      query: { limit: 5, tag: ["a", "b"] },
    });
    expect(captured.url).toBe("https://backend.test/v1/projects?limit=5&tag=a&tag=b");
  });
});
