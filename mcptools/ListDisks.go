package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const listDisksInputSchema = `{
  "properties": {
    "limit": {
      "description": "Maximum number of disks to return.",
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

func NewListDisksMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListDisks",
		"List sandbox disks owned by the caller.",
		[]byte(listDisksInputSchema),
	)
}
