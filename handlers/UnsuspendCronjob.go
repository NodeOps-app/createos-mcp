package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type UnsuspendCronjobParams struct {
	ProjectID string `json:"project_id"`
	CronjobID string `json:"cronjob_id"`
}

func UnsuspendCronjobHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[UnsuspendCronjobParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makePostRequest(fmt.Sprintf("/v1/projects/%s/cronjobs/%s/unsuspend", params.ProjectID, params.CronjobID), nil, authInfo)
}
