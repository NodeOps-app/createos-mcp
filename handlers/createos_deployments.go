package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type CreateOSDeploymentListParams struct {
	ProjectID string  `json:"project_id"`
	Limit     *int    `json:"limit"`
	Cursor    *string `json:"cursor"`
}

type createOSDeploymentListItem struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	Environment string `json:"environment_id,omitempty"`
}

type createOSDeploymentListResponse struct {
	Items []createOSDeploymentListItem `json:"items"`
	Page  listPage                     `json:"page"`
}

func CreateOSDeploymentListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateOSDeploymentListParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	limit := defaultListLimit
	if params.Limit != nil {
		limit = *params.Limit
	}
	if limit <= 0 {
		limit = defaultListLimit
	}
	if limit > maxListLimit {
		limit = maxListLimit
	}

	cursor := ""
	if params.Cursor != nil {
		cursor = *params.Cursor
	}
	offset, err := mcputils.DecodeCursor(cursor)
	if err != nil {
		return nil, fmt.Errorf("invalid cursor: %w", err)
	}

	queryParams := map[string]string{
		"limit":  strconv.Itoa(limit),
		"offset": strconv.Itoa(offset),
	}

	resp, err := mcputils.Get(fmt.Sprintf("/v1/projects/%s/deployments", params.ProjectID), queryParams, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	var backendResp backendListResponse[struct {
		ID            string `json:"id"`
		Status        string `json:"status"`
		CreatedAt     string `json:"createdAt"`
		UpdatedAt     string `json:"updatedAt"`
		EnvironmentID string `json:"environmentId"`
	}]
	if err := json.Unmarshal(resp.Body(), &backendResp); err != nil {
		return nil, fmt.Errorf("failed to parse deployments response: %w", err)
	}

	items := make([]createOSDeploymentListItem, 0, len(backendResp.Data.Data))
	for _, item := range backendResp.Data.Data {
		items = append(items, createOSDeploymentListItem{
			ID:          item.ID,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			Environment: item.EnvironmentID,
		})
	}

	nextCursor := mcputils.EncodeCursor(backendResp.Data.Pagination.Offset + backendResp.Data.Pagination.Count)
	response := createOSDeploymentListResponse{
		Items: items,
		Page:  buildListPage(backendResp.Data.Pagination, nextCursor),
	}

	return writeJSONResult(response)
}

type CreateOSDeploymentGetParams struct {
	DeploymentID string `json:"deployment_id"`
}

type createOSDeploymentGetResponse struct {
	ID            string `json:"id"`
	Status        string `json:"status"`
	EnvironmentID string `json:"environment_id,omitempty"`
	ProjectID     string `json:"project_id,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
}

func CreateOSDeploymentGetHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateOSDeploymentGetParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	resp, err := mcputils.Get(fmt.Sprintf("/v1/deployments/%s", params.DeploymentID), nil, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	var backendResp backendGetResponse[struct {
		ID            string `json:"id"`
		Status        string `json:"status"`
		EnvironmentID string `json:"environmentId"`
		ProjectID     string `json:"projectId"`
		CreatedAt     string `json:"createdAt"`
		UpdatedAt     string `json:"updatedAt"`
	}]
	if err := json.Unmarshal(resp.Body(), &backendResp); err != nil {
		return nil, fmt.Errorf("failed to parse deployment response: %w", err)
	}

	result := createOSDeploymentGetResponse{
		ID:            backendResp.Data.ID,
		Status:        backendResp.Data.Status,
		EnvironmentID: backendResp.Data.EnvironmentID,
		ProjectID:     backendResp.Data.ProjectID,
		CreatedAt:     backendResp.Data.CreatedAt,
		UpdatedAt:     backendResp.Data.UpdatedAt,
	}

	return writeJSONResult(result)
}

type CreateOSDeploymentTriggerLatestParams struct {
	ProjectID string `json:"project_id"`
}

func CreateOSDeploymentTriggerLatestHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateOSDeploymentTriggerLatestParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	resp, err := mcputils.Post(fmt.Sprintf("/v1/projects/%s/deployments/latest/trigger", params.ProjectID), nil, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{Content: []mcp.Content{mcp.NewTextContent(string(resp.Body()))}}, nil
}

type CreateOSDeploymentCancelParams struct {
	DeploymentID string `json:"deployment_id"`
}

func CreateOSDeploymentCancelHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateOSDeploymentCancelParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	resp, err := mcputils.Post(fmt.Sprintf("/v1/deployments/%s/cancel", params.DeploymentID), nil, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{Content: []mcp.Content{mcp.NewTextContent(string(resp.Body()))}}, nil
}
