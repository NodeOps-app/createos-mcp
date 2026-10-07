package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/createos-mcp/helpers"
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
	resp, err := mcputils.Get(authInfo.Context, path, queryParams, authInfo.Method, authInfo.Value)
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
	resp, err := mcputils.Post(authInfo.Context, path, body, authInfo.Method, authInfo.Value)
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
	resp, err := mcputils.Put(authInfo.Context, path, body, authInfo.Method, authInfo.Value)
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
	resp, err := mcputils.Patch(authInfo.Context, path, body, authInfo.Method, authInfo.Value)
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
	resp, err := mcputils.Delete(authInfo.Context, path, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}

func makeSandboxPostRequest(path string, body interface{}, authInfo *AuthInfo) (*mcp.CallToolResult, error) {
	resp, err := mcputils.SandboxPost(authInfo.Context, path, body, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}

func makeSandboxGetRequest(path string, queryParams map[string]string, authInfo *AuthInfo) (*mcp.CallToolResult, error) {
	resp, err := mcputils.SandboxGet(authInfo.Context, path, queryParams, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}

func makeSandboxPatchRequest(path string, body interface{}, authInfo *AuthInfo) (*mcp.CallToolResult, error) {
	resp, err := mcputils.SandboxPatch(authInfo.Context, path, body, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}

func makeSandboxDeleteRequest(path string, authInfo *AuthInfo) (*mcp.CallToolResult, error) {
	return makeSandboxDeleteRequestWithQuery(path, nil, authInfo)
}

func makeSandboxDeleteRequestWithQuery(path string, queryParams map[string]string, authInfo *AuthInfo) (*mcp.CallToolResult, error) {
	resp, err := mcputils.SandboxDelete(authInfo.Context, path, queryParams, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}
