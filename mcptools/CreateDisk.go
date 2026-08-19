package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const createDiskInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "name": {
          "description": "User-scoped disk name. Lowercase alphanumeric plus dash, 1 to 63 chars.",
          "pattern": "^[a-z0-9][a-z0-9-]{0,62}$",
          "example": "my-data",
          "type": "string"
        },
        "kind": {
          "description": "Disk backend kind. Currently only s3 is supported.",
          "enum": ["s3"],
          "type": "string"
        },
        "config": {
          "properties": {
            "bucket": { "type": "string", "example": "my-data-bucket" },
            "endpoint": { "type": "string", "example": "https://s3.amazonaws.com" },
            "region": { "type": "string", "example": "us-east-1" },
            "use_path_style": { "type": "boolean" }
          },
          "required": ["bucket", "endpoint"],
          "type": "object"
        },
        "credentials": {
          "properties": {
            "access_key": { "type": "string" },
            "secret_key": { "format": "password", "type": "string" }
          },
          "required": ["access_key", "secret_key"],
          "type": "object"
        }
      },
      "required": ["name", "kind", "config", "credentials"],
      "type": "object"
    }
  },
  "required": ["body"],
  "type": "object"
}`

func NewCreateDiskMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"CreateDisk",
		"Register an S3-compatible bucket as a sandbox disk.",
		[]byte(createDiskInputSchema),
	)
}
