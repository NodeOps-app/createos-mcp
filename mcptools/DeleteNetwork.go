package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const deleteNetworkInputSchema = `{
  "properties": {
    "id": {
      "description": "Network name or net-<ulid> id.",
      "type": "string"
    }
  },
  "required": ["id"],
  "type": "object"
}`

func NewDeleteNetworkMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"DeleteNetwork",
		"Delete a private sandbox network. The network must have no active members.",
		[]byte(deleteNetworkInputSchema),
	)
}
