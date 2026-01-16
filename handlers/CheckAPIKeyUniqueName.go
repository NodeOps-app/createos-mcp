package handler

import (
	"context"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type CheckAPIKeyUniqueNameParams struct {
	Body map[string]interface{} `json:"body"`
}

func CheckAPIKeyUniqueNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CheckAPIKeyUniqueNameParams](args)
	if err != nil {
		return nil, err
	}

	return makePostRequest("/v1/api-keys/available-unique-name", params.Body, authInfo)
}
