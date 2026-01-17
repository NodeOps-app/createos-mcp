package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type RemoveProjectsFromAppParams struct {
	AppID string                 `json:"app_id"`
	Body  map[string]interface{} `json:"body"`
}

func RemoveProjectsFromAppHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[RemoveProjectsFromAppParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	// POST /v1/apps/{app_id}/projects/remove
	return makePostRequest(fmt.Sprintf("/v1/apps/%s/projects/remove", params.AppID), params.Body, authInfo)
}
