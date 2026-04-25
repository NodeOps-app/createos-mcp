package codemode

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

// MCP Resources + Prompts that expose Code Mode usage to clients.
// Tools/list alone tells the model the names; these give it the
// patterns. Pulled on-demand — zero token cost until requested.

const introContent = `# CreateOS MCP — Code Mode v2

This server exposes 10 tools:

  Native fast-path (7):
    GetQuotas, GetSupportedProjectTypes, CheckProjectUniqueName,
    CreateProject, UploadDeploymentBase64Files, GetDeployment,
    CancelDeployment

  Code Mode (3):
    search(code)   — read-only spec introspection in a JS sandbox
    execute(code)  — typed CreateOS API calls in a JS sandbox
    pollJob(jobId) — drain long-running execute jobs

The full ~100-endpoint API is reachable via execute. Use search first
to discover an endpoint, then execute to call it.

## Workflow

1. search: 'async () => Object.keys(spec.paths).filter(p => p.includes("project"))'
2. search: 'async () => spec.paths["/v1/projects"]'   // inspect schema
3. execute: 'async () => (await api.projects.listProjects({limit:5})).data'

If execute exceeds 90s it returns {status:"running", jobId}; loop
pollJob(jobId) until status != "running".

## Auth

Provide X-Api-Key or Authorization: Bearer <token> on the MCP request.
For stdio transport, set CREATEOS_API_KEY or CREATEOS_BEARER env vars.

## Sandbox

workerd V8 isolate. No fs, no env, no ambient fetch. Only:
  spec     — OpenAPI document, $refs resolved, descriptions stripped
  api      — typed proxy (execute mode only)
  console  — log/info/warn/error captured into result.logs
  sleep    — async, capped at 60s per call

## Limits

  search:  5s timeout
  execute: 90s sync window, 600s wall cap, 50 concurrent
  jobs:    30 min TTL after terminal, 500 total cap
  code:    1 MB
  result:  64 KB (oversized truncated with __truncated:true)
`

const apiShapeContent = `# Code Mode — api shape

## Typed accessor

  api.<group>.<operationId>(args)

  group        = first OpenAPI tag for the operation (lowercase, alnum_).
  operationId  = exact operationId from the spec.
  args         = single object. Path params by name; query params by name;
                 body inferred from remaining keys, or pass {body:{...}}
                 explicitly.

  Example:
    api.projects.getProject({id:"p_abc"})
    api.deployments.createDeployment({projectId:"p_abc", body:{...}})

## Escape hatch

  api.raw(method, path, {query, headers, body})

  Use when an operation lacks an operationId or you need a custom call.
  Path MUST be server-relative (starts with "/"); dot-segments rejected.

## Return shape

  Success: { status, headers, data }   data = parsed JSON or raw text.
  Failure: throws ApiError {
    name:    "ApiError",
    status:  number,           // backend HTTP status
    body:    parsed | string,
    path:    request path,
  }

  Catch in user code if you want to inspect failures without rejecting:

    try { return (await api.foo.listFoo()).data; }
    catch (e) { return { kind: e.name, status: e.status, body: e.body }; }

## Discovering operationIds

  search: 'async () => Object.entries(spec.paths).flatMap(([p,m]) =>
    Object.entries(m).filter(([k]) => k !== "parameters")
      .map(([method, op]) => ({ path:p, method, op:op.operationId, tags:op.tags }))
  )'

## Forbidden

  - Dynamic code evaluation primitives are not exposed; user code is
    compiled to a static ESM module.
  - fetch — only api/api.raw reach the backend; no other network.
  - Set-Cookie / Authorization / X-Api-Key in api.raw headers — stripped.
`

func newCodeModeResource(uri, name, content string) (mcp.Resource, ResourceHandlerFunc) {
	r := mcp.NewResource(
		uri, name,
		mcp.WithMIMEType("text/markdown"),
		mcp.WithResourceDescription("Code Mode reference. Pull on demand."),
	)
	body := content
	handler := func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      uri,
				MIMEType: "text/markdown",
				Text:     body,
			},
		}, nil
	}
	return r, handler
}

