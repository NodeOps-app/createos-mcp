package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the UpdateApp tool
const updateAppInputSchema = `{
  "properties": {
    "app_id": {
      "description": "App unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "body": {
      "properties": {
        "color": {
          "description": "Hex color code for the app",
          "example": "#3B82F6",
          "pattern": "^#[0-9A-Fa-f]{6}$",
          "type": "string"
        },
        "description": {
          "description": "Description of the app",
          "example": "Updated description",
          "maxLength": 2048,
          "minLength": 1,
          "type": [
            "string",
            "null"
          ]
        },
        "name": {
          "description": "Name of the app",
          "example": "My Updated Application",
          "maxLength": 255,
          "minLength": 1,
          "type": "string"
        }
      },
      "required": [
        "name"
      ],
      "type": "object"
    }
  },
  "required": [
    "app_id",
    "body"
  ],
  "type": "object"
}`

// Response Template for the UpdateApp tool (Status: 200, Content-Type: application/json)
const UpdateAppResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> App updated successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: string):
      - Example: 'app updated successfully'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the UpdateApp tool (Status: 400, Content-Type: application/json)
const UpdateAppResponseTemplate_B = `# API Response Information

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

// Response Template for the UpdateApp tool (Status: 401, Content-Type: application/json)
const UpdateAppResponseTemplate_C = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 401

**Content-Type:** application/json

> Unauthorized - authentication required

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the UpdateApp tool (Status: 403, Content-Type: application/json)
const UpdateAppResponseTemplate_D = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 403

**Content-Type:** application/json

> Forbidden - user does not have permission

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the UpdateApp tool (Status: 404, Content-Type: application/json)
const UpdateAppResponseTemplate_E = `# API Response Information

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

// Response Template for the UpdateApp tool (Status: 500, Content-Type: application/json)
const UpdateAppResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 500

**Content-Type:** application/json

> Internal server error

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// NewUpdateAppMCPTool creates the MCP Tool instance for UpdateApp
func NewUpdateAppMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"UpdateApp",
		"Update app - Updates an app's name, description, and color",
		[]byte(updateAppInputSchema),
	)
}

// UpdateAppHandler is the handler function for the UpdateApp tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func UpdateAppHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "UpdateApp")
}
