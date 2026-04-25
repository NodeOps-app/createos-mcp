#!/usr/bin/env bun
import { encoding_for_model } from "tiktoken";

const TOOLS: { name: string; description: string; inputSchema: object }[] = [
  { name: "GetQuotas", description: "Get the user's quota limits.", inputSchema: { type: "object", properties: {} } },
  { name: "GetSupportedProjectTypes", description: "List project types the platform supports.", inputSchema: { type: "object", properties: {} } },
  { name: "CheckProjectUniqueName", description: "Check whether a project unique name is available.", inputSchema: { type: "object", properties: { uniqueName: { type: "string" } }, required: ["uniqueName"] } },
  { name: "CreateProject", description: "Create a new project.", inputSchema: { type: "object", properties: { name: { type: "string" }, projectType: { type: "string" } }, required: ["name", "projectType"] } },
  { name: "UploadDeploymentBase64Files", description: "Upload base64-encoded files for a deployment.", inputSchema: { type: "object", properties: { projectId: { type: "string" }, files: { type: "array" } }, required: ["projectId", "files"] } },
  { name: "GetDeployment", description: "Get a deployment by id.", inputSchema: { type: "object", properties: { id: { type: "string" } }, required: ["id"] } },
  { name: "CancelDeployment", description: "Cancel an in-progress deployment.", inputSchema: { type: "object", properties: { id: { type: "string" } }, required: ["id"] } },

  { name: "search", description: "Search the CreateOS OpenAPI spec. Write a JS async arrow function that reads `spec` and returns matches. No network access in this mode. Use to discover endpoints, parameters, and shapes before calling execute().", inputSchema: { type: "object", properties: { code: { type: "string" } }, required: ["code"] } },
  { name: "execute", description: "Execute JavaScript against the CreateOS API. Single sandbox call may chain multiple api calls. Returns either {status:'done', result, logs} or {status:'running', jobId} (long-running ops, polled via pollJob).", inputSchema: { type: "object", properties: { code: { type: "string" } }, required: ["code"] } },
  { name: "pollJob", description: "Poll a long-running execute() job. Blocks up to ~90s; returns {status: 'running'|'done'|'error', ...}. Loop until status != 'running'.", inputSchema: { type: "object", properties: { jobId: { type: "string" } }, required: ["jobId"] } },
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
