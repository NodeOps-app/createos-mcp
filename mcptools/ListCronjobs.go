package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListCronjobs tool
const listCronjobsInputSchema = `{
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

// Response Template for the ListCronjobs tool (Status: 200, Content-Type: application/json)
const ListCronjobsResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Cronjobs retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: array):
    - **Items** (Type: object):
      - **projectId** (Type: string, uuid):
      - **schedule** (Type: string):
      - **type** (Type: string):
          - Enum: ['http', 'artifact']
      - **suspendText** (Type: string, nullable):
          - Nullable: true
      - **createdAt** (Type: string, date-time):
      - **name** (Type: string):
      - **id** (Type: string, uuid):
      - **status** (Type: string):
          - Enum: ['pending', 'building', 'active', 'suspending', 'suspended', 'deleting', 'deleted']
      - **suspendedAt** (Type: string, date-time, nullable):
          - Nullable: true
      - **metadata** (Type: object, nullable):
          - Nullable: true
      - **userRequestTerminatedAt** (Type: string, date-time, nullable):
          - Nullable: true
      - **updatedAt** (Type: string, date-time):
      - **suspendedByOwner** (Type: boolean, nullable):
          - Nullable: true
      - **environmentId** (Type: string, uuid):
      - **settings** (Type: object, nullable):
          - Nullable: true
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the ListCronjobs tool (Status: 401, Content-Type: application/json)
const ListCronjobsResponseTemplate_B = `# API Response Information

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

// Response Template for the ListCronjobs tool (Status: 403, Content-Type: application/json)
const ListCronjobsResponseTemplate_C = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 403

**Content-Type:** application/json

> Forbidden - user does not have permission

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the ListCronjobs tool (Status: 404, Content-Type: application/json)
const ListCronjobsResponseTemplate_D = `# API Response Information

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

// Response Template for the ListCronjobs tool (Status: 500, Content-Type: application/json)
const ListCronjobsResponseTemplate_E = `# API Response Information

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

// NewListCronjobsMCPTool creates the MCP Tool instance for ListCronjobs
func NewListCronjobsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListCronjobs",
		"List cronjobs - Lists all cronjobs for a project.",
		[]byte(listCronjobsInputSchema),
	)
}

// ListCronjobsHandler is the handler function for the ListCronjobs tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListCronjobsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListCronjobs")
}
