package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the UpdateCronjob tool
const updateCronjobInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "name": {
          "description": "Name of the cronjob",
          "example": "health-check-updated",
          "maxLength": 255,
          "minLength": 1,
          "type": "string"
        },
        "schedule": {
          "description": "Cron schedule expression",
          "example": "0 * * * *",
          "maxLength": 256,
          "type": "string"
        },
        "settings": {
          "description": "Cronjob type-specific settings",
          "oneOf": [
            {
              "properties": {
                "body": {
                  "description": "Request body",
                  "maxLength": 65536,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "headers": {
                  "additionalProperties": {
                    "type": "string"
                  },
                  "description": "HTTP headers to include in the request",
                  "example": {
                    "Authorization": "Bearer token"
                  },
                  "type": "object"
                },
                "method": {
                  "description": "HTTP method",
                  "enum": [
                    "GET",
                    "POST",
                    "PUT",
                    "DELETE",
                    "PATCH",
                    "HEAD"
                  ],
                  "example": "GET",
                  "type": "string"
                },
                "path": {
                  "description": "Request path (must start with /)",
                  "example": "/api/health",
                  "maxLength": 2048,
                  "type": "string"
                },
                "timeoutInSeconds": {
                  "description": "Request timeout in seconds",
                  "example": 5,
                  "maximum": 5,
                  "minimum": 1,
                  "type": "integer"
                }
              },
              "required": [
                "path",
                "method"
              ],
              "type": "object"
            }
          ]
        }
      },
      "required": [
        "name",
        "schedule"
      ],
      "type": "object"
    },
    "cronjob_id": {
      "description": "Cronjob unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "project_id": {
      "description": "Project unique identifier",
      "format": "uuid",
      "type": "string"
    }
  },
  "required": [
    "project_id",
    "cronjob_id",
    "body"
  ],
  "type": "object"
}`

// Response Template for the UpdateCronjob tool (Status: 200, Content-Type: application/json)
const UpdateCronjobResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Cronjob updated successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: string):
      - Example: 'cronjob updated'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the UpdateCronjob tool (Status: 400, Content-Type: application/json)
const UpdateCronjobResponseTemplate_B = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 400

**Content-Type:** application/json

> Bad request - validation error

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the UpdateCronjob tool (Status: 401, Content-Type: application/json)
const UpdateCronjobResponseTemplate_C = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 401

**Content-Type:** application/json

> Unauthorized - authentication required

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the UpdateCronjob tool (Status: 403, Content-Type: application/json)
const UpdateCronjobResponseTemplate_D = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 403

**Content-Type:** application/json

> Forbidden - user does not have permission

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the UpdateCronjob tool (Status: 404, Content-Type: application/json)
const UpdateCronjobResponseTemplate_E = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 404

**Content-Type:** application/json

> Not found - resource does not exist

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the UpdateCronjob tool (Status: 500, Content-Type: application/json)
const UpdateCronjobResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 500

**Content-Type:** application/json

> Internal server error

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// NewUpdateCronjobMCPTool creates the MCP Tool instance for UpdateCronjob
func NewUpdateCronjobMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"UpdateCronjob",
		"Update a cronjob - Updates the name, schedule, and settings of a cronjob.",
		[]byte(updateCronjobInputSchema),
	)
}

// UpdateCronjobHandler is the handler function for the UpdateCronjob tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func UpdateCronjobHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "UpdateCronjob")
}
