package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const attachSandboxNetworkInputSchema = `{
  "properties": {
    "id": {
      "description": "Sandbox identifier.",
      "type": "string"
    },
    "body": {
      "properties": {
        "id": {
          "description": "Network name or net-<ulid> id.",
          "type": "string"
        }
      },
      "required": ["id"],
      "type": "object"
    }
  },
  "required": ["id", "body"],
  "type": "object"
}`

func NewAttachSandboxNetworkMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"AttachSandboxNetwork",
		"Attach a sandbox to a private network.",
		[]byte(attachSandboxNetworkInputSchema),
	)
}
