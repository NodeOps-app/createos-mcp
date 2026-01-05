package handler

import (
	"context"
	"fmt"
	"strconv"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type ListProjectsParams struct {
	Limit       *int  `json:"limit"`
	Offset      *int  `json:"offset"`
	ShowDeleted *bool `json:"show-deleted"`
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
	if params.ShowDeleted != nil {
		queryParams["show-deleted"] = strconv.FormatBool(*params.ShowDeleted)
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
