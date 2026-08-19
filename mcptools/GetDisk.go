package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const getDiskInputSchema = `{
  "properties": {
    "id_or_name": {
      "description": "Disk id, such as disk_<ulid>, or user-scoped disk name.",
      "type": "string"
    }
  },
  "required": ["id_or_name"],
  "type": "object"
}`

func NewGetDiskMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetDisk",
		"Get sandbox disk metadata by id or name. Credentials are never returned by the API.",
		[]byte(getDiskInputSchema),
	)
}
