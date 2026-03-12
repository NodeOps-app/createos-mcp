package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type CreateDomainParams struct {
	ProjectID string                 `json:"project_id"`
	Body      map[string]interface{} `json:"body"`
}

func CreateDomainHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateDomainParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makePostRequest(fmt.Sprintf("/v1/projects/%s/domains", params.ProjectID), params.Body, authInfo)
}
