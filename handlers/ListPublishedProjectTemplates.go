package handler

import (
	"context"
	"fmt"
	"strconv"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type ListPublishedProjectTemplatesParams struct {
	Limit      *int    `json:"limit"`
	Offset     *int    `json:"offset"`
	Categories *string `json:"categories"`
	Name       *string `json:"name"`
}

func ListPublishedProjectTemplatesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[ListPublishedProjectTemplatesParams](args)
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
	if params.Categories != nil {
		queryParams["categories"] = *params.Categories
	}
	if params.Name != nil {
		queryParams["name"] = *params.Name
	}

	return makeGetRequest("/v1/project-templates", queryParams, authInfo)
}
