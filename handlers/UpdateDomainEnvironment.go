package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type UpdateDomainEnvironmentParams struct {
	ProjectID string                 `json:"project_id"`
	DomainID  string                 `json:"domain_id"`
	Body      map[string]interface{} `json:"body"`
}

func UpdateDomainEnvironmentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[UpdateDomainEnvironmentParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	// PUT /v1/projects/{project_id}/domains/{domain_id}/environment
	return makePutRequest(fmt.Sprintf("/v1/projects/%s/domains/%s/environment", params.ProjectID, params.DomainID), params.Body, authInfo)
}
