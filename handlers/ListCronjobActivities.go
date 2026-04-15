package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type ListCronjobActivitiesParams struct {
	ProjectID string `json:"project_id"`
	CronjobID string `json:"cronjob_id"`
}

func ListCronjobActivitiesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[ListCronjobActivitiesParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makeGetRequest(fmt.Sprintf("/v1/projects/%s/cronjobs/%s/activities", params.ProjectID, params.CronjobID), nil, authInfo)
}
