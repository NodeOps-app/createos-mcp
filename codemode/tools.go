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
      "description": "JavaScript async arrow function. Available globals: spec (CreateOS OpenAPI, $refs resolved), console, sleep. Return any JSON-serialisable value."
    }
  },
  "required": ["code"]
}`

func NewSearchTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"search",
		"Search the CreateOS OpenAPI spec. Write a JS async arrow function that reads `spec` and returns matches. No network access in this mode. Use to discover endpoints, parameters, and shapes before calling execute().",
		[]byte(searchInputSchema),
	)
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
      "description": "Job id returned by execute() when status was running."
    }
  },
  "required": ["jobId"]
}`

func NewPollJobTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"pollJob",
		"Poll a long-running execute() job. Blocks up to ~90s; returns {status: 'running'|'done'|'error', ...}. Loop until status != 'running'.",
		[]byte(pollJobInputSchema),
	)
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
      "description": "JavaScript async arrow function. Available globals: spec, api, console, sleep. api.<group>.<operationId>(args) calls the CreateOS API with caller auth. api.raw(method, path, opts) is an escape hatch. Throws ApiError on non-2xx. Return any JSON-serialisable value."
    }
  },
  "required": ["code"]
}`

func NewExecuteTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"execute",
		"Execute JavaScript against the CreateOS API. Single sandbox call may chain multiple api calls. Returns either {status:'done', result, logs} or {status:'running', jobId} (long-running ops, polled via pollJob).",
		[]byte(executeInputSchema),
	)
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
