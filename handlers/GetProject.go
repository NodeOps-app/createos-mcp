package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type GetProjectParams struct {
	ProjectID string `json:"project_id"`
}

// GetProjectHandler is the handler function for the GetProject tool.
func GetProjectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, err := GetAuthInfo(ctx, request)
	if err != nil {
		return nil, err
	}

	// Parse parameters
	args, ok := request.Params.Arguments.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid arguments type")
	}

	params, err := mcputils.ParamsParser[GetProjectParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	// Make GET request
	resp, err := mcputils.Get(fmt.Sprintf("/v1/projects/%s", params.ProjectID), nil, authInfo.Method, authInfo.Value)
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
