package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListAPIKeys tool
const listAPIKeysInputSchema = `{
  "type": "object"
}`

// Response Template for the ListAPIKeys tool (Status: 200, Content-Type: application/json)
const ListAPIKeysResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> API keys retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: array):
    - **Items** (Type: object):
      - **updatedAt** (Type: string, date-time):
      - **userId** (Type: string):
      - **createdAt** (Type: string, date-time):
      - **description** (Type: string, nullable):
          - Nullable: true
      - **expiredAt** (Type: string, date-time, nullable):
          - Nullable: true
      - **id** (Type: string, uuid):
      - **name** (Type: string, nullable):
          - Nullable: true
      - **revokedAt** (Type: string, date-time, nullable):
          - Nullable: true
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the ListAPIKeys tool (Status: 401, Content-Type: application/json)
const ListAPIKeysResponseTemplate_B = `# API Response Information

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

// Response Template for the ListAPIKeys tool (Status: 404, Content-Type: application/json)
const ListAPIKeysResponseTemplate_C = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 404

**Content-Type:** application/json

> Not found - no API keys found

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the ListAPIKeys tool (Status: 500, Content-Type: application/json)
const ListAPIKeysResponseTemplate_D = `# API Response Information

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

// NewListAPIKeysMCPTool creates the MCP Tool instance for ListAPIKeys
func NewListAPIKeysMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListAPIKeys",
		"List API keys - Returns all API keys for the authenticated user",
		[]byte(listAPIKeysInputSchema),
	)
}

// ListAPIKeysHandler is the handler function for the ListAPIKeys tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListAPIKeysHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListAPIKeys")
}
