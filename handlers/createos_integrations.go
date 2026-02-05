package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

type CreateOSDomainListParams struct {
	ProjectID string `json:"project_id"`
}

type createOSDomainListResponse struct {
	Items []struct {
		ID         string `json:"id"`
		Domain     string `json:"domain"`
		Status     string `json:"status,omitempty"`
		CreatedAt  string `json:"created_at,omitempty"`
		UpdatedAt  string `json:"updated_at,omitempty"`
		ProjectID  string `json:"project_id,omitempty"`
		EnvID      string `json:"environment_id,omitempty"`
		IsPrimary  *bool  `json:"is_primary,omitempty"`
		IsVerified *bool  `json:"is_verified,omitempty"`
	} `json:"items"`
}

func CreateOSDomainListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateOSDomainListParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	resp, err := mcputils.Get(fmt.Sprintf("/v1/projects/%s/domains", params.ProjectID), nil, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	var backendResp backendGetResponse[[]struct {
		ID            string `json:"id"`
		Domain        string `json:"domain"`
		Status        string `json:"status"`
		CreatedAt     string `json:"createdAt"`
		UpdatedAt     string `json:"updatedAt"`
		ProjectID     string `json:"projectId"`
		EnvironmentID string `json:"environmentId"`
		IsPrimary     *bool  `json:"isPrimary"`
		IsVerified    *bool  `json:"isVerified"`
	}]
	if err := json.Unmarshal(resp.Body(), &backendResp); err != nil {
		return nil, fmt.Errorf("failed to parse domains response: %w", err)
	}

	items := make([]struct {
		ID         string `json:"id"`
		Domain     string `json:"domain"`
		Status     string `json:"status,omitempty"`
		CreatedAt  string `json:"created_at,omitempty"`
		UpdatedAt  string `json:"updated_at,omitempty"`
		ProjectID  string `json:"project_id,omitempty"`
		EnvID      string `json:"environment_id,omitempty"`
		IsPrimary  *bool  `json:"is_primary,omitempty"`
		IsVerified *bool  `json:"is_verified,omitempty"`
	}, 0, len(backendResp.Data))

	for _, item := range backendResp.Data {
		items = append(items, struct {
			ID         string `json:"id"`
			Domain     string `json:"domain"`
			Status     string `json:"status,omitempty"`
			CreatedAt  string `json:"created_at,omitempty"`
			UpdatedAt  string `json:"updated_at,omitempty"`
			ProjectID  string `json:"project_id,omitempty"`
			EnvID      string `json:"environment_id,omitempty"`
			IsPrimary  *bool  `json:"is_primary,omitempty"`
			IsVerified *bool  `json:"is_verified,omitempty"`
		}{
			ID:         item.ID,
			Domain:     item.Domain,
			Status:     item.Status,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
			ProjectID:  item.ProjectID,
			EnvID:      item.EnvironmentID,
			IsPrimary:  item.IsPrimary,
			IsVerified: item.IsVerified,
		})
	}

	return writeJSONResult(createOSDomainListResponse{Items: items})
}

type CreateOSGithubRepositoriesListParams struct {
	InstallationID string `json:"installation_id"`
}

type createOSGithubRepositoriesListResponse struct {
	Items []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name,omitempty"`
		Private  *bool  `json:"private,omitempty"`
		URL      string `json:"url,omitempty"`
	} `json:"items"`
}

func CreateOSGithubRepositoriesListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateOSGithubRepositoriesListParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	queryParams := map[string]string{"installation_id": params.InstallationID}
	resp, err := mcputils.Get(fmt.Sprintf("/v1/app-installations/github/accounts/%s/repositories", params.InstallationID), queryParams, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	var backendResp backendGetResponse[[]struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"fullName"`
		Private  *bool  `json:"private"`
		URL      string `json:"url"`
	}]
	if err := json.Unmarshal(resp.Body(), &backendResp); err != nil {
		return nil, fmt.Errorf("failed to parse repositories response: %w", err)
	}

	items := make([]struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name,omitempty"`
		Private  *bool  `json:"private,omitempty"`
		URL      string `json:"url,omitempty"`
	}, 0, len(backendResp.Data))

	for _, item := range backendResp.Data {
		items = append(items, struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			FullName string `json:"full_name,omitempty"`
			Private  *bool  `json:"private,omitempty"`
			URL      string `json:"url,omitempty"`
		}{
			ID:       item.ID,
			Name:     item.Name,
			FullName: item.FullName,
			Private:  item.Private,
			URL:      item.URL,
		})
	}

	return writeJSONResult(createOSGithubRepositoriesListResponse{Items: items})
}

type CreateOSGithubRepositoryBranchesListParams struct {
	InstallationID string  `json:"installation_id"`
	Repository     string  `json:"repository"`
	Limit          *int    `json:"limit"`
	Cursor         *string `json:"cursor"`
}

type createOSGithubRepositoryBranchesListResponse struct {
	Items []struct {
		Name string `json:"name"`
		SHA  string `json:"sha,omitempty"`
	} `json:"items"`
	Page listPage `json:"page"`
}

func CreateOSGithubRepositoryBranchesListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateOSGithubRepositoryBranchesListParams](args)
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

	page := 1
	offset := 0
	if params.Cursor != nil && *params.Cursor != "" {
		decoded, err := mcputils.DecodeCursor(*params.Cursor)
		if err != nil {
			return nil, fmt.Errorf("invalid cursor: %w", err)
		}
		offset = decoded
		page = (decoded / limit) + 1
	}
	if offset == 0 {
		offset = (page - 1) * limit
	}

	queryParams := map[string]string{
		"repository": params.Repository,
		"page":       strconv.Itoa(page),
		"per-page":   strconv.Itoa(limit),
	}
	resp, err := mcputils.Get(fmt.Sprintf("/v1/app-installations/github/accounts/%s/repositories/branches", params.InstallationID), queryParams, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	var backendResp backendGetResponse[[]struct {
		Name string `json:"name"`
		SHA  string `json:"sha"`
	}]
	if err := json.Unmarshal(resp.Body(), &backendResp); err != nil {
		return nil, fmt.Errorf("failed to parse branches response: %w", err)
	}

	items := make([]struct {
		Name string `json:"name"`
		SHA  string `json:"sha,omitempty"`
	}, 0, len(backendResp.Data))
	for _, item := range backendResp.Data {
		items = append(items, struct {
			Name string `json:"name"`
			SHA  string `json:"sha,omitempty"`
		}{
			Name: item.Name,
			SHA:  item.SHA,
		})
	}

	hasMore := len(items) == limit
	nextCursor := ""
	if hasMore {
		nextCursor = mcputils.EncodeCursor(offset + len(items))
	}

	return writeJSONResult(createOSGithubRepositoryBranchesListResponse{
		Items: items,
		Page: listPage{
			Limit:      limit,
			Count:      len(items),
			TotalCount: 0,
			HasMore:    hasMore,
			NextCursor: nextCursor,
		},
	})
}
