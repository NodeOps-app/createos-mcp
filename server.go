package main

import (
	handler "github.com/NodeOps-app/autogen-backend-v2-mcp/handlers"
	"github.com/NodeOps-app/autogen-backend-v2-mcp/mcptools"
	"github.com/mark3labs/mcp-go/server"
)

func NewMCPServer(toolset string) *server.MCPServer {
	// Create a new MCP server
	s := server.NewMCPServer(
		"CreateOS MCP Server",
		"0.0.1",
		server.WithToolCapabilities(true),
		server.WithLogging(),
	)

	registerCoreTools(s)
	if toolset == "all" {
		registerLegacyTools(s)
	}

	return s
}

func registerCoreTools(s *server.MCPServer) {
	s.AddTool(mcptools.NewCreateOSProjectListMCPTool(), handler.CreateOSProjectListHandler)
	s.AddTool(mcptools.NewCreateOSProjectGetMCPTool(), handler.CreateOSProjectGetHandler)
	s.AddTool(mcptools.NewCreateOSProjectCreateVCSMCPTool(), handler.CreateOSProjectCreateVCSHandler)
	s.AddTool(mcptools.NewCreateOSProjectCreateImageMCPTool(), handler.CreateOSProjectCreateImageHandler)
	s.AddTool(mcptools.NewCreateOSProjectCreateUploadMCPTool(), handler.CreateOSProjectCreateUploadHandler)
	s.AddTool(mcptools.NewCreateOSDeploymentListMCPTool(), handler.CreateOSDeploymentListHandler)
	s.AddTool(mcptools.NewCreateOSDeploymentGetMCPTool(), handler.CreateOSDeploymentGetHandler)
	s.AddTool(mcptools.NewCreateOSDeploymentTriggerLatestMCPTool(), handler.CreateOSDeploymentTriggerLatestHandler)
	s.AddTool(mcptools.NewCreateOSDeploymentCancelMCPTool(), handler.CreateOSDeploymentCancelHandler)
	s.AddTool(mcptools.NewCreateOSDomainListMCPTool(), handler.CreateOSDomainListHandler)
	s.AddTool(mcptools.NewCreateOSGithubRepositoriesListMCPTool(), handler.CreateOSGithubRepositoriesListHandler)
	s.AddTool(mcptools.NewCreateOSGithubRepositoryBranchesListMCPTool(), handler.CreateOSGithubRepositoryBranchesListHandler)
	s.AddTool(mcptools.NewCreateOSProjectEnvironmentAnalyticsMCPTool(), handler.CreateOSProjectEnvironmentAnalyticsHandler)
}

