package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const detachSandboxNetworkInputSchema = `{
  "properties": {
    "id": {
      "description": "Sandbox identifier.",
      "type": "string"
    },
    "network": {
      "description": "Network name or net-<ulid> id.",
      "type": "string"
    }
  },
  "required": ["id", "network"],
  "type": "object"
}`

func NewDetachSandboxNetworkMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"DetachSandboxNetwork",
		"Detach a sandbox from a private network.",
		[]byte(detachSandboxNetworkInputSchema),
	)
}
