package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type GetSecurityScanDownloadUriParams struct {
	ProjectID    string `json:"project_id"`
	DeploymentID string `json:"deployment_id"`
}

func GetSecurityScanDownloadUriHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[GetSecurityScanDownloadUriParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	// GET /v1/projects/{project_id}/deployments/{deployment_id}/security-scans/download
	return makeGetRequest(fmt.Sprintf("/v1/projects/%s/deployments/%s/security-scans/download", params.ProjectID, params.DeploymentID), nil, authInfo)
}
