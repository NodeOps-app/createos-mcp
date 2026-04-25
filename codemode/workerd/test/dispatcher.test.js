import { describe, expect, test } from "bun:test";
import dispatcher from "../dispatcher.js";

const MINI_YAML = `openapi: 3.0.0
info:
  title: t
  version: 1.0.0
paths:
  /v1/foo:
    get:
      operationId: listFoo
      responses:
        '200': { description: ok }
`;

function mockEnv() {
  return {
    API: {
      async fetch(_url, init) {
        const body = JSON.parse(init.body);
        if (body.mode === "bootstrap") {
          return new Response(
            JSON.stringify({ status: 200, headers: {}, body: MINI_YAML }),
            { status: 200 },
          );
        }
        return new Response(JSON.stringify({ error: "unhandled" }), { status: 400 });
      },
    },
  };
}

describe("dispatcher", () => {
  test("/health returns ok with spec stats", async () => {
    const req = new Request("http://localhost/health");
    const res = await dispatcher.fetch(req, mockEnv(), {});
    expect(res.status).toBe(200);
    const body = await res.json();
    expect(body.ok).toBe(true);
    expect(body.specVersion).toBe("1.0.0");
    expect(body.endpointCount).toBe(1);
  });

  test("unknown path 404", async () => {
    const req = new Request("http://localhost/whatever");
    const res = await dispatcher.fetch(req, mockEnv(), {});
    expect(res.status).toBe(404);
  });

});
