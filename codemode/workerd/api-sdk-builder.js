// Mirrors scripts/gen-sdk.ts but operates on in-memory spec.
// Returns an ESM source string defining buildApi(callback) → typed proxy
// keyed by tag.operationId, with `raw(method, path, opts)` escape hatch.

function collectOps(spec) {
  const ops = [];
  for (const [path, methods] of Object.entries(spec.paths ?? {})) {
    for (const [method, raw] of Object.entries(methods)) {
      if (method === "parameters" || typeof raw !== "object" || raw === null) continue;
      const op = raw;
      if (!op.operationId) continue;
      const group = (op.tags?.[0] ?? "default").replace(/[^a-zA-Z0-9_]/g, "");
      ops.push({
        operationId: op.operationId,
        method: method.toUpperCase(),
        path,
        group,
        paramsIn: (op.parameters ?? []).map((p) => ({
          in: p.in, name: p.name, required: !!p.required,
        })),
        hasBody: !!op.requestBody,
      });
    }
  }
  return ops;
}

export function buildSdkModule(spec) {
  const ops = collectOps(spec);
  const groups = {};
  for (const op of ops) (groups[op.group] ??= []).push(op);

  const lines = [];
  lines.push(`function __operation(callback, op) {
    return async function (args = {}) {
      const path = op.path.replace(/\\{([^}]+)\\}/g, (_, name) => {
        const v = args[name];
        if (v == null) throw new Error("missing path param: " + name);
        return encodeURIComponent(String(v));
      });
      const query = {};
      let body;
      for (const p of op.paramsIn) {
        if (p.in === "query" && args[p.name] !== undefined) query[p.name] = args[p.name];
      }
      if (op.hasBody && args.body !== undefined) body = args.body;
      else if (op.hasBody) {
        const inferred = {};
        for (const [k, v] of Object.entries(args)) {
          if (op.paramsIn.some((p) => p.name === k)) continue;
          if (k === "body") continue;
          inferred[k] = v;
        }
        body = inferred;
      }
      const res = await callback({ method: op.method, path, query, body });
      if (res.status >= 400) {
        const err = new Error("ApiError " + res.status + " " + op.method + " " + path);
        err.name = "ApiError";
        err.status = res.status;
        err.body = res.body;
        err.path = path;
        throw err;
      }
      return { status: res.status, headers: res.headers, data: res.body };
    };
  }`);
  lines.push("");
  lines.push("export function buildApi(callback) {");
  lines.push("  return {");
  for (const [group, list] of Object.entries(groups)) {
    lines.push(`    ${JSON.stringify(group)}: {`);
    for (const op of list) {
      const meta = JSON.stringify({
        method: op.method, path: op.path, paramsIn: op.paramsIn, hasBody: op.hasBody,
      });
      lines.push(`      ${JSON.stringify(op.operationId)}: __operation(callback, ${meta}),`);
    }
    lines.push("    },");
  }
  lines.push("    raw: (method, path, opts = {}) => callback({ method, path, ...opts }),");
  lines.push("  };");
  lines.push("}");
  return lines.join("\n");
}
