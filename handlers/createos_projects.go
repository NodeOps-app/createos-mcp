package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	mcputils "github.com/NodeOps-app/autogen-backend-v2-mcp/helpers"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	defaultListLimit = 20
	maxListLimit     = 50
)

type CreateOSProjectListParams struct {
	App    *string `json:"app"`
	Name   *string `json:"name"`
	Status *string `json:"status"`
	Type   *string `json:"type"`
	Limit  *int    `json:"limit"`
	Cursor *string `json:"cursor"`
}

type createOSProjectListItem struct {
	ID          string  `json:"id"`
	AppID       *string `json:"app_id,omitempty"`
	DisplayName string  `json:"display_name"`
	UniqueName  string  `json:"unique_name"`
	Status      string  `json:"status"`
	Type        string  `json:"type"`
	CreatedAt   string  `json:"created_at,omitempty"`
	UpdatedAt   string  `json:"updated_at,omitempty"`
}

type createOSProjectListResponse struct {
	Items []createOSProjectListItem `json:"items"`
	Page  listPage                  `json:"page"`
}

func CreateOSProjectListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateOSProjectListParams](args)
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
	if params.App != nil {
		queryParams["app"] = *params.App
	}
	if params.Name != nil {
		queryParams["name"] = *params.Name
	}
	if params.Status != nil {
		queryParams["status"] = *params.Status
	}
	if params.Type != nil {
		queryParams["type"] = *params.Type
	}

	resp, err := mcputils.Get("/v1/projects", queryParams, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	var backendResp backendListResponse[struct {
		ID          string  `json:"id"`
		AppID       *string `json:"appId"`
		DisplayName string  `json:"displayName"`
		UniqueName  string  `json:"uniqueName"`
		Status      string  `json:"status"`
		Type        string  `json:"type"`
		CreatedAt   string  `json:"createdAt"`
		UpdatedAt   string  `json:"updatedAt"`
	}]
	if err := json.Unmarshal(resp.Body(), &backendResp); err != nil {
		return nil, fmt.Errorf("failed to parse list projects response: %w", err)
	}

	items := make([]createOSProjectListItem, 0, len(backendResp.Data.Data))
	for _, item := range backendResp.Data.Data {
		items = append(items, createOSProjectListItem{
			ID:          item.ID,
			AppID:       item.AppID,
			DisplayName: item.DisplayName,
			UniqueName:  item.UniqueName,
			Status:      item.Status,
			Type:        item.Type,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	nextCursor := mcputils.EncodeCursor(backendResp.Data.Pagination.Offset + backendResp.Data.Pagination.Count)
	response := createOSProjectListResponse{
		Items: items,
		Page:  buildListPage(backendResp.Data.Pagination, nextCursor),
	}

	return writeJSONResult(response)
}

type CreateOSProjectGetParams struct {
	ProjectID string `json:"project_id"`
}

type createOSProjectGetResponse struct {
	ID          string  `json:"id"`
	AppID       *string `json:"app_id,omitempty"`
	UniqueName  string  `json:"unique_name"`
	DisplayName string  `json:"display_name"`
	Type        string  `json:"type"`
	Status      string  `json:"status"`
	Description *string `json:"description,omitempty"`
	CreatedAt   string  `json:"created_at,omitempty"`
	UpdatedAt   string  `json:"updated_at,omitempty"`
}

func CreateOSProjectGetHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateOSProjectGetParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	resp, err := mcputils.Get(fmt.Sprintf("/v1/projects/%s", params.ProjectID), nil, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	var backendResp backendGetResponse[struct {
		ID          string  `json:"id"`
		AppID       *string `json:"appId"`
		UniqueName  string  `json:"uniqueName"`
		DisplayName string  `json:"displayName"`
		Type        string  `json:"type"`
		Status      string  `json:"status"`
		Description *string `json:"description"`
		CreatedAt   string  `json:"createdAt"`
		UpdatedAt   string  `json:"updatedAt"`
	}]
	if err := json.Unmarshal(resp.Body(), &backendResp); err != nil {
		return nil, fmt.Errorf("failed to parse project response: %w", err)
	}

	result := createOSProjectGetResponse{
		ID:          backendResp.Data.ID,
		AppID:       backendResp.Data.AppID,
		UniqueName:  backendResp.Data.UniqueName,
		DisplayName: backendResp.Data.DisplayName,
		Type:        backendResp.Data.Type,
		Status:      backendResp.Data.Status,
		Description: backendResp.Data.Description,
		CreatedAt:   backendResp.Data.CreatedAt,
		UpdatedAt:   backendResp.Data.UpdatedAt,
	}

	return writeJSONResult(result)
}

type CreateOSProjectCreateVCSParams struct {
	AppID               *string           `json:"app_id"`
	UniqueName          string            `json:"unique_name"`
	DisplayName         string            `json:"display_name"`
	Description         *string           `json:"description"`
	EnabledSecurityScan *bool             `json:"enabled_security_scan"`
	VCSName             *string           `json:"vcs_name"`
	VCSInstallationID   string            `json:"vcs_installation_id"`
	VCSRepoID           string            `json:"vcs_repo_id"`
	BuildCommand        *string           `json:"build_command"`
	BuildDir            *string           `json:"build_dir"`
	BuildFlag           *string           `json:"build_flag"`
	BuildFlags          map[string]string `json:"build_flags"`
	BuildVars           map[string]string `json:"build_vars"`
	DirectoryPath       *string           `json:"directory_path"`
	Framework           *string           `json:"framework"`
	HasDockerfile       *bool             `json:"has_dockerfile"`
	IgnoreBranches      []string          `json:"ignore_branches"`
	InstallCommand      *string           `json:"install_command"`
	Port                *int              `json:"port"`
	RunCommand          *string           `json:"run_command"`
	RunEnvs             map[string]string `json:"run_envs"`
	RunFlag             *string           `json:"run_flag"`
	Runtime             *string           `json:"runtime"`
	UseBuildAI          *bool             `json:"use_build_ai"`
}

func CreateOSProjectCreateVCSHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateOSProjectCreateVCSParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	vcsName := "github"
	if params.VCSName != nil && *params.VCSName != "" {
		vcsName = *params.VCSName
	}

	settings := map[string]interface{}{}
	if params.BuildCommand != nil {
		settings["buildCommand"] = *params.BuildCommand
	}
	if params.BuildDir != nil {
		settings["buildDir"] = *params.BuildDir
	}
	if params.BuildFlag != nil {
		settings["buildFlag"] = *params.BuildFlag
	}
	if len(params.BuildFlags) > 0 {
		settings["buildFlags"] = params.BuildFlags
	}
	if len(params.BuildVars) > 0 {
		settings["buildVars"] = params.BuildVars
	}
	if params.DirectoryPath != nil {
		settings["directoryPath"] = *params.DirectoryPath
	}
	if params.Framework != nil {
		settings["framework"] = *params.Framework
	}
	if params.HasDockerfile != nil {
		settings["hasDockerfile"] = *params.HasDockerfile
	}
	if len(params.IgnoreBranches) > 0 {
		settings["ignoreBranches"] = params.IgnoreBranches
	}
	if params.InstallCommand != nil {
		settings["installCommand"] = *params.InstallCommand
	}
	if params.Port != nil {
		settings["port"] = *params.Port
	}
	if params.RunCommand != nil {
		settings["runCommand"] = *params.RunCommand
	}
	if len(params.RunEnvs) > 0 {
		settings["runEnvs"] = params.RunEnvs
	}
	if params.RunFlag != nil {
		settings["runFlag"] = *params.RunFlag
	}
	if params.Runtime != nil {
		settings["runtime"] = *params.Runtime
	}
	if params.UseBuildAI != nil {
		settings["useBuildAI"] = *params.UseBuildAI
	}

	body := map[string]interface{}{
		"uniqueName": params.UniqueName,
		"displayName": params.DisplayName,
		"type":        "vcs",
		"settings":    settings,
		"source": map[string]interface{}{
			"vcsName":           vcsName,
			"vcsInstallationId": params.VCSInstallationID,
			"vcsRepoId":         params.VCSRepoID,
		},
	}
	if params.AppID != nil {
		body["appId"] = *params.AppID
	}
	if params.Description != nil {
		body["description"] = *params.Description
	}
	if params.EnabledSecurityScan != nil {
		body["enabledSecurityScan"] = *params.EnabledSecurityScan
	}

	resp, err := mcputils.Post("/v1/projects", body, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{Content: []mcp.Content{mcp.NewTextContent(string(resp.Body()))}}, nil
}

type CreateOSProjectCreateImageParams struct {
	AppID               *string           `json:"app_id"`
	UniqueName          string            `json:"unique_name"`
	DisplayName         string            `json:"display_name"`
	Description         *string           `json:"description"`
	EnabledSecurityScan *bool             `json:"enabled_security_scan"`
	Port                int               `json:"port"`
	RunEnvs             map[string]string `json:"run_envs"`
}

func CreateOSProjectCreateImageHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateOSProjectCreateImageParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	settings := map[string]interface{}{
		"port": params.Port,
	}
	if len(params.RunEnvs) > 0 {
		settings["runEnvs"] = params.RunEnvs
	}

	body := map[string]interface{}{
		"uniqueName": params.UniqueName,
		"displayName": params.DisplayName,
		"type":        "image",
		"settings":    settings,
		"source":      map[string]interface{}{},
	}
	if params.AppID != nil {
		body["appId"] = *params.AppID
	}
	if params.Description != nil {
		body["description"] = *params.Description
	}
	if params.EnabledSecurityScan != nil {
		body["enabledSecurityScan"] = *params.EnabledSecurityScan
	}

	resp, err := mcputils.Post("/v1/projects", body, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{Content: []mcp.Content{mcp.NewTextContent(string(resp.Body()))}}, nil
}

type CreateOSProjectCreateUploadParams struct {
	AppID               *string           `json:"app_id"`
	UniqueName          string            `json:"unique_name"`
	DisplayName         string            `json:"display_name"`
	Description         *string           `json:"description"`
	EnabledSecurityScan *bool             `json:"enabled_security_scan"`
	BuildCommand        *string           `json:"build_command"`
	BuildDir            *string           `json:"build_dir"`
	BuildFlag           *string           `json:"build_flag"`
	BuildVars           map[string]string `json:"build_vars"`
	DirectoryPath       *string           `json:"directory_path"`
	Framework           *string           `json:"framework"`
	HasDockerfile       *bool             `json:"has_dockerfile"`
	InstallCommand      *string           `json:"install_command"`
	Port                *int              `json:"port"`
	RunCommand          *string           `json:"run_command"`
	RunEnvs             map[string]string `json:"run_envs"`
	RunFlag             *string           `json:"run_flag"`
	Runtime             *string           `json:"runtime"`
	UseBuildAI          *bool             `json:"use_build_ai"`
}

func CreateOSProjectCreateUploadHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	authInfo, args, err := handleRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	params, err := mcputils.ParamsParser[CreateOSProjectCreateUploadParams](args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	settings := map[string]interface{}{}
	if params.BuildCommand != nil {
		settings["buildCommand"] = *params.BuildCommand
	}
	if params.BuildDir != nil {
		settings["buildDir"] = *params.BuildDir
	}
	if params.BuildFlag != nil {
		settings["buildFlag"] = *params.BuildFlag
	}
	if len(params.BuildVars) > 0 {
		settings["buildVars"] = params.BuildVars
	}
	if params.DirectoryPath != nil {
		settings["directoryPath"] = *params.DirectoryPath
	}
	if params.Framework != nil {
		settings["framework"] = *params.Framework
	}
	if params.HasDockerfile != nil {
		settings["hasDockerfile"] = *params.HasDockerfile
	}
	if params.InstallCommand != nil {
		settings["installCommand"] = *params.InstallCommand
	}
	if params.Port != nil {
		settings["port"] = *params.Port
	}
	if params.RunCommand != nil {
		settings["runCommand"] = *params.RunCommand
	}
	if len(params.RunEnvs) > 0 {
		settings["runEnvs"] = params.RunEnvs
	}
	if params.RunFlag != nil {
		settings["runFlag"] = *params.RunFlag
	}
	if params.Runtime != nil {
		settings["runtime"] = *params.Runtime
	}
	if params.UseBuildAI != nil {
		settings["useBuildAI"] = *params.UseBuildAI
	}

	body := map[string]interface{}{
		"uniqueName": params.UniqueName,
		"displayName": params.DisplayName,
		"type":        "upload",
		"settings":    settings,
		"source":      map[string]interface{}{},
	}
	if params.AppID != nil {
		body["appId"] = *params.AppID
	}
	if params.Description != nil {
		body["description"] = *params.Description
	}
	if params.EnabledSecurityScan != nil {
		body["enabledSecurityScan"] = *params.EnabledSecurityScan
	}

	resp, err := mcputils.Post("/v1/projects", body, authInfo.Method, authInfo.Value)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResult{Content: []mcp.Content{mcp.NewTextContent(string(resp.Body()))}}, nil
}
