package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const listNetworksInputSchema = `{
  "properties": {
    "limit": {
      "description": "Maximum number of networks to return.",
      "maximum": 500,
      "type": "integer"
    },
    "offset": {
      "description": "Pagination offset.",
      "minimum": 0,
      "type": "integer"
    }
  },
  "type": "object"
}`

func NewListNetworksMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListNetworks",
		"List private sandbox networks owned by the caller.",
		[]byte(listNetworksInputSchema),
	)
}
