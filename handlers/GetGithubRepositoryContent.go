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

	return makePostRequest(fmt.Sprintf("/v1/app-installations/github/accounts/%s/repositories/content", params.InstallationID), params.Body, apiKey)
}
