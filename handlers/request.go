package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

// AuthInfo contains authentication method and value
type AuthInfo struct {
	Method string // "api-key" or "bearer-token"
	Value  string
}

// GetAuthInfo extracts authentication information from the request
func GetAuthInfo(ctx context.Context, request mcp.CallToolRequest) (*AuthInfo, error) {
	// Check for X-Api-Key first
	apiKey := request.Header.Get("X-Api-Key")
	if apiKey != "" {
		return &AuthInfo{
			Method: "api-key",
			Value:  apiKey,
		}, nil
	}

	// Check for Bearer token
	authHeader := request.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) >= 2 && strings.ToLower(parts[0]) == "bearer" {
			return &AuthInfo{
				Method: "bearer-token",
				Value:  parts[1],
			}, nil
		}
	}

	return nil, fmt.Errorf("authentication required: either X-Api-Key or Authorization Bearer token")
}

// GetAPIKey is kept for backward compatibility but now uses GetAuthInfo
func GetAPIKey(ctx context.Context, request mcp.CallToolRequest) (string, error) {
	authInfo, err := GetAuthInfo(ctx, request)
	if err != nil {
		return "", err
	}
	return authInfo.Value, nil
}
