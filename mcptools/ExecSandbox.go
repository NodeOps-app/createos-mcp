package mcptools

import "github.com/mark3labs/mcp-go/mcp"

// Input Schema for the ExecSandbox tool
const execSandboxInputSchema = `{
  "properties": {
    "id": {
      "description": "Sandbox identifier.",
      "type": "string"
    },
    "body": {
      "properties": {
        "cmd": {
          "description": "Program to execute inside the VM, absolute or PATH-resolved.",
          "example": "ls",
          "type": "string"
        },
        "args": {
          "description": "Arguments passed directly to the executable without shell parsing.",
          "items": { "type": "string" },
          "type": "array"
        },
        "stdin": {
          "description": "Optional stdin passed to the process.",
          "type": "string"
        },
        "env": {
          "additionalProperties": { "type": "string" },
          "description": "Per-exec environment overrides. Keys must have been declared in the sandbox envs at create time.",
          "type": "object"
        },
        "stream": {
          "description": "Streaming exec is not supported by this MCP tool yet. Leave false or omit.",
          "type": "boolean"
        }
      },
      "required": ["cmd"],
      "type": "object"
    }
  },
  "required": ["id", "body"],
  "type": "object"
}`

// NewExecSandboxMCPTool creates the MCP Tool instance for ExecSandbox
func NewExecSandboxMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ExecSandbox",
		"Run a buffered command inside a sandbox VM.",
		[]byte(execSandboxInputSchema),
	)
}
