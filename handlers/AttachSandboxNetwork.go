package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type AttachSandboxNetworkParams struct {
	ID   string                 `json:"id"`
	Body map[string]interface{} `json:"body"`
}

func AttachSandboxNetworkHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[AttachSandboxNetworkParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makeSandboxPostRequest(fmt.Sprintf("/v1/sandboxes/%s/networks", params.ID), params.Body, authInfo)
}
