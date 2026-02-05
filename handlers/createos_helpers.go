package handler

import (
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

type backendPagination struct {
	Total  int `json:"total"`
	Count  int `json:"count"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type backendListResponse[T any] struct {
	Data struct {
		Data       []T               `json:"data"`
		Pagination backendPagination `json:"pagination"`
	} `json:"data"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type backendGetResponse[T any] struct {
	Data    T      `json:"data"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type listPage struct {
	Limit      int    `json:"limit"`
	Count      int    `json:"count"`
	TotalCount int    `json:"total_count"`
	HasMore    bool   `json:"has_more"`
	NextCursor string `json:"next_cursor,omitempty"`
}

func writeJSONResult(payload interface{}) (*mcp.CallToolResult, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to encode response: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(body)),
		},
	}, nil
}

func buildListPage(pagination backendPagination, nextCursor string) listPage {
	hasMore := pagination.Offset+pagination.Count < pagination.Total
	if !hasMore {
		nextCursor = ""
	}

	return listPage{
		Limit:      pagination.Limit,
		Count:      pagination.Count,
		TotalCount: pagination.Total,
		HasMore:    hasMore,
		NextCursor: nextCursor,
	}
}
