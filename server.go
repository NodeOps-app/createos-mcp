package main

import (
	"os"

	"github.com/NodeOps-app/createos-mcp/codemode"
	handler "github.com/NodeOps-app/createos-mcp/handlers"
	"github.com/NodeOps-app/createos-mcp/mcptools"
	"github.com/mark3labs/mcp-go/server"
)

// NewMCPServer creates and returns an MCP server with all tools registered.
//
// v2: 7 native fast-path tools + 3 code-mode tools (search, execute, pollJob).
// The previous 87 long-tail tools are reachable via search/execute against the
// CreateOS OpenAPI spec, served by the workerd sidecar.
func NewMCPServer() *server.MCPServer {
	s := server.NewMCPServer(
		"CreateOS MCP (Code Mode v2)",
		"2.0.0",
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(false, false),
		server.WithPromptCapabilities(false),
		server.WithLogging(),
	)

	// Native fast-path tools (v2: 7 retained).
	s.AddTool(mcptools.NewGetQuotasMCPTool(), handler.GetQuotasHandler)
	s.AddTool(mcptools.NewGetSupportedProjectTypesMCPTool(), handler.GetSupportedProjectTypesHandler)
	s.AddTool(mcptools.NewCheckProjectUniqueNameMCPTool(), handler.CheckProjectUniqueNameHandler)
	s.AddTool(mcptools.NewCreateProjectMCPTool(), handler.CreateProjectHandler)
	s.AddTool(mcptools.NewUploadDeploymentBase64FilesMCPTool(), handler.UploadDeploymentBase64FilesHandler)
	s.AddTool(mcptools.NewGetDeploymentMCPTool(), handler.GetDeploymentHandler)
	s.AddTool(mcptools.NewCancelDeploymentMCPTool(), handler.CancelDeploymentHandler)

	// Code Mode tools.
	workerdURL := os.Getenv("WORKERD_URL")
	if workerdURL == "" {
		workerdURL = "http://127.0.0.1:8787"
	}
	cmHandler := &codemode.Handler{Client: codemode.NewClient(workerdURL)}
	s.AddTool(codemode.NewSearchTool(), cmHandler.Search)
	s.AddTool(codemode.NewExecuteTool(), cmHandler.Execute)
	s.AddTool(codemode.NewPollJobTool(), cmHandler.PollJob)

	// Discovery surface so clients/agents can pull deeper docs on demand.
	introRes, introHandler := codemode.NewIntroResource()
	s.AddResource(introRes, server.ResourceHandlerFunc(introHandler))
	apiRes, apiHandler := codemode.NewAPIShapeResource()
	s.AddResource(apiRes, server.ResourceHandlerFunc(apiHandler))

	deployPrompt, deployHandler := codemode.NewDeployExamplePrompt()
	s.AddPrompt(deployPrompt, server.PromptHandlerFunc(deployHandler))
	discoveryPrompt, discoveryHandler := codemode.NewAPIDiscoveryPrompt()
	s.AddPrompt(discoveryPrompt, server.PromptHandlerFunc(discoveryHandler))

	return s
}
