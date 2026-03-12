package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type CancelDeploymentParams struct {
	ProjectID    string `json:"project_id"`
	DeploymentID string `json:"deployment_id"`
}

func CancelDeploymentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CancelDeploymentParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makePostRequest(fmt.Sprintf("/v1/projects/%s/deployments/%s/cancel", params.ProjectID, params.DeploymentID), nil, authInfo)
}
