package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type TriggerLatestDeploymentParams struct {
	ProjectID string `json:"project_id"`
	Branch    string `json:"branch"`
}

func TriggerLatestDeploymentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[TriggerLatestDeploymentParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	queryParams := map[string]string{
		"branch": params.Branch,
	}

	return makePostRequest(fmt.Sprintf("/v1/projects/%s/trigger-latest", params.ProjectID), queryParams, authInfo)
}
