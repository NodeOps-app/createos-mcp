package handler

import (
	"context"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type CreateAPIKeyParams struct {
	Body map[string]interface{} `json:"body"`
}

func CreateAPIKeyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateAPIKeyParams](args)
	if err != nil {
		return nil, err
	}

	return makePostRequest("/v1/api-keys", params.Body, authInfo)
}
