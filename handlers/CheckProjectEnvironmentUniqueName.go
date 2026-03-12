package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type CheckProjectEnvironmentUniqueNameParams struct {
	ProjectID string                 `json:"project_id"`
	Body      map[string]interface{} `json:"body"`
}

func CheckProjectEnvironmentUniqueNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CheckProjectEnvironmentUniqueNameParams](args)
	if err != nil {
		return nil, err
	}

	return makePostRequest(fmt.Sprintf("/v1/projects/%s/environments/available-unique-name", params.ProjectID), params.Body, authInfo)
}
