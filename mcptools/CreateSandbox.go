package mcptools

import "github.com/mark3labs/mcp-go/mcp"

// Input Schema for the CreateSandbox tool
const createSandboxInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "shape": {
          "description": "Required sandbox shape. Use the sandbox API's GET /v1/shapes endpoint to discover available values.",
          "example": "s-1vcpu-256mb",
          "type": "string"
        },
        "rootfs": {
          "description": "Root filesystem catalog name. Empty means host default.",
          "example": "devbox:1",
          "type": "string"
        },
        "name": {
          "description": "User-facing sandbox name, unique per user among non-terminal sandboxes. Omit to auto-generate.",
          "example": "brave-otter",
          "type": "string"
        },
        "networks": {
          "description": "Private networks to join at create time. Each entry is {\"id\":\"<name-or-network-id>\"}.",
          "items": {
            "properties": {
              "id": { "type": "string" }
            },
            "required": ["id"],
            "type": "object"
          },
          "type": "array"
        },
        "disk_mib": {
          "description": "Disk size in MiB. 0 or omitted uses the shape default.",
          "format": "int64",
          "type": "integer"
        },
        "egress": {
          "description": "Network egress allowlist entries. Use host[:port], ip[:port], cidr[:port], or *.",
          "items": { "type": "string" },
          "type": "array"
        },
        "envs": {
          "additionalProperties": { "type": "string" },
          "description": "Environment variables exported into every exec invocation. Keys must be declared here before exec overrides can use them.",
          "type": "object"
        },
        "ssh_pubkeys": {
          "description": "OpenSSH public keys authorized for SSH gateway tunnel/shell access.",
          "items": { "type": "string" },
          "type": "array"
        },
        "host_id": {
          "description": "Optional host pin. Empty lets the scheduler choose.",
          "type": "string"
        },
        "region": {
          "description": "Optional placement region.",
          "example": "us",
          "type": "string"
        },
        "ingress_enabled": {
          "description": "Whether HTTP ingress is enabled for this sandbox.",
          "type": "boolean"
        },
        "auto_pause_after_seconds": {
          "description": "Idle timeout in seconds. Valid range is 60 to 86400. Null or omitted disables auto-pause.",
          "maximum": 86400,
          "minimum": 60,
          "type": ["integer", "null"]
        },
        "disks": {
          "description": "S3 disks to mount at create time.",
          "items": {
            "properties": {
              "disk_id": { "type": "string" },
              "mount_path": { "type": "string" }
            },
            "required": ["disk_id", "mount_path"],
            "type": "object"
          },
          "type": "array"
        }
      },
      "required": ["shape"],
      "type": "object"
    }
  },
  "required": ["body"],
  "type": "object"
}`

// NewCreateSandboxMCPTool creates the MCP Tool instance for CreateSandbox
func NewCreateSandboxMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"CreateSandbox",
		"Create a sandbox VM in the sandbox control plane.",
		[]byte(createSandboxInputSchema),
	)
}
