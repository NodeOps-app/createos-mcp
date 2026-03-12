package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type UpdateProjectEnvironmentResourcesParams struct {
	ProjectID    string                 `json:"project_id"`
	ProjectEnvID string                 `json:"environment_id"`
	Body         map[string]interface{} `json:"body"`
}

func UpdateProjectEnvironmentResourcesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[UpdateProjectEnvironmentResourcesParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makePutRequest(fmt.Sprintf("/v1/projects/%s/environments/%s/resources", params.ProjectID, params.ProjectEnvID), params.Body, authInfo)
}
