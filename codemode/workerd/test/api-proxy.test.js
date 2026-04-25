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

  test("authenticated mode injects X-Access-Token (matches native)", async () => {
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
    expect(captured.init.headers["X-Access-Token"]).toBe("tok");
    expect(captured.init.headers["Authorization"]).toBeUndefined();
  });

  test("BACKEND_URL with path prefix is preserved", async () => {
    let captured;
    globalThis.fetch = mock(async (url, init) => {
      captured = { url, init };
      return new Response("[]", { status: 200, headers: { "content-type": "application/json" } });
    });
    const r = await apiProxy.fetch(
      new Request("https://internal/proxy", {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify({
          mode: "authenticated",
          method: "GET",
          path: "/v1/projects",
          authCtx: { apiKey: "k" },
        }),
      }),
      { BACKEND_URL: "https://backend.test/api" },
    );
    expect(captured.url).toBe("https://backend.test/api/v1/projects");
    expect(r.status).toBe(200);
  });

  test("bootstrap honours BACKEND_URL path prefix", async () => {
    let captured;
    globalThis.fetch = mock(async (url) => {
      captured = url;
      return new Response("openapi: 3.0", { status: 200 });
    });
    await apiProxy.fetch(
      new Request("https://internal/proxy", {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify({ mode: "bootstrap", method: "GET", path: "/api-docs/openapi.yaml" }),
      }),
      { BACKEND_URL: "https://backend.test/api" },
    );
    expect(captured).toBe("https://backend.test/api/api-docs/openapi.yaml");
  });

  test("rejects absolute URL in path (SSRF guard)", async () => {
    let called = false;
    globalThis.fetch = mock(async () => { called = true; return new Response("nope"); });
    const r = await call({
      mode: "authenticated",
      method: "GET",
      path: "https://attacker.example/x",
      authCtx: { apiKey: "secret" },
    });
    expect(r.status).toBe(400);
    expect(called).toBe(false);
  });

  test("rejects dot-segment escape ('/../admin')", async () => {
    let called = false;
    globalThis.fetch = mock(async () => { called = true; return new Response("nope"); });
    const r = await apiProxy.fetch(
      new Request("https://internal/proxy", {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify({
          mode: "authenticated",
          method: "GET",
          path: "/../admin",
          authCtx: { apiKey: "k" },
        }),
      }),
      { BACKEND_URL: "https://backend.test/api" },
    );
    expect(r.status).toBe(400);
    expect(called).toBe(false);
  });

  test("rejects percent-encoded dot-segment ('/%2e%2e/admin')", async () => {
    let called = false;
    globalThis.fetch = mock(async () => { called = true; return new Response("nope"); });
    const r = await apiProxy.fetch(
      new Request("https://internal/proxy", {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify({
          mode: "authenticated",
          method: "GET",
          path: "/%2e%2e/admin",
          authCtx: { apiKey: "k" },
        }),
      }),
      { BACKEND_URL: "https://backend.test/api" },
    );
    expect(r.status).toBe(400);
    expect(called).toBe(false);
  });

  test("rejects mixed-case encoded dot-segment ('/%2E%2e/admin')", async () => {
    let called = false;
    globalThis.fetch = mock(async () => { called = true; return new Response("nope"); });
    const r = await apiProxy.fetch(
      new Request("https://internal/proxy", {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify({
          mode: "authenticated",
          method: "GET",
          path: "/%2E%2e/admin",
          authCtx: { apiKey: "k" },
        }),
      }),
      { BACKEND_URL: "https://backend.test/api" },
    );
    expect(r.status).toBe(400);
    expect(called).toBe(false);
  });

  test("rejects partial-encoded dot-segment ('/%2e./admin')", async () => {
    let called = false;
    globalThis.fetch = mock(async () => { called = true; return new Response("nope"); });
    const r = await apiProxy.fetch(
      new Request("https://internal/proxy", {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify({
          mode: "authenticated",
          method: "GET",
          path: "/%2e./admin",
          authCtx: { apiKey: "k" },
        }),
      }),
      { BACKEND_URL: "https://backend.test/api" },
    );
    expect(r.status).toBe(400);
    expect(called).toBe(false);
  });

  test("rejects invalid percent-encoding ('/%ZZ/admin')", async () => {
    let called = false;
    globalThis.fetch = mock(async () => { called = true; return new Response("nope"); });
    const r = await apiProxy.fetch(
      new Request("https://internal/proxy", {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify({
          mode: "authenticated",
          method: "GET",
          path: "/%ZZ/admin",
          authCtx: { apiKey: "k" },
        }),
      }),
      { BACKEND_URL: "https://backend.test/api" },
    );
    expect(r.status).toBe(400);
    expect(called).toBe(false);
  });

  test("rejects dot-segment in middle of path ('/v1/foo/../../admin')", async () => {
    let called = false;
    globalThis.fetch = mock(async () => { called = true; return new Response("nope"); });
    const r = await apiProxy.fetch(
      new Request("https://internal/proxy", {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify({
          mode: "authenticated",
          method: "GET",
          path: "/v1/foo/../../admin",
          authCtx: { apiKey: "k" },
        }),
      }),
      { BACKEND_URL: "https://backend.test/api" },
    );
    expect(r.status).toBe(400);
    expect(called).toBe(false);
  });

  test("invalid JSON body with json content-type returns raw text without throwing", async () => {
    globalThis.fetch = mock(async () => new Response("not json {{{", {
      status: 200,
      headers: { "content-type": "application/json" },
    }));
    const r = await call({ mode: "authenticated", method: "GET", path: "/v1/x", authCtx: { apiKey: "k" } });
    expect(r.status).toBe(200);
    const body = await r.json();
    expect(body.body).toBe("not json {{{");
  });

  test("rejects protocol-relative // path", async () => {
    let called = false;
    globalThis.fetch = mock(async () => { called = true; return new Response("nope"); });
    const r = await call({
      mode: "authenticated",
      method: "GET",
      path: "//attacker.example/x",
      authCtx: { apiKey: "secret" },
    });
    expect(r.status).toBe(400);
    expect(called).toBe(false);
  });

  test("user-injected X-Access-Token / X-Api-Key are stripped (no credential smuggling)", async () => {
    let captured;
    globalThis.fetch = mock(async (_url, init) => {
      captured = init;
      return new Response("[]", { status: 200, headers: { "content-type": "application/json" } });
    });
    await call({
      mode: "authenticated",
      method: "GET",
      path: "/v1/projects",
      authCtx: { apiKey: "real-key" }, // bearer NOT set
      headers: {
        "X-Access-Token": "smuggled-bearer",
        "X-Api-Key":      "smuggled-key",
        "Authorization":  "Bearer smuggled",
        "Cookie":         "session=stolen",
      },
    });
    expect(captured.headers["X-Api-Key"]).toBe("real-key");
    expect(captured.headers["X-Access-Token"]).toBeUndefined();
    expect(captured.headers["Authorization"]).toBeUndefined();
    expect(captured.headers["Cookie"]).toBeUndefined();
  });

  test("user headers cannot override auth", async () => {
    let captured;
    globalThis.fetch = mock(async (_url, init) => {
      captured = init;
      return new Response("[]", { status: 200, headers: { "content-type": "application/json" } });
    });
    await call({
      mode: "authenticated",
      method: "GET",
      path: "/v1/projects",
      authCtx: { apiKey: "real-key" },
      headers: { "X-Api-Key": "spoofed", "Authorization": "Bearer spoofed", "X-Internal-Bypass": "1" },
    });
    expect(captured.headers["X-Api-Key"]).toBe("real-key");
    expect(captured.headers["Authorization"]).toBeUndefined();
    expect(captured.headers["X-Internal-Bypass"]).toBeUndefined();
  });

  test("strips Set-Cookie + internal x-* response headers", async () => {
    globalThis.fetch = mock(async () => new Response("[]", {
      status: 200,
      headers: {
        "content-type": "application/json",
        "set-cookie": "session=secret",
        "x-internal-token": "leak",
        "etag": "abc",
      },
    }));
    const r = await call({ mode: "authenticated", method: "GET", path: "/v1/projects", authCtx: { apiKey: "k" } });
    const body = await r.json();
    expect(body.headers["set-cookie"]).toBeUndefined();
    expect(body.headers["x-internal-token"]).toBeUndefined();
    expect(body.headers["etag"]).toBe("abc");
  });

  test("invalid JSON request body returns 400", async () => {
    const r = await apiProxy.fetch(
      new Request("https://internal/proxy", {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: "not json",
      }),
      env,
    );
    expect(r.status).toBe(400);
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
