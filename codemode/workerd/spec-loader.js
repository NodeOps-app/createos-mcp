import YAML from "yaml";

function resolveRefs(obj, root = obj, seen = new Set()) {
  if (obj == null || typeof obj !== "object") return obj;
  if (Array.isArray(obj)) return obj.map((x) => resolveRefs(x, root, seen));
  if (typeof obj.$ref === "string") {
    if (seen.has(obj.$ref)) return { __circularRef: obj.$ref };
    seen.add(obj.$ref);
    const path = obj.$ref.replace(/^#\//, "").split("/");
    let target = root;
    for (const seg of path) target = target?.[seg];
    const resolved = resolveRefs(target, root, seen);
    seen.delete(obj.$ref);
    return resolved;
  }
  const out = {};
  for (const [k, v] of Object.entries(obj)) out[k] = resolveRefs(v, root, seen);
  return out;
}

function strip(obj) {
  if (obj == null || typeof obj !== "object") return obj;
  if (Array.isArray(obj)) return obj.map(strip);
  const out = {};
  for (const [k, v] of Object.entries(obj)) {
    if (k === "description" || k === "example" || k === "examples") continue;
    out[k] = strip(v);
  }
  return out;
}

export async function loadSpec(apiBinding) {
  const r = await apiBinding.fetch("https://internal/proxy", {
    method: "POST",
    headers: { "content-type": "application/json" },
    body: JSON.stringify({
      mode: "bootstrap",
      method: "GET",
      path: "/api-docs/openapi.yaml",
    }),
  });
  if (!r.ok) throw new Error(`spec bootstrap http ${r.status}`);
  const wrapped = await r.json();
  if (wrapped.status !== 200) {
    throw new Error(`backend returned ${wrapped.status} for openapi.yaml`);
  }
  const parsed = YAML.parse(wrapped.body);
  const resolved = resolveRefs(parsed);
  const trimmed = strip(resolved);

  let endpointCount = 0;
  for (const methods of Object.values(trimmed.paths ?? {})) {
    for (const k of Object.keys(methods)) {
      if (k !== "parameters") endpointCount++;
    }
  }

  return {
    spec: trimmed,
    version: trimmed.info?.version ?? "unknown",
    endpointCount,
    loadedAt: Date.now(),
  };
}
