package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the UpdateAPIKey tool
const updateAPIKeyInputSchema = `{
  "properties": {
    "api_key_id": {
      "description": "API key unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "body": {
      "properties": {
        "description": {
          "example": "Updated API key description",
          "maxLength": 2048,
          "minLength": 4,
          "pattern": "^[a-zA-Z0-9 _-]+$",
          "type": [
            "string",
            "null"
          ]
        },
        "name": {
          "example": "updated-api-key",
          "maxLength": 48,
          "minLength": 4,
          "pattern": "^[a-zA-Z0-9-]+$",
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
    "api_key_id",
    "body"
  ],
  "type": "object"
}`

// Response Template for the UpdateAPIKey tool (Status: 200, Content-Type: application/json)
const UpdateAPIKeyResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> API key updated successfully

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'true'
  - **data** (Type: string):
      - Example: 'api key updated successfully'
`

// Response Template for the UpdateAPIKey tool (Status: 400, Content-Type: application/json)
const UpdateAPIKeyResponseTemplate_B = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 400

**Content-Type:** application/json

> Bad request - validation error

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the UpdateAPIKey tool (Status: 401, Content-Type: application/json)
const UpdateAPIKeyResponseTemplate_C = `# API Response Information

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

// Response Template for the UpdateAPIKey tool (Status: 404, Content-Type: application/json)
const UpdateAPIKeyResponseTemplate_D = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 404

**Content-Type:** application/json

> Not found - resource does not exist

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the UpdateAPIKey tool (Status: 500, Content-Type: application/json)
const UpdateAPIKeyResponseTemplate_E = `# API Response Information

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

// NewUpdateAPIKeyMCPTool creates the MCP Tool instance for UpdateAPIKey
func NewUpdateAPIKeyMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"UpdateAPIKey",
		"Update API key - Updates the name and description of an API key",
		[]byte(updateAPIKeyInputSchema),
	)
}

// UpdateAPIKeyHandler is the handler function for the UpdateAPIKey tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func UpdateAPIKeyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "UpdateAPIKey")
}
