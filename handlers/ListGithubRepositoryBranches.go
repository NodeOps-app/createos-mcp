package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type ListGithubRepositoryBranchesParams struct {
	InstallationID string `json:"installation_id"`
	Repository     string `json:"repository"`
}

func ListGithubRepositoryBranchesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiKey, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[ListGithubRepositoryBranchesParams](args)
	if err != nil {
		return nil, err
	}

	queryParams := map[string]string{
		"repository": params.Repository,
	}
	return makeGetRequest(fmt.Sprintf("/v1/app-installations/github/accounts/%s/repositories/branches", params.InstallationID), queryParams, apiKey)
}
