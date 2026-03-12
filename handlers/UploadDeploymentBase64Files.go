package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type UploadDeploymentBase64FilesParams struct {
	ProjectID string                 `json:"project_id"`
	Body      map[string]interface{} `json:"body"`
}

func UploadDeploymentBase64FilesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[UploadDeploymentBase64FilesParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	// PUT /v1/projects/{project_id}/deployments/files/base64
	return makePutRequest(fmt.Sprintf("/v1/projects/%s/deployments/files/base64", params.ProjectID), params.Body, authInfo)
}
