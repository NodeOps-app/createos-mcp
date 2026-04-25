package codemode

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

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
		return mcp.NewToolResultError(fmt.Sprintf("code mode unavailable: %v", err)), nil
	}
	out, err := json.Marshal(result)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("encode: %v", err)), nil
	}
	return mcp.NewToolResultText(string(out)), nil
}
