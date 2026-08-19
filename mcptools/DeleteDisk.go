package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const deleteDiskInputSchema = `{
  "properties": {
    "id_or_name": {
      "description": "Disk id, such as disk_<ulid>, or user-scoped disk name.",
      "type": "string"
    }
  },
  "required": ["id_or_name"],
  "type": "object"
}`

func NewDeleteDiskMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"DeleteDisk",
		"Soft-delete a sandbox disk. Bucket contents are not touched.",
		[]byte(deleteDiskInputSchema),
	)
}
