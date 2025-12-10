package handler

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func GetAPIKey(ctx context.Context, request mcp.CallToolRequest) (string, error) {
	apiKey := request.Header.Get("X-Api-Key")
	if apiKey == "" {
		return "", fmt.Errorf("X-Api-Key header is required")
	}
	return apiKey, nil
}
