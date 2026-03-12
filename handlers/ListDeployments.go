package handler

import (
	"context"
	"fmt"
	"strconv"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type ListDeploymentsParams struct {
	ProjectID string `json:"project_id"`
	Limit     *int   `json:"limit"`
	Offset    *int   `json:"offset"`
}

func ListDeploymentsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[ListDeploymentsParams](args)
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

	return makeGetRequest(fmt.Sprintf("/v1/projects/%s/deployments", params.ProjectID), queryParams, authInfo)
}
