package handler

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/NodeOps-app/createos-mcp/pkg/requestid"
	"github.com/mark3labs/mcp-go/mcp"
)

// AuthInfo contains authentication method and value
type AuthInfo struct {
	Context context.Context
	Method  string // "api-key" or "bearer-token"
	Value   string
}

// GetAuthInfo extracts authentication information from the request
func GetAuthInfo(ctx context.Context, request mcp.CallToolRequest) (*AuthInfo, error) {
	id := requestid.FromContext(ctx)
	if id == "" {
		id = requestid.Resolve(request.Header.Get(requestid.Header))
	}
	ctx = requestid.WithContext(ctx, id)
	log.Printf("request_id=%s MCP tool requested tool=%q", id, request.Params.Name)
	// Check for X-Api-Key first
	apiKey := request.Header.Get("X-Api-Key")
	if apiKey != "" {
		return &AuthInfo{
			Context: ctx,
			Method:  "api-key",
			Value:   apiKey,
		}, nil
	}

	// Check for Bearer token
	authHeader := request.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) >= 2 && strings.ToLower(parts[0]) == "bearer" {
			return &AuthInfo{
				Context: ctx,
				Method:  "bearer-token",
				Value:   parts[1],
			}, nil
		}
	}

	return nil, fmt.Errorf("authentication required: either X-Api-Key or Authorization Bearer token")
}
