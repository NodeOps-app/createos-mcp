package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type ListProjectTransferHistoryParams struct {
	ProjectID string `json:"project_id"`
}

func ListProjectTransferHistoryHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[ListProjectTransferHistoryParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	// GET /v1/projects/{project_id}/transfer-history
	return makeGetRequest(fmt.Sprintf("/v1/projects/%s/transfer-history", params.ProjectID), nil, authInfo)
}
