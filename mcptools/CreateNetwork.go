package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const createNetworkInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "name": {
          "description": "User-facing network name.",
          "example": "backend",
          "type": "string"
        }
      },
      "required": ["name"],
      "type": "object"
    }
  },
  "required": ["body"],
  "type": "object"
}`

func NewCreateNetworkMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"CreateNetwork",
		"Create a private sandbox network.",
		[]byte(createNetworkInputSchema),
	)
}
