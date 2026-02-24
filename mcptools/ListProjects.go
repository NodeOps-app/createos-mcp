package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListProjects tool
const listProjectsInputSchema = `{
  "properties": {
    "app": {
      "description": "Filter by app ID. Use 'null' to filter projects not assigned to any app",
      "type": "string"
    },
    "limit": {
      "default": 10,
      "description": "Maximum number of projects to return (default: 10, max: 20)",
      "maximum": 20,
      "minimum": 1,
      "type": "integer"
    },
    "name": {
      "description": "Filter by project name (partial match, max 100 characters)",
      "maxLength": 100,
      "type": "string"
    },
    "offset": {
      "default": 0,
      "description": "Number of projects to skip (default: 0)",
      "minimum": 0,
      "type": "integer"
    },
    "status": {
      "description": "Filter by project status. Defaults to active and deleting (excludes deleted)",
      "enum": [
        "active",
        "deleting",
        "deleted"
      ],
      "type": "string"
    },
    "type": {
      "description": "Filter by project type",
      "enum": [
        "vcs",
        "image",
        "upload"
      ],
      "type": "string"
    }
  },
  "type": "object"
}`

// Response Template for the ListProjects tool (Status: 200, Content-Type: application/json)
const ListProjectsResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> List of projects retrieved successfully

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'true'
  - **data** (Type: object):
    - **data** (Type: array):
      - **Items** (Type: object):
        - **enabledSecurityScan**: Whether security scanning is enabled for the project (Type: boolean):
        - **userId** (Type: string):
        - **status** (Type: string):
            - Enum: ['active', 'deleting', 'deleted']
        - **displayName** (Type: string):
        - **deletedAt** (Type: string, date-time, nullable):
            - Nullable: true
        - **type** (Type: string):
            - Enum: ['vcs', 'image', 'upload']
        - **id** (Type: string, uuid):
        - **updatedAt** (Type: string, date-time):
        - **expireAt**: Expiration time for temporary projects (e.g., GPT-created projects) (Type: string, date-time, nullable):
            - Nullable: true
        - **uniqueName** (Type: string):
        - **createdAt** (Type: string, date-time):
        - **appId**: App ID the project is associated with (Type: string, uuid, nullable):
            - Nullable: true
        - **source** (Type: object):
        - **description** (Type: string, nullable):
            - Nullable: true
        - **settings**: Project settings (schema depends on project type) (Type: object):
    - **pagination** (Type: object):
      - **count**: Number of items in the current page (Type: integer):
          - Example: '10'
      - **limit**: Maximum number of items per page (Type: integer):
          - Example: '10'
      - **offset**: Number of items skipped (Type: integer):
          - Example: '0'
      - **total**: Total number of items available (Type: integer):
          - Example: '100'
`

// Response Template for the ListProjects tool (Status: 400, Content-Type: application/json)
const ListProjectsResponseTemplate_B = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 400

**Content-Type:** application/json

> Bad request - invalid status or type filter

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the ListProjects tool (Status: 401, Content-Type: application/json)
const ListProjectsResponseTemplate_C = `# API Response Information

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

// Response Template for the ListProjects tool (Status: 403, Content-Type: application/json)
const ListProjectsResponseTemplate_D = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 403

**Content-Type:** application/json

> Forbidden - user cannot list projects

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the ListProjects tool (Status: 500, Content-Type: application/json)
const ListProjectsResponseTemplate_E = `# API Response Information

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

// NewListProjectsMCPTool creates the MCP Tool instance for ListProjects
func NewListProjectsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListProjects",
		"List all projects - Returns a list of all projects for the authenticated user",
		[]byte(listProjectsInputSchema),
	)
}

// ListProjectsHandler is the handler function for the ListProjects tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListProjectsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListProjects")
}
