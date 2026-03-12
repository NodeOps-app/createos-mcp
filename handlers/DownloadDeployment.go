package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type DownloadDeploymentParams struct {
	ProjectID    string `json:"project_id"`
	DeploymentID string `json:"deployment_id"`
}

func DownloadDeploymentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[DownloadDeploymentParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makeGetRequest(fmt.Sprintf("/v1/projects/%s/deployments/%s/download", params.ProjectID, params.DeploymentID), nil, authInfo)
}
