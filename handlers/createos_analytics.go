package handler

import (
	"context"
	"fmt"
	"strconv"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type CreateOSProjectEnvironmentAnalyticsParams struct {
	ProjectID     string `json:"project_id"`
	EnvironmentID string `json:"environment_id"`
	Metric        string `json:"metric"`
	Start         *int64 `json:"start"`
	End           *int64 `json:"end"`
}

func CreateOSProjectEnvironmentAnalyticsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateOSProjectEnvironmentAnalyticsParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	metricPath, ok := map[string]string{
		"requests_over_time": "requests-over-time",
		"rpm":                "rpm",
		"success_percentage": "success-percentage",
		"request_distribution": "request-distribution",
		"top_error_paths":    "top-error-paths",
		"top_hit_paths":      "top-hit-paths",
		"overall_requests":   "overall-requests",
	}[params.Metric]
	if !ok {
		return nil, fmt.Errorf("invalid metric: %s", params.Metric)
	}

	queryParams := make(map[string]string)
	if params.Start != nil {
		queryParams["start"] = strconv.FormatInt(*params.Start, 10)
	}
	if params.End != nil {
		queryParams["end"] = strconv.FormatInt(*params.End, 10)
	}

	resp, err := mcputils.Get(fmt.Sprintf("/v1/projects/%s/environments/%s/analytics/%s", params.ProjectID, params.EnvironmentID, metricPath), queryParams, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{Content: []mcp.Content{mcp.NewTextContent(string(resp.Body()))}}, nil
}
