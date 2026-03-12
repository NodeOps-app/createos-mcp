package handler

import (
	"context"
	"fmt"
	"strconv"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type GetProjectEnvironmentAnalyticsRequestsOverTimeParams struct {
	ProjectID     string `json:"project_id"`
	EnvironmentID string `json:"environment_id"`
	Start         int64  `json:"start,omitempty"`
	End           int64  `json:"end,omitempty"`
}

func GetProjectEnvironmentAnalyticsRequestsOverTimeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[GetProjectEnvironmentAnalyticsRequestsOverTimeParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	queryParams := make(map[string]string)
	if params.Start > 0 {
		queryParams["start"] = strconv.FormatInt(params.Start, 10)
	}
	if params.End > 0 {
		queryParams["end"] = strconv.FormatInt(params.End, 10)
	}

	// GET /v1/projects/{project_id}/environments/{environment_id}/analytics/requests-over-time
	return makeGetRequest(fmt.Sprintf("/v1/projects/%s/environments/%s/analytics/requests-over-time", params.ProjectID, params.EnvironmentID), queryParams, authInfo)
}
