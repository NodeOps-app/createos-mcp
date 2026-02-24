package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type GetProjectTemplateParams struct {
	TemplateID string `json:"template_id"`
}

func GetProjectTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[GetProjectTemplateParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makeGetRequest(fmt.Sprintf("/v1/project-templates/%s", params.TemplateID), nil, authInfo)
}
