package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type GetProjectEnvironmentLogsParams struct {
	ProjectID     string `json:"project_id"`
	EnvironmentID string `json:"environment_id"`
}

func GetProjectEnvironmentLogsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[GetProjectEnvironmentLogsParams](args)
	if err != nil {
		return nil, err
	}

	return makeGetRequest(fmt.Sprintf("/v1/projects/%s/environments/%s/logs", params.ProjectID, params.EnvironmentID), nil, authInfo)
}
