package mcptools

import "github.com/mark3labs/mcp-go/mcp"

// Input Schema for the DeleteSandbox tool
const deleteSandboxInputSchema = `{
  "properties": {
    "id": {
      "description": "Sandbox identifier.",
      "type": "string"
    }
  },
  "required": ["id"],
  "type": "object"
}`

// NewDeleteSandboxMCPTool creates the MCP Tool instance for DeleteSandbox
func NewDeleteSandboxMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"DeleteSandbox",
		"Destroy a sandbox VM. The backend treats deleting an already terminal sandbox idempotently.",
		[]byte(deleteSandboxInputSchema),
	)
}
