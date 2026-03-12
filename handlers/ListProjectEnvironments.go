package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type ListProjectEnvironmentsParams struct {
	ProjectID string `json:"project_id"`
}

func ListProjectEnvironmentsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[ListProjectEnvironmentsParams](args)
	if err != nil {
		return nil, err
	}

	return makeGetRequest(fmt.Sprintf("/v1/projects/%s/environments", params.ProjectID), nil, authInfo)
}
