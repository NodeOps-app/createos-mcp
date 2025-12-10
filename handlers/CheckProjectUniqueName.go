package handler

import (
	"context"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type CheckProjectUniqueNameParams struct {
	Body map[string]interface{} `json:"body"`
}

func CheckProjectUniqueNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiKey, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CheckProjectUniqueNameParams](args)
	if err != nil {
		return nil, err
	}

	return makePostRequest("/v1/projects/available-unique-name", params.Body, apiKey)
}
