package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

// CreateProjectParams represents the input parameters for CreateProject
type CreateProjectParams struct {
	Body map[string]interface{} `json:"body"`
}

// CreateProjectHandler is the handler function for the CreateProject tool.
func CreateProjectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, err := GetAuthInfo(ctx, request)
	if err != nil {
		return nil, err
	}

	// Parse parameters
	args, ok := request.Params.Arguments.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid arguments type")
	}

	params, err := mcputils.ParamsParser[CreateProjectParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	// Make POST request with body
	resp, err := mcputils.Post("/v1/projects", params.Body, authInfo.Method, authInfo.Value)
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
