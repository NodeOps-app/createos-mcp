#!/usr/bin/env bun
import { encoding_for_model } from "tiktoken";

const SEARCH_DESC = `Code Mode: introspect the CreateOS OpenAPI spec from a sandboxed JS arrow fn.

When: discover endpoints/params/schemas BEFORE calling execute. No network in this mode.

Globals: spec ($refs resolved, descriptions stripped), console, sleep(ms).
Result envelope: {status:"done"|"error", result, logs} | {status:"error", errorKind:"userCode"|"timeout"|"infra", error, stack}.
Limits: 5s timeout, 1MB code, 64KB result.

See resource code-mode://intro and prompt code-mode/api-discovery.`;

const EXECUTE_DESC = `Code Mode: run JS against the CreateOS API in a fresh V8 isolate. Auth forwarded from caller.

API shape:
  api.<group>.<operationId>(args)   typed call; group = first OpenAPI tag, operationId from spec.
  api.raw(method, path, {query, headers, body})  escape hatch for unmapped routes.
  Path params come from args by name: api.projects.getProject({id:"..."}).
  Query/body inferred from operation; pass {body:{...}} explicitly to override.

Errors: non-2xx throws ApiError {name, status, body, path}. Catch in user code if you want to handle them.

Result envelope:
  sync     -> {status:"done", result, logs}
  >90s     -> {status:"running", jobId} -> loop pollJob(jobId)
  failure  -> {status:"error", errorKind:"userCode"|"timeout"|"infra"|"capacity", error, stack, logs}

Limits: 90s sync window, 600s wall cap, 1MB code, 64KB result, 50 concurrent jobs.
Sandbox: no fs, no env, no ambient fetch. Only api/console/sleep/spec.

See resource code-mode://api-shape and prompt code-mode/deploy-example.`;

const POLLJOB_DESC = `Code Mode: drain a long-running execute() job.

When: a prior execute() returned {status:"running", jobId}. Loop until status != "running".
Blocks up to ~90s per call (long-poll).

Result envelope: {status:"running", jobId, logsSoFar} | {status:"done", result, logs} | {status:"error", errorKind:"userCode"|"timeout"|"infra"|"jobMissing", error, stack}.
Wall cap on the underlying job: 600s. Jobs evicted 30 min after terminal.`;

const TOOLS: { name: string; description: string; inputSchema: object }[] = [
  { name: "GetQuotas", description: "Get the user's quota limits.", inputSchema: { type: "object", properties: {} } },
  { name: "GetSupportedProjectTypes", description: "List project types the platform supports.", inputSchema: { type: "object", properties: {} } },
  { name: "CheckProjectUniqueName", description: "Check whether a project unique name is available.", inputSchema: { type: "object", properties: { uniqueName: { type: "string" } }, required: ["uniqueName"] } },
  { name: "CreateProject", description: "Create a new project.", inputSchema: { type: "object", properties: { name: { type: "string" }, projectType: { type: "string" } }, required: ["name", "projectType"] } },
  { name: "UploadDeploymentBase64Files", description: "Upload base64-encoded files for a deployment.", inputSchema: { type: "object", properties: { projectId: { type: "string" }, files: { type: "array" } }, required: ["projectId", "files"] } },
  { name: "GetDeployment", description: "Get a deployment by id.", inputSchema: { type: "object", properties: { id: { type: "string" } }, required: ["id"] } },
  { name: "CancelDeployment", description: "Cancel an in-progress deployment.", inputSchema: { type: "object", properties: { id: { type: "string" } }, required: ["id"] } },

  { name: "search", description: SEARCH_DESC, inputSchema: { type: "object", properties: { code: { type: "string" } }, required: ["code"] } },
  { name: "execute", description: EXECUTE_DESC, inputSchema: { type: "object", properties: { code: { type: "string" } }, required: ["code"] } },
  { name: "pollJob", description: POLLJOB_DESC, inputSchema: { type: "object", properties: { jobId: { type: "string" } }, required: ["jobId"] } },
];

const enc = encoding_for_model("gpt-4");
let total = 0;
for (const tool of TOOLS) {
  const blob = JSON.stringify(tool);
  const tokens = enc.encode(blob).length;
  total += tokens;
  console.log(`${tool.name}: ${tokens}`);
}
enc.free();

console.log(`---`);
console.log(`total: ${total}`);

const BUDGET = 1500;
if (total > BUDGET) {
  console.error(`token budget exceeded: ${total} > ${BUDGET}`);
  process.exit(1);
}
console.log(`within budget (${BUDGET})`);
