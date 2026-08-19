package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type GetDiskParams struct {
	IDOrName string `json:"id_or_name"`
}

func GetDiskHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[GetDiskParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makeSandboxGetRequest(fmt.Sprintf("/v1/disks/%s", params.IDOrName), nil, authInfo)
}
