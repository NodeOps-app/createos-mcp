package handler

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

func GetSupportedProjectTypesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiKey, _, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	return makeGetRequest("/v1/projects/types", nil, apiKey)
}

