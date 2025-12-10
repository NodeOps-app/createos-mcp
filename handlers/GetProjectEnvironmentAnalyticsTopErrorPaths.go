package handler

import (
	"context"
	"fmt"
	"strconv"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type GetProjectEnvironmentAnalyticsTopErrorPathsParams struct {
	ProjectID     string `json:"project_id"`
	EnvironmentID string `json:"environment_id"`
	Start         *int64 `json:"start"`
	End           *int64 `json:"end"`
}

func GetProjectEnvironmentAnalyticsTopErrorPathsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiKey, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[GetProjectEnvironmentAnalyticsTopErrorPathsParams](args)
	if err != nil {
		return nil, err
	}

	queryParams := make(map[string]string)
	if params.Start != nil {
		queryParams["start"] = strconv.FormatInt(*params.Start, 10)
	}
	if params.End != nil {
		queryParams["end"] = strconv.FormatInt(*params.End, 10)
	}

	return makeGetRequest(fmt.Sprintf("/v1/projects/%s/environments/%s/analytics/top-error-paths", params.ProjectID, params.EnvironmentID), queryParams, apiKey)
}

