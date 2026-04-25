// Builds the source for a per-call `host.js` ESM module.
// Exports spec/console/sleep. `api` is constructed in main.js from the
// callback installed via setApi() RPC.
export function buildHostModule({ specRef }) {
  const parts = [];
  parts.push(`const spec = JSON.parse(${JSON.stringify(JSON.stringify(specRef))});`);
  parts.push(`const __logs = [];`);
  parts.push(`function __push(level, args) {
    if (__logs.length > 1000) return;
    let msg = args.map((a) => {
      try { return typeof a === "string" ? a : JSON.stringify(a); }
      catch { return String(a); }
    }).join(" ");
    if (msg.length > 4000) msg = msg.slice(0, 4000) + "…";
    __logs.push({ level, msg, t: Date.now() });
  }`);
  parts.push(`const console = {
    log: (...a) => __push("log", a),
    info: (...a) => __push("info", a),
    warn: (...a) => __push("warn", a),
    error: (...a) => __push("error", a),
  };`);
  parts.push(`async function sleep(ms) {
    if (typeof ms !== "number" || ms < 0) throw new Error("sleep: ms must be non-negative");
    if (ms > 60000) ms = 60000;
    return new Promise((r) => setTimeout(r, ms));
  }`);
  parts.push(`export { spec, console, sleep };`);
  parts.push(`export function __getLogs() { return __logs; }`);
  return parts.join("\n");
}

export function wrapMainModule(userCode) {
  // userCode is interpolated into the module body. Note: a syntax error in
  // userCode causes the dynamic ESM module compile to fail BEFORE run() is
  // ever called, so it surfaces from the dispatcher as kind=infra (the
  // worker loader factory throws). Only runtime errors thrown from inside
  // userFn are caught by the try below and returned as kind=userCode.
  // The dispatcher distinguishes them by inspecting the error name
  // ("SyntaxError" → reclassify to userCode) before responding.
  return `
import { spec, console, sleep, __getLogs } from "host.js";
import { buildApi } from "api-sdk.js";
export default {
  async run(callback) {
    const api = callback ? buildApi(callback) : undefined;
    try {
      const userFn = (${userCode});
      const result = await userFn();
      return { ok: true, result, logs: __getLogs() };
    } catch (e) {
      return {
        ok: false,
        error: String(e?.message ?? e),
        kind: e?.name ?? "Error",
        stack: e?.stack ?? null,
        logs: __getLogs(),
      };
    }
  },
};
`;
}
