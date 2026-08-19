package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type ExecSandboxParams struct {
	ID   string                 `json:"id"`
	Body map[string]interface{} `json:"body"`
}

func ExecSandboxHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[ExecSandboxParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}
	if stream, ok := params.Body["stream"].(bool); ok && stream {
		return nil, fmt.Errorf("streaming exec is not supported by ExecSandbox yet; omit stream or set it to false")
	}

	return makeSandboxPostRequest(fmt.Sprintf("/v1/sandboxes/%s/exec", params.ID), params.Body, authInfo)
}
