package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const createOSDeploymentListInputSchema = `{
  "type": "object",
  "required": ["project_id"],
  "properties": {
    "project_id": {"type": "string", "description": "Project ID (UUID)."},
    "limit": {"type": "integer", "minimum": 1, "maximum": 50, "default": 20},
    "cursor": {"type": "string", "description": "Pagination cursor from a previous response."}
  }
}`

const createOSDeploymentGetInputSchema = `{
  "type": "object",
  "required": ["deployment_id"],
  "properties": {
    "deployment_id": {"type": "string", "description": "Deployment ID (UUID)."}
  }
}`

const createOSDeploymentTriggerLatestInputSchema = `{
  "type": "object",
  "required": ["project_id"],
  "properties": {
    "project_id": {"type": "string", "description": "Project ID (UUID)."}
  }
}`

const createOSDeploymentCancelInputSchema = `{
  "type": "object",
  "required": ["deployment_id"],
  "properties": {
    "deployment_id": {"type": "string", "description": "Deployment ID (UUID)."}
  }
}`

func NewCreateOSDeploymentListMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"createos.deployment.list",
		"List deployments for a project. Use when the user asks about recent deployments.",
		[]byte(createOSDeploymentListInputSchema),
	)
}

func NewCreateOSDeploymentGetMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"createos.deployment.get",
		"Get deployment details by ID. Use when a specific deployment ID is referenced.",
		[]byte(createOSDeploymentGetInputSchema),
	)
}

func NewCreateOSDeploymentTriggerLatestMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"createos.deployment.trigger_latest",
		"Trigger the latest deployment for a project. Use when the user asks to redeploy.",
		[]byte(createOSDeploymentTriggerLatestInputSchema),
	)
}

func NewCreateOSDeploymentCancelMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"createos.deployment.cancel",
		"Cancel an in-progress deployment. Use when the user wants to stop a deployment.",
		[]byte(createOSDeploymentCancelInputSchema),
	)
}
