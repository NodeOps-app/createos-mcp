package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type AssignDeploymentToProjectEnvironmentParams struct {
	ProjectID    string                 `json:"project_id"`
	DeploymentID string                 `json:"deployment_id"`
	ProjectEnvID string                 `json:"environment_id"`
	Body         map[string]interface{} `json:"body"`
}

func AssignDeploymentToProjectEnvironmentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiKey, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[AssignDeploymentToProjectEnvironmentParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makePostRequest(fmt.Sprintf("/v1/projects/%s/environments/%s/assign-deployment", params.ProjectID, params.ProjectEnvID), params.Body, apiKey)
}
