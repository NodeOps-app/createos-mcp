package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type DeleteProjectEnvironmentParams struct {
	ProjectID    string `json:"project_id"`
	ProjectEnvID string `json:"environment_id"`
}

func DeleteProjectEnvironmentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[DeleteProjectEnvironmentParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makeDeleteRequest(fmt.Sprintf("/v1/projects/%s/environments/%s", params.ProjectID, params.ProjectEnvID), authInfo)
}
