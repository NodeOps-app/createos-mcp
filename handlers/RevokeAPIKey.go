package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type RevokeAPIKeyParams struct {
	APIKeyID string `json:"api_key_id"`
}

func RevokeAPIKeyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiKey, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[RevokeAPIKeyParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makeDeleteRequest(fmt.Sprintf("/v1/api-keys/%s", params.APIKeyID), apiKey)
}

