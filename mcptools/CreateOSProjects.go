package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const createOSProjectListInputSchema = `{
  "type": "object",
  "properties": {
    "app": {
      "type": "string",
      "description": "Filter by app ID (UUID). Use to narrow results."
    },
    "name": {
      "type": "string",
      "description": "Filter by project name (partial match)."
    },
    "status": {
      "type": "string",
      "enum": ["active", "deleting", "deleted"],
      "description": "Project status filter."
    },
    "type": {
      "type": "string",
      "enum": ["vcs", "image", "upload"],
      "description": "Project type filter."
    },
    "limit": {
      "type": "integer",
      "minimum": 1,
      "maximum": 50,
      "default": 20,
      "description": "Maximum number of items per page."
    },
    "cursor": {
      "type": "string",
      "description": "Pagination cursor from a previous response."
    }
  }
}`

const createOSProjectGetInputSchema = `{
  "type": "object",
  "required": ["project_id"],
  "properties": {
    "project_id": {
      "type": "string",
      "description": "Project ID (UUID)."
    }
  }
}`

const createOSProjectCreateVCSInputSchema = `{
  "type": "object",
  "required": ["unique_name", "display_name", "vcs_installation_id", "vcs_repo_id"],
  "properties": {
    "app_id": {"type": "string", "description": "Optional app ID to associate."},
    "unique_name": {"type": "string", "description": "Unique name (lowercase, hyphenated)."},
    "display_name": {"type": "string", "description": "Human readable name."},
    "description": {"type": "string"},
    "enabled_security_scan": {"type": "boolean", "default": false},
    "vcs_name": {"type": "string", "enum": ["github"], "default": "github"},
    "vcs_installation_id": {"type": "string", "description": "GitHub App installation ID."},
    "vcs_repo_id": {"type": "string", "description": "GitHub repository ID."},
    "build_command": {"type": "string"},
    "build_dir": {"type": "string"},
    "build_flag": {"type": "string"},
    "build_flags": {"type": "object", "additionalProperties": {"type": "string"}},
    "build_vars": {"type": "object", "additionalProperties": {"type": "string"}},
    "directory_path": {"type": "string"},
    "framework": {"type": "string"},
    "has_dockerfile": {"type": "boolean"},
    "ignore_branches": {"type": "array", "items": {"type": "string"}},
    "install_command": {"type": "string"},
    "port": {"type": "integer", "minimum": 1, "maximum": 65535},
    "run_command": {"type": "string"},
    "run_envs": {"type": "object", "additionalProperties": {"type": "string"}},
    "run_flag": {"type": "string"},
    "runtime": {"type": "string"},
    "use_build_ai": {"type": "boolean"}
  }
}`

const createOSProjectCreateImageInputSchema = `{
  "type": "object",
  "required": ["unique_name", "display_name", "port"],
  "properties": {
    "app_id": {"type": "string", "description": "Optional app ID to associate."},
    "unique_name": {"type": "string"},
    "display_name": {"type": "string"},
    "description": {"type": "string"},
    "enabled_security_scan": {"type": "boolean", "default": false},
    "port": {"type": "integer", "minimum": 1, "maximum": 65535},
    "run_envs": {"type": "object", "additionalProperties": {"type": "string"}}
  }
}`

const createOSProjectCreateUploadInputSchema = `{
  "type": "object",
  "required": ["unique_name", "display_name"],
  "properties": {
    "app_id": {"type": "string", "description": "Optional app ID to associate."},
    "unique_name": {"type": "string"},
    "display_name": {"type": "string"},
    "description": {"type": "string"},
    "enabled_security_scan": {"type": "boolean", "default": false},
    "build_command": {"type": "string"},
    "build_dir": {"type": "string"},
    "build_flag": {"type": "string"},
    "build_vars": {"type": "object", "additionalProperties": {"type": "string"}},
    "directory_path": {"type": "string"},
    "framework": {"type": "string"},
    "has_dockerfile": {"type": "boolean"},
    "install_command": {"type": "string"},
    "port": {"type": "integer", "minimum": 1, "maximum": 65535},
    "run_command": {"type": "string"},
    "run_envs": {"type": "object", "additionalProperties": {"type": "string"}},
    "run_flag": {"type": "string"},
    "runtime": {"type": "string"},
    "use_build_ai": {"type": "boolean"}
  }
}`

func NewCreateOSProjectListMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"createos.project.list",
		"List projects for the authenticated user. Use when the user asks for project inventory or wants a filtered list.",
		[]byte(createOSProjectListInputSchema),
	)
}

func NewCreateOSProjectGetMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"createos.project.get",
		"Get a single project by ID. Use when the user references a specific project ID.",
		[]byte(createOSProjectGetInputSchema),
	)
}

func NewCreateOSProjectCreateVCSMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"createos.project.create_vcs",
		"Create a VCS-backed project in one step. Use when the user wants a GitHub-backed project.",
		[]byte(createOSProjectCreateVCSInputSchema),
	)
}

func NewCreateOSProjectCreateImageMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"createos.project.create_image",
		"Create an image-based project (prebuilt container). Use when the user has an image and port.",
		[]byte(createOSProjectCreateImageInputSchema),
	)
}

func NewCreateOSProjectCreateUploadMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"createos.project.create_upload",
		"Create an upload-based project. Use when the user wants to deploy uploaded files.",
		[]byte(createOSProjectCreateUploadInputSchema),
	)
}
