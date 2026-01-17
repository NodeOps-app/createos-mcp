package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the CheckProjectEnvironmentUniqueName tool
const checkProjectEnvironmentUniqueNameInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "uniqueName": {
          "example": "production",
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
    },
    "project_id": {
      "description": "Project unique identifier",
      "format": "uuid",
      "type": "string"
    }
  },
  "required": [
    "project_id",
    "body"
  ],
  "type": "object"
}`

// Response Template for the CheckProjectEnvironmentUniqueName tool (Status: 200, Content-Type: application/json)
const CheckProjectEnvironmentUniqueNameResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Availability check completed

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'true'
  - **data** (Type: object):
    - **isAvailable** (Type: boolean):
        - Example: 'true'
`

// Response Template for the CheckProjectEnvironmentUniqueName tool (Status: 400, Content-Type: application/json)
const CheckProjectEnvironmentUniqueNameResponseTemplate_B = `# API Response Information

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

// Response Template for the CheckProjectEnvironmentUniqueName tool (Status: 401, Content-Type: application/json)
const CheckProjectEnvironmentUniqueNameResponseTemplate_C = `# API Response Information

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

// Response Template for the CheckProjectEnvironmentUniqueName tool (Status: 500, Content-Type: application/json)
const CheckProjectEnvironmentUniqueNameResponseTemplate_D = `# API Response Information

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

// NewCheckProjectEnvironmentUniqueNameMCPTool creates the MCP Tool instance for CheckProjectEnvironmentUniqueName
func NewCheckProjectEnvironmentUniqueNameMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"CheckProjectEnvironmentUniqueName",
		"Check if environment unique name is available - Checks if an environment unique name is available for a project",
		[]byte(checkProjectEnvironmentUniqueNameInputSchema),
	)
}

// CheckProjectEnvironmentUniqueNameHandler is the handler function for the CheckProjectEnvironmentUniqueName tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func CheckProjectEnvironmentUniqueNameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "CheckProjectEnvironmentUniqueName")
}
