package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the CheckAPIKeyUniqueName tool
const checkAPIKeyUniqueNameInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "uniqueName": {
          "description": "API key name to check for availability",
          "example": "my-api-key",
          "maxLength": 48,
          "minLength": 4,
          "pattern": "^[a-zA-Z0-9-]+$",
          "type": "string"
        }
      },
      "required": [
        "uniqueName"
      ],
      "type": "object"
    }
  },
  "required": [
    "body"
  ],
  "type": "object"
}`

// Response Template for the CheckAPIKeyUniqueName tool (Status: 200, Content-Type: application/json)
const CheckAPIKeyUniqueNameResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Availability check completed

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **isAvailable** (Type: boolean):
        - Example: 'true'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the CheckAPIKeyUniqueName tool (Status: 400, Content-Type: application/json)
const CheckAPIKeyUniqueNameResponseTemplate_B = `# API Response Information

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

// Response Template for the CheckAPIKeyUniqueName tool (Status: 401, Content-Type: application/json)
const CheckAPIKeyUniqueNameResponseTemplate_C = `# API Response Information

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

// Response Template for the CheckAPIKeyUniqueName tool (Status: 500, Content-Type: application/json)
const CheckAPIKeyUniqueNameResponseTemplate_D = `# API Response Information

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

// NewCheckAPIKeyUniqueNameMCPTool creates the MCP Tool instance for CheckAPIKeyUniqueName
func NewCheckAPIKeyUniqueNameMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"CheckAPIKeyUniqueName",
		"Check if API key name is available - Checks if an API key name is available for the authenticated user",
		[]byte(checkAPIKeyUniqueNameInputSchema),
	)
}

// CheckAPIKeyUniqueNameHandler is the handler function for the CheckAPIKeyUniqueName tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func CheckAPIKeyUniqueNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "CheckAPIKeyUniqueName")
}
