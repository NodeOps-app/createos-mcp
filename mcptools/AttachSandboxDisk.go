package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const attachSandboxDiskInputSchema = `{
  "properties": {
    "id": {
      "description": "Sandbox identifier.",
      "type": "string"
    },
    "body": {
      "properties": {
        "disk_id": {
          "description": "Disk id or user-scoped disk name.",
          "example": "my-data",
          "type": "string"
        },
        "mount_path": {
          "description": "Absolute mount path inside the sandbox guest.",
          "example": "/mnt/data",
          "type": "string"
        },
        "sub_path": {
          "description": "Optional bucket prefix to mount.",
          "example": "team-a/",
          "type": "string"
        }
      },
      "required": ["disk_id", "mount_path"],
      "type": "object"
    }
  },
  "required": ["id", "body"],
  "type": "object"
}`

func NewAttachSandboxDiskMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"AttachSandboxDisk",
		"Live-attach a disk to a running sandbox.",
		[]byte(attachSandboxDiskInputSchema),
	)
}
