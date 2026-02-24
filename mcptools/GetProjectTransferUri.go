package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the GetProjectTransferUri tool
const getProjectTransferUriInputSchema = `{
  "properties": {
    "project_id": {
      "description": "Project unique identifier",
      "format": "uuid",
      "type": "string"
    }
  },
  "required": [
    "project_id"
  ],
  "type": "object"
}`

// Response Template for the GetProjectTransferUri tool (Status: 200, Content-Type: application/json)
const GetProjectTransferUriResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Transfer URI generated successfully

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'true'
  - **data** (Type: object):
    - **token**: JWT token for project transfer (valid for 6 hours) (Type: string):
        - Example: 'eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...'
    - **expiresAt**: Token expiration time (Type: string, date-time):
        - Example: '2025-01-18T12:00:00Z'
    - **frontendUri**: Complete frontend URL with token for accepting the transfer (Type: string):
        - Example: 'https://app.example.com/projects/550e8400-e29b-41d4-a716-446655440000/transfer?token=eyJ...&source=ChatGPT'
`

// Response Template for the GetProjectTransferUri tool (Status: 400, Content-Type: application/json)
const GetProjectTransferUriResponseTemplate_B = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 400

**Content-Type:** application/json

> Bad request - project not in active state

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the GetProjectTransferUri tool (Status: 401, Content-Type: application/json)
const GetProjectTransferUriResponseTemplate_C = `# API Response Information

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

// Response Template for the GetProjectTransferUri tool (Status: 403, Content-Type: application/json)
const GetProjectTransferUriResponseTemplate_D = `# API Response Information

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

// Response Template for the GetProjectTransferUri tool (Status: 404, Content-Type: application/json)
const GetProjectTransferUriResponseTemplate_E = `# API Response Information

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

// Response Template for the GetProjectTransferUri tool (Status: 500, Content-Type: application/json)
const GetProjectTransferUriResponseTemplate_F = `# API Response Information

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

// NewGetProjectTransferUriMCPTool creates the MCP Tool instance for GetProjectTransferUri
func NewGetProjectTransferUriMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetProjectTransferUri",
		"Get project transfer URI - Generates a transfer URI and token that can be used to transfer project ownership to another user. The token is valid for 6 hours.",
		[]byte(getProjectTransferUriInputSchema),
	)
}

// GetProjectTransferUriHandler is the handler function for the GetProjectTransferUri tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func GetProjectTransferUriHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "GetProjectTransferUri")
}
