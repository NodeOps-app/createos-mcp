package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type ListProjectTemplateCountsParams struct {
	Name *string `json:"name"`
}

func ListProjectTemplateCountsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[ListProjectTemplateCountsParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	queryParams := make(map[string]string)
	if params.Name != nil {
		queryParams["name"] = *params.Name
	}

	return makeGetRequest("/v1/project-templates/counts", queryParams, authInfo)
}
