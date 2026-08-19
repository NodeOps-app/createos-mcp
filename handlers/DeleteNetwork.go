package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type DeleteNetworkParams struct {
	ID string `json:"id"`
}

func DeleteNetworkHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[DeleteNetworkParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makeSandboxDeleteRequest(fmt.Sprintf("/v1/networks/%s", params.ID), authInfo)
}
