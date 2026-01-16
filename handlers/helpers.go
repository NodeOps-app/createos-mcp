package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

func handleRequest(ctx context.Context, request mcp.CallToolRequest) (*AuthInfo, map[string]interface{}, error) {
	authInfo, err := GetAuthInfo(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	args, ok := request.Params.Arguments.(map[string]interface{})
	if !ok {
		return nil, nil, fmt.Errorf("invalid arguments type")
	}

	return authInfo, args, nil
}

func makeGetRequest(path string, queryParams map[string]string, authInfo *AuthInfo) (*mcp.CallToolResult, error) {
	resp, err := mcputils.Get(path, queryParams, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}

func makePostRequest(path string, body interface{}, authInfo *AuthInfo) (*mcp.CallToolResult, error) {
	resp, err := mcputils.Post(path, body, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}

func makePutRequest(path string, body interface{}, authInfo *AuthInfo) (*mcp.CallToolResult, error) {
	resp, err := mcputils.Put(path, body, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}

func makePatchRequest(path string, body interface{}, authInfo *AuthInfo) (*mcp.CallToolResult, error) {
	resp, err := mcputils.Patch(path, body, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}

func makeDeleteRequest(path string, authInfo *AuthInfo) (*mcp.CallToolResult, error) {
	resp, err := mcputils.Delete(path, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}
