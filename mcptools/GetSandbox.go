package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const getSandboxInputSchema = `{
  "properties": {
    "id": {
      "description": "Sandbox identifier.",
      "type": "string"
    }
  },
  "required": ["id"],
  "type": "object"
}`

func NewGetSandboxMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetSandbox",
		"Get one sandbox VM by id.",
		[]byte(getSandboxInputSchema),
	)
}
