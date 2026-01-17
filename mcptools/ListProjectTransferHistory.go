package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListProjectTransferHistory tool
const listProjectTransferHistoryInputSchema = `{
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

// Response Template for the ListProjectTransferHistory tool (Status: 200, Content-Type: application/json)
const ListProjectTransferHistoryResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Transfer history retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: array):
    - **Items** (Type: object):
      - **projectId**: Project ID that was transferred (Type: string, uuid):
          - Example: '550e8400-e29b-41d4-a716-446655440000'
      - **toUserId**: User ID of the new owner (Type: string):
          - Example: 'user_xyz789'
      - **updatedAt**: Last update time of the record (Type: string, date-time):
      - **createdAt**: When the transfer occurred (Type: string, date-time):
      - **fromUserId**: User ID of the previous owner (Type: string):
          - Example: 'user_abc123'
      - **id**: Unique identifier of the transfer record (Type: string, uuid):
          - Example: '550e8400-e29b-41d4-a716-446655440001'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the ListProjectTransferHistory tool (Status: 400, Content-Type: application/json)
const ListProjectTransferHistoryResponseTemplate_B = `# API Response Information

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

// Response Template for the ListProjectTransferHistory tool (Status: 401, Content-Type: application/json)
const ListProjectTransferHistoryResponseTemplate_C = `# API Response Information

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

// Response Template for the ListProjectTransferHistory tool (Status: 403, Content-Type: application/json)
const ListProjectTransferHistoryResponseTemplate_D = `# API Response Information

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

// Response Template for the ListProjectTransferHistory tool (Status: 404, Content-Type: application/json)
const ListProjectTransferHistoryResponseTemplate_E = `# API Response Information

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

// Response Template for the ListProjectTransferHistory tool (Status: 500, Content-Type: application/json)
const ListProjectTransferHistoryResponseTemplate_F = `# API Response Information

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

// NewListProjectTransferHistoryMCPTool creates the MCP Tool instance for ListProjectTransferHistory
func NewListProjectTransferHistoryMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListProjectTransferHistory",
		"List project transfer history - Returns the transfer history of a project, showing all ownership transfers that have occurred.",
		[]byte(listProjectTransferHistoryInputSchema),
	)
}

// ListProjectTransferHistoryHandler is the handler function for the ListProjectTransferHistory tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListProjectTransferHistoryHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListProjectTransferHistory")
}
