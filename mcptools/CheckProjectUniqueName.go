package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the CheckProjectUniqueName tool
const checkProjectUniqueNameInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "uniqueName": {
          "example": "my-awesome-app",
          "maxLength": 32,
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

// Response Template for the CheckProjectUniqueName tool (Status: 200, Content-Type: application/json)
const CheckProjectUniqueNameResponseTemplate_A = `# API Response Information

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

// Response Template for the CheckProjectUniqueName tool (Status: 400, Content-Type: application/json)
const CheckProjectUniqueNameResponseTemplate_B = `# API Response Information

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

// Response Template for the CheckProjectUniqueName tool (Status: 401, Content-Type: application/json)
const CheckProjectUniqueNameResponseTemplate_C = `# API Response Information

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

// Response Template for the CheckProjectUniqueName tool (Status: 500, Content-Type: application/json)
const CheckProjectUniqueNameResponseTemplate_D = `# API Response Information

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

// NewCheckProjectUniqueNameMCPTool creates the MCP Tool instance for CheckProjectUniqueName
func NewCheckProjectUniqueNameMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"CheckProjectUniqueName",
		"Check if project unique name is available - Checks if a project unique name is available for the authenticated user",
		[]byte(checkProjectUniqueNameInputSchema),
	)
}

// CheckProjectUniqueNameHandler is the handler function for the CheckProjectUniqueName tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func CheckProjectUniqueNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "CheckProjectUniqueName")
}
