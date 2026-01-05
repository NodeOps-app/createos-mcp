package handler

import (
	"context"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type InstallGithubAppParams struct {
	Body map[string]interface{} `json:"body"`
}

func InstallGithubAppHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[InstallGithubAppParams](args)
	if err != nil {
		return nil, err
	}

	return makePostRequest("/v1/app-installations/github", params.Body, authInfo)
}
