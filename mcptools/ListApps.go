package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListApps tool
const listAppsInputSchema = `{
  "type": "object"
}`

// Response Template for the ListApps tool (Status: 200, Content-Type: application/json)
const ListAppsResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Apps retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: array):
    - **Items** (Type: object):
      - **updatedAt** (Type: string, date-time):
      - **userId**: User ID of the app owner (Type: string):
      - **color**: Hex color code for the app (Type: string):
          - Example: '#3B82F6'
      - **createdAt** (Type: string, date-time):
      - **description**: Description of the app (Type: string, nullable):
          - Nullable: true
          - Example: 'A collection of related projects and services'
      - **id**: Unique identifier of the app (Type: string, uuid):
          - Example: '550e8400-e29b-41d4-a716-446655440000'
      - **name**: Name of the app (Type: string):
          - Example: 'My Application'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the ListApps tool (Status: 401, Content-Type: application/json)
const ListAppsResponseTemplate_B = `# API Response Information

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

// Response Template for the ListApps tool (Status: 500, Content-Type: application/json)
const ListAppsResponseTemplate_C = `# API Response Information

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

// NewListAppsMCPTool creates the MCP Tool instance for ListApps
func NewListAppsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListApps",
		"List apps - Returns all apps for the authenticated user",
		[]byte(listAppsInputSchema),
	)
}

// ListAppsHandler is the handler function for the ListApps tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListAppsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListApps")
}
