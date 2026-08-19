package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type DetachSandboxDiskParams struct {
	ID        string `json:"id"`
	DiskID    string `json:"disk_id"`
	MountPath string `json:"mount_path"`
}

func DetachSandboxDiskHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[DetachSandboxDiskParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	queryParams := map[string]string{
		"mount_path": params.MountPath,
	}

	return makeSandboxDeleteRequestWithQuery(
		fmt.Sprintf("/v1/sandboxes/%s/disks/%s", params.ID, params.DiskID),
		queryParams,
		authInfo,
	)
}
