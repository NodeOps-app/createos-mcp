package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type CreateDeploymentParams struct {
	ProjectID string                 `json:"project_id"`
	Body      map[string]interface{} `json:"body"`
}

func CreateDeploymentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateDeploymentParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makePostRequest(fmt.Sprintf("/v1/projects/%s/deployments", params.ProjectID), params.Body, authInfo)
}

