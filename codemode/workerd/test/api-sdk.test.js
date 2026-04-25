import { describe, expect, test } from "bun:test";
import { buildSdkModule } from "../api-sdk-builder.js";

const SPEC = {
  paths: {
    "/v1/projects": {
      get: {
        operationId: "listProjects",
        tags: ["projects"],
        parameters: [{ in: "query", name: "limit", required: false }],
      },
      post: {
        operationId: "createProject",
        tags: ["projects"],
        requestBody: { content: { "application/json": { schema: { type: "object" } } } },
      },
    },
    "/v1/projects/{id}": {
      get: {
        operationId: "getProject",
        tags: ["projects"],
        parameters: [{ in: "path", name: "id", required: true }],
      },
    },
  },
};

describe("buildSdkModule", () => {
  test("generates buildApi with grouped operations", () => {
    const src = buildSdkModule(SPEC);
    expect(src).toContain("export function buildApi(callback)");
    expect(src).toContain('"projects"');
    expect(src).toContain('"listProjects"');
    expect(src).toContain('"createProject"');
    expect(src).toContain('"getProject"');
    expect(src).toContain("raw:");
  });

  test("operation closure preserves path + method", () => {
    const src = buildSdkModule(SPEC);
    expect(src).toContain('"path":"/v1/projects/{id}"');
    expect(src).toContain('"method":"GET"');
  });
});
