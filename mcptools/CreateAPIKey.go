package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the CreateAPIKey tool
const createAPIKeyInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "description": {
          "example": "API key for production use",
          "maxLength": 2048,
          "minLength": 4,
          "pattern": "^[a-zA-Z0-9 _-]+$",
          "type": [
            "string",
            "null"
          ]
        },
        "expiryAt": {
          "description": "Expiration date and time (must be in the future)",
          "example": "2025-12-31T23:59:59Z",
          "format": "date-time",
          "type": "string"
        },
        "name": {
          "example": "my-api-key",
          "maxLength": 48,
          "minLength": 4,
          "pattern": "^[a-zA-Z0-9-]+$",
          "type": "string"
        }
      },
      "required": [
        "name",
        "expiryAt"
      ],
      "type": "object"
    }
  },
  "required": [
    "body"
  ],
  "type": "object"
}`

// Response Template for the CreateAPIKey tool (Status: 201, Content-Type: application/json)
const CreateAPIKeyResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 201

**Content-Type:** application/json

> API key created successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **id** (Type: string, uuid):
    - **secret**: Full API key secret (only shown once on creation) (Type: string):
        - Example: 'skp_abc123def456...'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the CreateAPIKey tool (Status: 400, Content-Type: application/json)
const CreateAPIKeyResponseTemplate_B = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 400

**Content-Type:** application/json

> Bad request - validation error or limit reached

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the CreateAPIKey tool (Status: 401, Content-Type: application/json)
const CreateAPIKeyResponseTemplate_C = `# API Response Information

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

// Response Template for the CreateAPIKey tool (Status: 500, Content-Type: application/json)
const CreateAPIKeyResponseTemplate_D = `# API Response Information

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

// NewCreateAPIKeyMCPTool creates the MCP Tool instance for CreateAPIKey
func NewCreateAPIKeyMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"CreateAPIKey",
		"Create API key - Creates a new API key for the authenticated user",
		[]byte(createAPIKeyInputSchema),
	)
}

// CreateAPIKeyHandler is the handler function for the CreateAPIKey tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func CreateAPIKeyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "CreateAPIKey")
}
