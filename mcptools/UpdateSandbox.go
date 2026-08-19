package mcptools

import "github.com/mark3labs/mcp-go/mcp"

// Input Schema for the UpdateSandbox tool
const updateSandboxInputSchema = `{
  "properties": {
    "id": {
      "description": "Sandbox identifier.",
      "type": "string"
    },
    "body": {
      "description": "Partial sandbox update. Omitted fields are left unchanged.",
      "properties": {
        "ingress_enabled": {
          "description": "Enable or disable HTTP ingress for this sandbox.",
          "type": "boolean"
        },
        "auto_pause_after_seconds": {
          "description": "Idle timeout in seconds. Valid range is 60 to 86400.",
          "maximum": 86400,
          "minimum": 60,
          "type": ["integer", "null"]
        },
        "disable_auto_pause": {
          "description": "Set true to clear auto_pause_after_seconds.",
          "type": "boolean"
        }
      },
      "type": "object"
    }
  },
  "required": ["id", "body"],
  "type": "object"
}`

// NewUpdateSandboxMCPTool creates the MCP Tool instance for UpdateSandbox
func NewUpdateSandboxMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"UpdateSandbox",
		"Partially update sandbox settings such as ingress and auto-pause.",
		[]byte(updateSandboxInputSchema),
	)
}
