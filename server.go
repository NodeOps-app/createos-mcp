package main

import (
	handler "github.com/NodeOps-app/autogen-backend-v2-mcp/handlers"
	"github.com/NodeOps-app/autogen-backend-v2-mcp/mcptools"
	"github.com/mark3labs/mcp-go/server"
)

func NewMCPServer() *server.MCPServer {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Autogen MCP Server",
		"0.0.1",
		server.WithToolCapabilities(true),
		server.WithLogging(),
	)

	// Register all tools
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

	return s
}
