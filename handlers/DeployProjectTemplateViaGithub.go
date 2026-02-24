package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type DeployProjectTemplateViaGithubParams struct {
	PurchaseID string                 `json:"purchase_id"`
	Body       map[string]interface{} `json:"body"`
}

func DeployProjectTemplateViaGithubHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[DeployProjectTemplateViaGithubParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makePutRequest(fmt.Sprintf("/v1/project-templates/purchases/%s/deploy-via-github", params.PurchaseID), params.Body, authInfo)
}
