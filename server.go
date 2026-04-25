package main

import (
	"os"

	"github.com/NodeOps-app/createos-mcp/codemode"
	handler "github.com/NodeOps-app/createos-mcp/handlers"
	"github.com/NodeOps-app/createos-mcp/mcptools"
	"github.com/mark3labs/mcp-go/server"
)

// NewMCPServer creates and returns an MCP server with all tools registered
func NewMCPServer() *server.MCPServer {
	// Create a new MCP server
	s := server.NewMCPServer(
		"MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithLogging(),
	)

	// Register all tools
	s.AddTool(mcptools.NewAddProjectsToAppMCPTool(), handler.AddProjectsToAppHandler)
	s.AddTool(mcptools.NewAddServicesToAppMCPTool(), handler.AddServicesToAppHandler)
	s.AddTool(mcptools.NewAssignDeploymentToProjectEnvironmentMCPTool(), handler.AssignDeploymentToProjectEnvironmentHandler)
	s.AddTool(mcptools.NewBuyProjectTemplateMCPTool(), handler.BuyProjectTemplateHandler)
	s.AddTool(mcptools.NewCancelDeploymentMCPTool(), handler.CancelDeploymentHandler)
	s.AddTool(mcptools.NewCheckAPIKeyUniqueNameMCPTool(), handler.CheckAPIKeyUniqueNameHandler)
	s.AddTool(mcptools.NewCheckProjectEnvironmentUniqueNameMCPTool(), handler.CheckProjectEnvironmentUniqueNameHandler)
	s.AddTool(mcptools.NewCheckProjectUniqueNameMCPTool(), handler.CheckProjectUniqueNameHandler)
	s.AddTool(mcptools.NewCreateAPIKeyMCPTool(), handler.CreateAPIKeyHandler)
	s.AddTool(mcptools.NewCreateAppMCPTool(), handler.CreateAppHandler)
	s.AddTool(mcptools.NewCreateCronjobMCPTool(), handler.CreateCronjobHandler)
	s.AddTool(mcptools.NewCreateDeploymentMCPTool(), handler.CreateDeploymentHandler)
	s.AddTool(mcptools.NewCreateDomainMCPTool(), handler.CreateDomainHandler)
	s.AddTool(mcptools.NewCreateProjectMCPTool(), handler.CreateProjectHandler)
	s.AddTool(mcptools.NewCreateProjectEnvironmentMCPTool(), handler.CreateProjectEnvironmentHandler)
	s.AddTool(mcptools.NewCreateProjectTemplateMCPTool(), handler.CreateProjectTemplateHandler)
	s.AddTool(mcptools.NewDeleteAppMCPTool(), handler.DeleteAppHandler)
	s.AddTool(mcptools.NewDeleteCronjobMCPTool(), handler.DeleteCronjobHandler)
	s.AddTool(mcptools.NewDeleteDeploymentMCPTool(), handler.DeleteDeploymentHandler)
	s.AddTool(mcptools.NewDeleteDomainMCPTool(), handler.DeleteDomainHandler)
	s.AddTool(mcptools.NewDeleteProjectMCPTool(), handler.DeleteProjectHandler)
	s.AddTool(mcptools.NewDeleteProjectEnvironmentMCPTool(), handler.DeleteProjectEnvironmentHandler)
	s.AddTool(mcptools.NewDeleteProjectTemplateMCPTool(), handler.DeleteProjectTemplateHandler)
	s.AddTool(mcptools.NewDeployProjectTemplateViaGithubMCPTool(), handler.DeployProjectTemplateViaGithubHandler)
	s.AddTool(mcptools.NewDownloadDeploymentMCPTool(), handler.DownloadDeploymentHandler)
	s.AddTool(mcptools.NewDownloadProjectTemplateMCPTool(), handler.DownloadProjectTemplateHandler)
	s.AddTool(mcptools.NewGetBuildLogsMCPTool(), handler.GetBuildLogsHandler)
	s.AddTool(mcptools.NewGetCronjobMCPTool(), handler.GetCronjobHandler)
	s.AddTool(mcptools.NewGetCurrentUserMCPTool(), handler.GetCurrentUserHandler)
	s.AddTool(mcptools.NewGetDeploymentMCPTool(), handler.GetDeploymentHandler)
	s.AddTool(mcptools.NewGetDeploymentLogsMCPTool(), handler.GetDeploymentLogsHandler)
	s.AddTool(mcptools.NewGetGithubRepositoryContentMCPTool(), handler.GetGithubRepositoryContentHandler)
	s.AddTool(mcptools.NewGetProjectMCPTool(), handler.GetProjectHandler)
	s.AddTool(mcptools.NewGetProjectEnvironmentLogsMCPTool(), handler.GetProjectEnvironmentLogsHandler)
	s.AddTool(mcptools.NewGetProjectTemplateMCPTool(), handler.GetProjectTemplateHandler)
	s.AddTool(mcptools.NewGetProjectTransferUriMCPTool(), handler.GetProjectTransferUriHandler)
	s.AddTool(mcptools.NewGetQuotasMCPTool(), handler.GetQuotasHandler)
	s.AddTool(mcptools.NewGetSecurityScanMCPTool(), handler.GetSecurityScanHandler)
	s.AddTool(mcptools.NewGetSecurityScanDownloadUriMCPTool(), handler.GetSecurityScanDownloadUriHandler)
	s.AddTool(mcptools.NewGetSupportedProjectTypesMCPTool(), handler.GetSupportedProjectTypesHandler)
	s.AddTool(mcptools.NewInstallGithubAppMCPTool(), handler.InstallGithubAppHandler)
	s.AddTool(mcptools.NewListAPIKeysMCPTool(), handler.ListAPIKeysHandler)
	s.AddTool(mcptools.NewListAppsMCPTool(), handler.ListAppsHandler)
	s.AddTool(mcptools.NewListConnectedGithubAccountsMCPTool(), handler.ListConnectedGithubAccountsHandler)
	s.AddTool(mcptools.NewListCronjobActivitiesMCPTool(), handler.ListCronjobActivitiesHandler)
	s.AddTool(mcptools.NewListCronjobsMCPTool(), handler.ListCronjobsHandler)
	s.AddTool(mcptools.NewListDeploymentsMCPTool(), handler.ListDeploymentsHandler)
	s.AddTool(mcptools.NewListDomainsMCPTool(), handler.ListDomainsHandler)
	s.AddTool(mcptools.NewListGithubRepositoriesMCPTool(), handler.ListGithubRepositoriesHandler)
	s.AddTool(mcptools.NewListGithubRepositoryBranchesMCPTool(), handler.ListGithubRepositoryBranchesHandler)
	s.AddTool(mcptools.NewListMyProjectTemplatesMCPTool(), handler.ListMyProjectTemplatesHandler)
	s.AddTool(mcptools.NewListProjectEnvironmentsMCPTool(), handler.ListProjectEnvironmentsHandler)
	s.AddTool(mcptools.NewListProjectTemplateCountsMCPTool(), handler.ListProjectTemplateCountsHandler)
	s.AddTool(mcptools.NewListProjectTemplatePurchasesMCPTool(), handler.ListProjectTemplatePurchasesHandler)
	s.AddTool(mcptools.NewListProjectTemplateRejectionsMCPTool(), handler.ListProjectTemplateRejectionsHandler)
	s.AddTool(mcptools.NewListProjectTransferHistoryMCPTool(), handler.ListProjectTransferHistoryHandler)
	s.AddTool(mcptools.NewListProjectsMCPTool(), handler.ListProjectsHandler)
	s.AddTool(mcptools.NewListProjectsByAppMCPTool(), handler.ListProjectsByAppHandler)
	s.AddTool(mcptools.NewListPublishedProjectTemplatesMCPTool(), handler.ListPublishedProjectTemplatesHandler)
	s.AddTool(mcptools.NewListServicesByAppMCPTool(), handler.ListServicesByAppHandler)
	s.AddTool(mcptools.NewRefreshDomainMCPTool(), handler.RefreshDomainHandler)
	s.AddTool(mcptools.NewRemoveProjectsFromAppMCPTool(), handler.RemoveProjectsFromAppHandler)
	s.AddTool(mcptools.NewRemoveServicesFromAppMCPTool(), handler.RemoveServicesFromAppHandler)
	s.AddTool(mcptools.NewRetriggerDeploymentMCPTool(), handler.RetriggerDeploymentHandler)
	s.AddTool(mcptools.NewRetriggerSecurityScanMCPTool(), handler.RetriggerSecurityScanHandler)
	s.AddTool(mcptools.NewRevokeAPIKeyMCPTool(), handler.RevokeAPIKeyHandler)
	s.AddTool(mcptools.NewSuspendCronjobMCPTool(), handler.SuspendCronjobHandler)
	s.AddTool(mcptools.NewTransferProjectMCPTool(), handler.TransferProjectHandler)
	s.AddTool(mcptools.NewTriggerLatestDeploymentMCPTool(), handler.TriggerLatestDeploymentHandler)
	s.AddTool(mcptools.NewTriggerSecurityScanMCPTool(), handler.TriggerSecurityScanHandler)
	s.AddTool(mcptools.NewUnsuspendCronjobMCPTool(), handler.UnsuspendCronjobHandler)
	s.AddTool(mcptools.NewUpdateAPIKeyMCPTool(), handler.UpdateAPIKeyHandler)
	s.AddTool(mcptools.NewUpdateAppMCPTool(), handler.UpdateAppHandler)
	s.AddTool(mcptools.NewUpdateCronjobMCPTool(), handler.UpdateCronjobHandler)
	s.AddTool(mcptools.NewUpdateDomainEnvironmentMCPTool(), handler.UpdateDomainEnvironmentHandler)
	s.AddTool(mcptools.NewUpdateProjectMCPTool(), handler.UpdateProjectHandler)
	s.AddTool(mcptools.NewUpdateProjectEnvironmentMCPTool(), handler.UpdateProjectEnvironmentHandler)
	s.AddTool(mcptools.NewUpdateProjectEnvironmentEnvironmentVariablesMCPTool(), handler.UpdateProjectEnvironmentEnvironmentVariablesHandler)
	s.AddTool(mcptools.NewUpdateProjectEnvironmentResourcesMCPTool(), handler.UpdateProjectEnvironmentResourcesHandler)
	s.AddTool(mcptools.NewUpdateProjectSettingsMCPTool(), handler.UpdateProjectSettingsHandler)
	s.AddTool(mcptools.NewUpdateProjectTemplateMCPTool(), handler.UpdateProjectTemplateHandler)
	s.AddTool(mcptools.NewUpdateProjectTemplateStatusMCPTool(), handler.UpdateProjectTemplateStatusHandler)
	s.AddTool(mcptools.NewUploadDeploymentBase64FilesMCPTool(), handler.UploadDeploymentBase64FilesHandler)
	s.AddTool(mcptools.NewUploadDeploymentFilesMCPTool(), handler.UploadDeploymentFilesHandler)
	s.AddTool(mcptools.NewUploadDeploymentZipMCPTool(), handler.UploadDeploymentZipHandler)
	s.AddTool(mcptools.NewWakeupDeploymentMCPTool(), handler.WakeupDeploymentHandler)

	// Code Mode (v2): search tool. execute + pollJob registered in later phases.
	workerdURL := os.Getenv("WORKERD_URL")
	if workerdURL == "" {
		workerdURL = "http://127.0.0.1:8787"
	}
	cmHandler := &codemode.Handler{Client: codemode.NewClient(workerdURL)}
	s.AddTool(codemode.NewSearchTool(), cmHandler.Search)

	return s
}
