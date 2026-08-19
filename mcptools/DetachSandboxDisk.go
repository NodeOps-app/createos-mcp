package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const detachSandboxDiskInputSchema = `{
  "properties": {
    "id": {
      "description": "Sandbox identifier.",
      "type": "string"
    },
    "disk_id": {
      "description": "Disk id, always disk_<ulid> for detach.",
      "type": "string"
    },
    "mount_path": {
      "description": "Absolute mount path for the attachment to detach.",
      "example": "/mnt/data",
      "type": "string"
    }
  },
  "required": ["id", "disk_id", "mount_path"],
  "type": "object"
}`

func NewDetachSandboxDiskMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"DetachSandboxDisk",
		"Detach one disk mount from a sandbox. Bucket contents are not touched.",
		[]byte(detachSandboxDiskInputSchema),
	)
}
