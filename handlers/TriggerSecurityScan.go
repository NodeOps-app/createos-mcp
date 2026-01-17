package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type TriggerSecurityScanParams struct {
	ProjectID    string `json:"project_id"`
	DeploymentID string `json:"deployment_id"`
}

func TriggerSecurityScanHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[TriggerSecurityScanParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	// POST /v1/projects/{project_id}/deployments/{deployment_id}/security-scans/trigger
	return makePostRequest(fmt.Sprintf("/v1/projects/%s/deployments/%s/security-scans/trigger", params.ProjectID, params.DeploymentID), nil, authInfo)
}
