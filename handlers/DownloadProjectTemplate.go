package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type DownloadProjectTemplateParams struct {
	PurchaseID string `json:"purchase_id"`
}

func DownloadProjectTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[DownloadProjectTemplateParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	return makeGetRequest(fmt.Sprintf("/v1/project-templates/purchases/%s/download", params.PurchaseID), nil, authInfo)
}
