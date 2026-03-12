package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type UpdateProjectTemplateParams struct {
	TemplateID string                 `json:"template_id"`
	Body       map[string]interface{} `json:"body"`
}

func UpdateProjectTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[UpdateProjectTemplateParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makePutRequest(fmt.Sprintf("/v1/project-templates/%s", params.TemplateID), params.Body, authInfo)
}
