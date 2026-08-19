package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const listSandboxesInputSchema = `{
  "properties": {
    "limit": {
      "description": "Maximum number of sandboxes to return.",
      "maximum": 500,
      "type": "integer"
    },
    "offset": {
      "description": "Pagination offset.",
      "minimum": 0,
      "type": "integer"
    },
    "status": {
      "description": "Optional status filter. Empty or omitted returns any status.",
      "enum": ["running", "creating", "destroyed", "failed"],
      "type": "string"
    }
  },
  "type": "object"
}`

func NewListSandboxesMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListSandboxes",
		"List sandbox VMs owned by the caller.",
		[]byte(listSandboxesInputSchema),
	)
}
