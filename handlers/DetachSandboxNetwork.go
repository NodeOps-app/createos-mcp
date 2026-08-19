package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type DetachSandboxNetworkParams struct {
	ID      string `json:"id"`
	Network string `json:"network"`
}

func DetachSandboxNetworkHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[DetachSandboxNetworkParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makeSandboxDeleteRequest(fmt.Sprintf("/v1/sandboxes/%s/networks/%s", params.ID, params.Network), authInfo)
}
