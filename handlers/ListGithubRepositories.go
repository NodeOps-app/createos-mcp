package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type ListGithubRepositoriesParams struct {
	InstallationID string `json:"installation_id"`
}

func ListGithubRepositoriesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiKey, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[ListGithubRepositoriesParams](args)
	if err != nil {
		return nil, err
	}

	queryParams := map[string]string{"installation_id": params.InstallationID}
	return makeGetRequest(fmt.Sprintf("/v1/app-installations/github/accounts/%s/repositories", params.InstallationID), queryParams, apiKey)
}
