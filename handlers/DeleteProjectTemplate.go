package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type DeleteProjectTemplateParams struct {
	TemplateID string `json:"template_id"`
}

func DeleteProjectTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[DeleteProjectTemplateParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makeDeleteRequest(fmt.Sprintf("/v1/project-templates/%s", params.TemplateID), authInfo)
}
