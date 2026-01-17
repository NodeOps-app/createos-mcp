package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type TransferProjectParams struct {
	ProjectID string `json:"project_id"`
	Token     string `json:"token"`
}

func TransferProjectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[TransferProjectParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	queryParams := map[string]string{
		"token": params.Token,
	}

	// GET /v1/projects/{project_id}/transfer?token={token}
	return makeGetRequest(fmt.Sprintf("/v1/projects/%s/transfer", params.ProjectID), queryParams, authInfo)
}
