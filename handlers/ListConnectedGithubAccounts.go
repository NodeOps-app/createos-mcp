package handler

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

func ListConnectedGithubAccountsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, _, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	return makeGetRequest("/v1/app-installations/github/accounts", nil, authInfo)
}