func registerLegacyTools(s *server.MCPServer) {
	// Register all tools - existing handlers
	s.AddTool(mcptools.NewAssignDeploymentToProjectEnvironmentMCPTool(), handler.AssignDeploymentToProjectEnvironmentHandler)
	s.AddTool(mcptools.NewCancelDeploymentMCPTool(), handler.CancelDeploymentHandler)
	s.AddTool(mcptools.NewCheckAPIKeyUniqueNameMCPTool(), handler.CheckAPIKeyUniqueNameHandler)
	s.AddTool(mcptools.NewCheckProjectEnvironmentUniqueNameMCPTool(), handler.CheckProjectEnvironmentUniqueNameHandler)
	s.AddTool(mcptools.NewCheckProjectUniqueNameMCPTool(), handler.CheckProjectUniqueNameHandler)
	s.AddTool(mcptools.NewCreateAPIKeyMCPTool(), handler.CreateAPIKeyHandler)
	s.AddTool(mcptools.NewCreateDeploymentMCPTool(), handler.CreateDeploymentHandler)
	s.AddTool(mcptools.NewCreateProjectMCPTool(), handler.CreateProjectHandler)
	s.AddTool(mcptools.NewCreateProjectEnvironmentMCPTool(), handler.CreateProjectEnvironmentHandler)
	s.AddTool(mcptools.NewDeleteDeploymentMCPTool(), handler.DeleteDeploymentHandler)
	s.AddTool(mcptools.NewDeleteProjectMCPTool(), handler.DeleteProjectHandler)
	s.AddTool(mcptools.NewDeleteProjectEnvironmentMCPTool(), handler.DeleteProjectEnvironmentHandler)
	s.AddTool(mcptools.NewGetBuildLogsMCPTool(), handler.GetBuildLogsHandler)
	s.AddTool(mcptools.NewGetCurrentUserMCPTool(), handler.GetCurrentUserHandler)
	s.AddTool(mcptools.NewGetDeploymentMCPTool(), handler.GetDeploymentHandler)
	s.AddTool(mcptools.NewGetDeploymentLogsMCPTool(), handler.GetDeploymentLogsHandler)
	s.AddTool(mcptools.NewGetGithubRepositoryContentMCPTool(), handler.GetGithubRepositoryContentHandler)
	s.AddTool(mcptools.NewGetProjectMCPTool(), handler.GetProjectHandler)
	s.AddTool(mcptools.NewGetProjectEnvironmentAnalyticsMCPTool(), handler.GetProjectEnvironmentAnalyticsHandler)
	s.AddTool(mcptools.NewGetProjectEnvironmentAnalyticsOverallRequestsMCPTool(), handler.GetProjectEnvironmentAnalyticsOverallRequestsHandler)
	s.AddTool(mcptools.NewGetProjectEnvironmentAnalyticsRPMMCPTool(), handler.GetProjectEnvironmentAnalyticsRPMHandler)
	s.AddTool(mcptools.NewGetProjectEnvironmentAnalyticsRequestDistributionMCPTool(), handler.GetProjectEnvironmentAnalyticsRequestDistributionHandler)
	s.AddTool(mcptools.NewGetProjectEnvironmentAnalyticsSuccessPercentageMCPTool(), handler.GetProjectEnvironmentAnalyticsSuccessPercentageHandler)
	s.AddTool(mcptools.NewGetProjectEnvironmentAnalyticsTopErrorPathsMCPTool(), handler.GetProjectEnvironmentAnalyticsTopErrorPathsHandler)
	s.AddTool(mcptools.NewGetProjectEnvironmentAnalyticsTopHitPathsMCPTool(), handler.GetProjectEnvironmentAnalyticsTopHitPathsHandler)
	s.AddTool(mcptools.NewGetProjectEnvironmentLogsMCPTool(), handler.GetProjectEnvironmentLogsHandler)
	s.AddTool(mcptools.NewGetSupportedProjectTypesMCPTool(), handler.GetSupportedProjectTypesHandler)
	s.AddTool(mcptools.NewInstallGithubAppMCPTool(), handler.InstallGithubAppHandler)
	s.AddTool(mcptools.NewListAPIKeysMCPTool(), handler.ListAPIKeysHandler)
	s.AddTool(mcptools.NewListConnectedGithubAccountsMCPTool(), handler.ListConnectedGithubAccountsHandler)
	s.AddTool(mcptools.NewListDeploymentsMCPTool(), handler.ListDeploymentsHandler)
	s.AddTool(mcptools.NewListGithubRepositoriesMCPTool(), handler.ListGithubRepositoriesHandler)
	s.AddTool(mcptools.NewListGithubRepositoryBranchesMCPTool(), handler.ListGithubRepositoryBranchesHandler)
	s.AddTool(mcptools.NewListProjectEnvironmentsMCPTool(), handler.ListProjectEnvironmentsHandler)
	s.AddTool(mcptools.NewListProjectsMCPTool(), handler.ListProjectsHandler)
	s.AddTool(mcptools.NewRetriggerDeploymentMCPTool(), handler.RetriggerDeploymentHandler)
	s.AddTool(mcptools.NewRevokeAPIKeyMCPTool(), handler.RevokeAPIKeyHandler)
	s.AddTool(mcptools.NewTriggerLatestDeploymentMCPTool(), handler.TriggerLatestDeploymentHandler)
	s.AddTool(mcptools.NewUpdateAPIKeyMCPTool(), handler.UpdateAPIKeyHandler)
	s.AddTool(mcptools.NewUpdateProjectMCPTool(), handler.UpdateProjectHandler)
	s.AddTool(mcptools.NewUpdateProjectEnvironmentMCPTool(), handler.UpdateProjectEnvironmentHandler)
	s.AddTool(mcptools.NewUpdateProjectEnvironmentEnvironmentVariablesMCPTool(), handler.UpdateProjectEnvironmentEnvironmentVariablesHandler)
	s.AddTool(mcptools.NewUpdateProjectEnvironmentResourcesMCPTool(), handler.UpdateProjectEnvironmentResourcesHandler)
	s.AddTool(mcptools.NewUpdateProjectSettingsMCPTool(), handler.UpdateProjectSettingsHandler)
	s.AddTool(mcptools.NewWakeupDeploymentMCPTool(), handler.WakeupDeploymentHandler)

	// Register new tools
	s.AddTool(mcptools.NewAddProjectsToAppMCPTool(), handler.AddProjectsToAppHandler)
	s.AddTool(mcptools.NewAddServicesToAppMCPTool(), handler.AddServicesToAppHandler)
	s.AddTool(mcptools.NewCreateAppMCPTool(), handler.CreateAppHandler)
	s.AddTool(mcptools.NewCreateDomainMCPTool(), handler.CreateDomainHandler)
	s.AddTool(mcptools.NewDeleteAppMCPTool(), handler.DeleteAppHandler)
	s.AddTool(mcptools.NewDeleteDomainMCPTool(), handler.DeleteDomainHandler)
	s.AddTool(mcptools.NewDownloadDeploymentMCPTool(), handler.DownloadDeploymentHandler)
	s.AddTool(mcptools.NewGetProjectEnvironmentAnalyticsRequestsOverTimeMCPTool(), handler.GetProjectEnvironmentAnalyticsRequestsOverTimeHandler)
	s.AddTool(mcptools.NewGetProjectTransferUriMCPTool(), handler.GetProjectTransferUriHandler)
	s.AddTool(mcptools.NewGetQuotasMCPTool(), handler.GetQuotasHandler)
	s.AddTool(mcptools.NewGetSecurityScanMCPTool(), handler.GetSecurityScanHandler)
	s.AddTool(mcptools.NewGetSecurityScanDownloadUriMCPTool(), handler.GetSecurityScanDownloadUriHandler)
	s.AddTool(mcptools.NewListAppsMCPTool(), handler.ListAppsHandler)
	s.AddTool(mcptools.NewListDomainsMCPTool(), handler.ListDomainsHandler)
	s.AddTool(mcptools.NewListProjectTransferHistoryMCPTool(), handler.ListProjectTransferHistoryHandler)
	s.AddTool(mcptools.NewListProjectsByAppMCPTool(), handler.ListProjectsByAppHandler)
	s.AddTool(mcptools.NewListServicesByAppMCPTool(), handler.ListServicesByAppHandler)
	s.AddTool(mcptools.NewRefreshDomainMCPTool(), handler.RefreshDomainHandler)
	s.AddTool(mcptools.NewRemoveProjectsFromAppMCPTool(), handler.RemoveProjectsFromAppHandler)
	s.AddTool(mcptools.NewRemoveServicesFromAppMCPTool(), handler.RemoveServicesFromAppHandler)
	s.AddTool(mcptools.NewRetriggerSecurityScanMCPTool(), handler.RetriggerSecurityScanHandler)
	s.AddTool(mcptools.NewTransferProjectMCPTool(), handler.TransferProjectHandler)
	s.AddTool(mcptools.NewTriggerSecurityScanMCPTool(), handler.TriggerSecurityScanHandler)
	s.AddTool(mcptools.NewUpdateAppMCPTool(), handler.UpdateAppHandler)
	s.AddTool(mcptools.NewUpdateDomainEnvironmentMCPTool(), handler.UpdateDomainEnvironmentHandler)
	s.AddTool(mcptools.NewUploadDeploymentBase64FilesMCPTool(), handler.UploadDeploymentBase64FilesHandler)
	s.AddTool(mcptools.NewUploadDeploymentFilesMCPTool(), handler.UploadDeploymentFilesHandler)
	s.AddTool(mcptools.NewUploadDeploymentZipMCPTool(), handler.UploadDeploymentZipHandler)
}
