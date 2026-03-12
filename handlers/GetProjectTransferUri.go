package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type GetProjectTransferUriParams struct {
	ProjectID string `json:"project_id"`
}

func GetProjectTransferUriHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[GetProjectTransferUriParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	// GET /v1/projects/{project_id}/transfer-uri
	return makeGetRequest(fmt.Sprintf("/v1/projects/%s/transfer-uri", params.ProjectID), nil, authInfo)
}