// ResourceHandlerFunc mirrors server.ResourceHandlerFunc to avoid an import
// cycle on the server package; identical signature.
type ResourceHandlerFunc func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error)

// NewIntroResource exposes the high-level Code Mode usage doc.
func NewIntroResource() (mcp.Resource, ResourceHandlerFunc) {
	return newCodeModeResource("code-mode://intro", "Code Mode intro + workflow", introContent)
}

// NewAPIShapeResource exposes the api proxy / ApiError contract.
func NewAPIShapeResource() (mcp.Resource, ResourceHandlerFunc) {
	return newCodeModeResource("code-mode://api-shape", "Code Mode api shape + ApiError contract", apiShapeContent)
}

// Prompts ---------------------------------------------------------------

const deployExamplePromptText = `You are using CreateOS MCP Code Mode v2.
Goal: deploy a project from an upload and wait for the deployment URL.

Pattern (single execute call):

  execute({code: 'async () => {
    const { data: dep } = await api.deployments.createDeployment({
      projectId: PROJECT_ID,
      body: { source: "upload" },
    });
    for (let i = 0; i < 20; i++) {
      await sleep(30000);
      const { data: s } = await api.deployments.getDeployment({ id: dep.id });
      if (s.status === "deployed") return s.url;
      if (s.status === "failed")   throw new Error(s.error ?? "deploy failed");
    }
    throw new Error("deploy did not complete in 10 min");
  }'})

If the call exceeds 90s, the result is {status:"running", jobId}. Loop:

  pollJob({jobId: "j_..."})

until status != "running". Each pollJob blocks up to ~90s.

Replace PROJECT_ID with the real value from CreateProject (native tool)
or api.projects.listProjects().`

const apiDiscoveryPromptText = `You are using CreateOS MCP Code Mode v2.
Goal: find which endpoints + parameters are relevant to a task.

Use search (no network, fast):

  // 1. Find paths matching a keyword
  search({code: 'async () => Object.keys(spec.paths).filter(p => /domain/i.test(p))'})

  // 2. Inspect the operations on a path
  search({code: 'async () => spec.paths["/v1/domains"]'})

  // 3. Get the input schema of a specific operation
  search({code: 'async () => spec.paths["/v1/domains"].post.requestBody'})

  // 4. List all operationIds grouped by tag (= api.<tag>.<operationId>)
  search({code: 'async () => {
    const out = {};
    for (const [path, methods] of Object.entries(spec.paths)) {
      for (const [m, op] of Object.entries(methods)) {
        if (m === "parameters" || !op.operationId) continue;
        const tag = (op.tags?.[0] ?? "default");
        (out[tag] ||= []).push({op: op.operationId, method: m.toUpperCase(), path});
      }
    }
    return out;
  }'})

Then call execute with api.<tag>.<operationId>(args).`

func NewDeployExamplePrompt() (mcp.Prompt, PromptHandlerFunc) {
	p := mcp.NewPrompt(
		"code-mode/deploy-example",
		mcp.WithPromptDescription("Code Mode pattern: single execute() that deploys + waits, with pollJob fallback."),
	)
	handler := func(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		return &mcp.GetPromptResult{
			Description: "Deploy + wait pattern using Code Mode execute() and pollJob().",
			Messages: []mcp.PromptMessage{
				mcp.NewPromptMessage(mcp.RoleUser, mcp.NewTextContent(deployExamplePromptText)),
			},
		}, nil
	}
	return p, handler
}

func NewAPIDiscoveryPrompt() (mcp.Prompt, PromptHandlerFunc) {
	p := mcp.NewPrompt(
		"code-mode/api-discovery",
		mcp.WithPromptDescription("Code Mode pattern: 4 search() snippets that walk you from keyword to operationId."),
	)
	handler := func(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		return &mcp.GetPromptResult{
			Description: "Discover the right operationId via search().",
			Messages: []mcp.PromptMessage{
				mcp.NewPromptMessage(mcp.RoleUser, mcp.NewTextContent(apiDiscoveryPromptText)),
			},
		}, nil
	}
	return p, handler
}

// PromptHandlerFunc mirrors server.PromptHandlerFunc.
type PromptHandlerFunc func(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error)
