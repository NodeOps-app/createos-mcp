package handler

import (
	"context"
	"fmt"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

func handleRequest(ctx context.Context, request mcp.CallToolRequest) (string, map[string]interface{}, error) {
	apiKey, err := GetAPIKey(ctx, request)
	if err != nil {
		return "", nil, err
	}

	args, ok := request.Params.Arguments.(map[string]interface{})
	if !ok {
		return "", nil, fmt.Errorf("invalid arguments type")
	}

	return apiKey, args, nil
}

func makeGetRequest(path string, queryParams map[string]string, apiKey string) (*mcp.CallToolResult, error) {
	resp, err := mcputils.Get(path, queryParams, apiKey)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}

func makePostRequest(path string, body interface{}, apiKey string) (*mcp.CallToolResult, error) {
	resp, err := mcputils.Post(path, body, apiKey)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}

func makePutRequest(path string, body interface{}, apiKey string) (*mcp.CallToolResult, error) {
	resp, err := mcputils.Put(path, body, apiKey)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}

func makePatchRequest(path string, body interface{}, apiKey string) (*mcp.CallToolResult, error) {
	resp, err := mcputils.Patch(path, body, apiKey)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}

func makeDeleteRequest(path string, apiKey string) (*mcp.CallToolResult, error) {
	resp, err := mcputils.Delete(path, apiKey)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resp.Body())),
		},
	}, nil
}
