package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const getNetworkInputSchema = `{
  "properties": {
    "id": {
      "description": "Network name or net-<ulid> id.",
      "type": "string"
    }
  },
  "required": ["id"],
  "type": "object"
}`

func NewGetNetworkMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetNetwork",
		"Get one sandbox network, including members.",
		[]byte(getNetworkInputSchema),
	)
}
