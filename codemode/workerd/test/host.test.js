import { describe, expect, test } from "bun:test";
import { buildHostModule, wrapMainModule } from "../host.js";

describe("buildHostModule", () => {
  test("exports spec, console, sleep", () => {
    const src = buildHostModule({ specRef: { paths: {} } });
    expect(src).toContain("export { spec, console, sleep }");
  });

  test("includes sleep cap", () => {
    const src = buildHostModule({ specRef: {} });
    expect(src).toContain("60000");
  });
});

describe("wrapMainModule", () => {
  test("emits run entrypoint that builds api from callback", () => {
    const src = wrapMainModule("async () => 42");
    expect(src).toContain("async run(callback)");
    expect(src).toContain("buildApi");
  });
});
