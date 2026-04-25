package codemode

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func observe(mode Mode, outcome string, start time.Time) {
	RunDuration.WithLabelValues(string(mode), outcome).Observe(time.Since(start).Seconds())
}

func outcomeFromResult(r *RunResult) string {
	switch r.Status {
	case "done":
		return "success"
	case "running":
		return "async"
	case "error":
		if r.ErrorKind != "" {
			return r.ErrorKind
		}
		return "user_error"
	default:
		return "unknown"
	}
}

const searchInputSchema = `{
  "type": "object",
  "properties": {
    "code": {
      "type": "string",
      "description": "Async arrow fn. Globals: spec, console, sleep. Example: 'async () => Object.entries(spec.paths).filter(([p]) => p.includes(\"deploy\")).map(([p,m]) => ({path:p, ops:Object.keys(m)}))'."
    }
  },
  "required": ["code"]
}`

const searchToolDescription = `Code Mode: introspect the CreateOS OpenAPI spec from a sandboxed JS arrow fn.

When: discover endpoints/params/schemas BEFORE calling execute. No network in this mode.

Globals: spec ($refs resolved, descriptions stripped), console, sleep(ms).
Result envelope: {status:"done"|"error", result, logs} | {status:"error", errorKind:"userCode"|"timeout"|"infra", error, stack}.
Limits: 5s timeout, 1MB code, 64KB result.

See resource code-mode://intro and prompt code-mode/api-discovery.`

func toolMeta() *mcp.Meta {
	return &mcp.Meta{
		AdditionalFields: map[string]any{
			"mode":    "code",
			"sandbox": "workerd",
			"version": "v2",
		},
	}
}

func NewSearchTool() mcp.Tool {
	t := mcp.NewToolWithRawSchema("search", searchToolDescription, []byte(searchInputSchema))
	t.Meta = toolMeta()
	return t
}

type Handler struct {
	Client *Client
}

func (h *Handler) Search(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()
	outcome := "bad_input"
	defer func() { observe(ModeSearch, outcome, start) }()

	args, ok := req.Params.Arguments.(map[string]interface{})
	if !ok {
		return mcp.NewToolResultError("invalid arguments type"), nil
	}
	rawCode, ok := args["code"]
	if !ok {
		return mcp.NewToolResultError("missing required argument: code"), nil
	}
	code, ok := rawCode.(string)
	if !ok {
		return mcp.NewToolResultError("code must be a string"), nil
	}
	result, err := h.Client.Run(ctx, RunRequest{Mode: ModeSearch, Code: code})
	if err != nil {
		outcome = "infra"
		return mcp.NewToolResultError(fmt.Sprintf("code mode unavailable: %v", err)), nil
	}
	outcome = outcomeFromResult(result)
	out, err := json.Marshal(result)
	if err != nil {
		outcome = "encode"
		return mcp.NewToolResultError(fmt.Sprintf("encode: %v", err)), nil
	}
	return mcp.NewToolResultText(string(out)), nil
}

const pollJobInputSchema = `{
  "type": "object",
  "properties": {
    "jobId": {
      "type": "string",
      "description": "j_<uuid> from a prior execute() that returned status:'running'."
    }
  },
  "required": ["jobId"]
}`

const pollJobToolDescription = `Code Mode: drain a long-running execute() job.

When: a prior execute() returned {status:"running", jobId}. Loop until status != "running".
Blocks up to ~90s per call (long-poll).

Result envelope: {status:"running", jobId, logsSoFar} | {status:"done", result, logs} | {status:"error", errorKind:"userCode"|"timeout"|"infra"|"jobMissing", error, stack}.
Wall cap on the underlying job: 600s. Jobs evicted 30 min after terminal.`

func NewPollJobTool() mcp.Tool {
	t := mcp.NewToolWithRawSchema("pollJob", pollJobToolDescription, []byte(pollJobInputSchema))
	t.Meta = toolMeta()
	return t
}

func (h *Handler) PollJob(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, ok := req.Params.Arguments.(map[string]interface{})
	if !ok {
		return mcp.NewToolResultError("invalid arguments type"), nil
	}
	rawID, ok := args["jobId"]
	if !ok {
		return mcp.NewToolResultError("missing required argument: jobId"), nil
	}
	jobID, ok := rawID.(string)
	if !ok || jobID == "" {
		return mcp.NewToolResultError("jobId must be a non-empty string"), nil
	}
	result, err := h.Client.Poll(ctx, jobID)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("code mode unavailable: %v", err)), nil
	}
	out, err := json.Marshal(result)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("encode: %v", err)), nil
	}
	return mcp.NewToolResultText(string(out)), nil
}

const executeInputSchema = `{
  "type": "object",
  "properties": {
    "code": {
      "type": "string",
      "description": "Async arrow fn. Globals: spec, api, console, sleep. Example: 'async () => { const {data} = await api.projects.listProjects({limit:5}); return data; }'. Or escape hatch: 'async () => (await api.raw(\"POST\",\"/v1/x\",{body:{...}})).body'."
    }
  },
  "required": ["code"]
}`

const executeToolDescription = `Code Mode: run JS against the CreateOS API in a fresh V8 isolate. Auth forwarded from caller.

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

See resource code-mode://api-shape and prompt code-mode/deploy-example.`

func NewExecuteTool() mcp.Tool {
	t := mcp.NewToolWithRawSchema("execute", executeToolDescription, []byte(executeInputSchema))
	t.Meta = toolMeta()
	return t
}

func (h *Handler) Execute(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()
	outcome := "bad_input"
	defer func() { observe(ModeExecute, outcome, start) }()

	args, ok := req.Params.Arguments.(map[string]interface{})
	if !ok {
		return mcp.NewToolResultError("invalid arguments type"), nil
	}
	rawCode, ok := args["code"]
	if !ok {
		return mcp.NewToolResultError("missing required argument: code"), nil
	}
	code, ok := rawCode.(string)
	if !ok {
		return mcp.NewToolResultError("code must be a string"), nil
	}
	auth := AuthFromRequest(ctx, req)
	result, err := h.Client.Run(ctx, RunRequest{Mode: ModeExecute, Code: code, AuthCtx: auth})
	if err != nil {
		outcome = "infra"
		return mcp.NewToolResultError(fmt.Sprintf("code mode unavailable: %v", err)), nil
	}
	outcome = outcomeFromResult(result)
	out, err := json.Marshal(result)
	if err != nil {
		outcome = "encode"
		return mcp.NewToolResultError(fmt.Sprintf("encode: %v", err)), nil
	}
	return mcp.NewToolResultText(string(out)), nil
}
