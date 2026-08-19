package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const listSandboxDisksInputSchema = `{
  "properties": {
    "id": {
      "description": "Sandbox identifier.",
      "type": "string"
    }
  },
  "required": ["id"],
  "type": "object"
}`

func NewListSandboxDisksMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListSandboxDisks",
		"List disks attached to a sandbox, including mount status.",
		[]byte(listSandboxDisksInputSchema),
	)
}
