package handler

import (
	"context"
	"fmt"
	"strconv"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type ListProjectTemplatePurchasesParams struct {
	Limit  *int `json:"limit"`
	Offset *int `json:"offset"`
}

func ListProjectTemplatePurchasesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[ListProjectTemplatePurchasesParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	queryParams := make(map[string]string)
	if params.Limit != nil {
		queryParams["limit"] = strconv.Itoa(*params.Limit)
	}
	if params.Offset != nil {
		queryParams["offset"] = strconv.Itoa(*params.Offset)
	}

	return makeGetRequest("/v1/project-templates/purchases", queryParams, authInfo)
}
