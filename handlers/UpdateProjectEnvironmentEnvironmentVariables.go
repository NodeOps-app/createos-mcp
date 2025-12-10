package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type UpdateProjectEnvironmentEnvironmentVariablesParams struct {
	ProjectID    string                 `json:"project_id"`
	ProjectEnvID string                 `json:"environment_id"`
	Body         map[string]interface{} `json:"body"`
}

func UpdateProjectEnvironmentEnvironmentVariablesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiKey, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[UpdateProjectEnvironmentEnvironmentVariablesParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makePutRequest(fmt.Sprintf("/v1/projects/%s/environments/%s/environment-variables", params.ProjectID, params.ProjectEnvID), params.Body, apiKey)
}
