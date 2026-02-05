package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const createOSDomainListInputSchema = `{
  "type": "object",
  "required": ["project_id"],
  "properties": {
    "project_id": {"type": "string", "description": "Project ID (UUID)."}
  }
}`

const createOSGithubRepositoriesListInputSchema = `{
  "type": "object",
  "required": ["installation_id"],
  "properties": {
    "installation_id": {"type": "string", "description": "GitHub App installation ID."}
  }
}`

const createOSGithubRepositoryBranchesListInputSchema = `{
  "type": "object",
  "required": ["installation_id", "repository"],
  "properties": {
    "installation_id": {"type": "string", "description": "GitHub App installation ID."},
    "repository": {"type": "string", "description": "Repository full name (org/repo)."},
    "limit": {"type": "integer", "minimum": 1, "maximum": 50, "default": 20},
    "cursor": {"type": "string", "description": "Pagination cursor from a previous response."}
  }
}`

func NewCreateOSDomainListMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"createos.domain.list",
		"List domains attached to a project. Use when the user asks about custom domains.",
		[]byte(createOSDomainListInputSchema),
	)
}

func NewCreateOSGithubRepositoriesListMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"createos.github.repositories.list",
		"List GitHub repositories for a connected installation. Use when selecting a repo.",
		[]byte(createOSGithubRepositoriesListInputSchema),
	)
}

func NewCreateOSGithubRepositoryBranchesListMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"createos.github.repository.branches.list",
		"List branches for a GitHub repository. Use when the user needs a branch name.",
		[]byte(createOSGithubRepositoryBranchesListInputSchema),
	)
}
