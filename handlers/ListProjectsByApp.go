package handler

import (
	"context"
	"fmt"
	"strconv"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type ListProjectsByAppParams struct {
	AppID  string `json:"app_id"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
}

func ListProjectsByAppHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[ListProjectsByAppParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	queryParams := make(map[string]string)
	if params.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(params.Limit)
	}
	if params.Offset > 0 {
		queryParams["offset"] = strconv.Itoa(params.Offset)
	}

	// GET /v1/apps/{app_id}/projects
	return makeGetRequest(fmt.Sprintf("/v1/apps/%s/projects", params.AppID), queryParams, authInfo)
}
