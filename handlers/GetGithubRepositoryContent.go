package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type GetGithubRepositoryContentParams struct {
	InstallationID string                 `json:"installation_id"`
	Body           map[string]interface{} `json:"body"`
}

func GetGithubRepositoryContentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiKey, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[GetGithubRepositoryContentParams](args)
	if err != nil {
		return nil, err
	}

	// POST request with body and query params
	resp, err := mcputils.Client().R().
		SetHeader("X-Api-Key", apiKey).
		SetQueryParam("installation_id", params.InstallationID).
		SetBody(params.Body).
		Post("/v1/github/repositories/content")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode(), string(resp.Body()))
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}

