package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type UploadDeploymentZipParams struct {
	ProjectID string                 `json:"project_id"`
	Body      map[string]interface{} `json:"body"`
}

func UploadDeploymentZipHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[UploadDeploymentZipParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	// PUT /v1/projects/{project_id}/deployments/upload-zip
	return makePutRequest(fmt.Sprintf("/v1/projects/%s/deployments/upload-zip", params.ProjectID), params.Body, authInfo)
}
