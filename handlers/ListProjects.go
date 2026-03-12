package handler

import (
	"context"
	"fmt"
	"strconv"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type ListProjectsParams struct {
	Limit  *int    `json:"limit"`
	Offset *int    `json:"offset"`
	Status *string `json:"status"`
	App    *string `json:"app"`
	Name   *string `json:"name"`
	Type   *string `json:"type"`
}

// ListProjectsHandler is the handler function for the ListProjects tool.
func ListProjectsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, err := GetAuthInfo(ctx, request)
	if err != nil {
		return nil, err
	}

	// Parse parameters
	args, ok := request.Params.Arguments.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid arguments type")
	}

	params, err := mcputils.ParamsParser[ListProjectsParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	// Build query parameters
	queryParams := make(map[string]string)
	if params.Limit != nil {
		queryParams["limit"] = strconv.Itoa(*params.Limit)
	}
	if params.Offset != nil {
		queryParams["offset"] = strconv.Itoa(*params.Offset)
	}
	if params.Status != nil {
		queryParams["status"] = *params.Status
	}
	if params.App != nil {
		queryParams["app"] = *params.App
	}
	if params.Name != nil {
		queryParams["name"] = *params.Name
	}
	if params.Type != nil {
		queryParams["type"] = *params.Type
	}

	// Make GET request using Resty
	resp, err := mcputils.Get("/v1/projects", queryParams, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	// Return the response
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}
