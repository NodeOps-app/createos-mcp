import { describe, expect, test } from "bun:test";
import fs from "node:fs";
import { loadSpec } from "../spec-loader.js";

function makeBinding(yamlText) {
  return {
    async fetch(_url, init) {
      const body = JSON.parse(init.body);
      if (body.mode !== "bootstrap") throw new Error("expected bootstrap");
      return new Response(
        JSON.stringify({ status: 200, headers: {}, body: yamlText }),
        { status: 200 },
      );
    },
  };
}

describe("spec-loader", () => {
  test("resolves $refs and strips descriptions", async () => {
    const yamlText = fs.readFileSync(
      new URL("./fixtures/openapi-mini.yaml", import.meta.url),
      "utf8",
    );
    const result = await loadSpec(makeBinding(yamlText));
    expect(result.endpointCount).toBe(1);
    expect(result.version).toBe("1.0.0");
    const op = result.spec.paths["/v1/foo"].get;
    expect(op.operationId).toBe("listFoo");
    expect(op.description).toBeUndefined();
    expect(op.responses["200"].$ref).toBeUndefined();
    expect(op.responses["200"].content["application/json"].schema.type).toBe(
      "object",
    );
  });

  test("non-200 backend response throws", async () => {
    const binding = {
      async fetch() {
        return new Response(
          JSON.stringify({ status: 500, headers: {}, body: "oops" }),
          { status: 200 },
        );
      },
    };
    await expect(loadSpec(binding)).rejects.toThrow(/500/);
  });
});
