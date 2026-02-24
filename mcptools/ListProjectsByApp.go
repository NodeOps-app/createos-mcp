package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListProjectsByApp tool
const listProjectsByAppInputSchema = `{
  "properties": {
    "app_id": {
      "description": "App unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "limit": {
      "default": 10,
      "description": "Maximum number of projects to return (default: 10, max: 20)",
      "maximum": 20,
      "minimum": 1,
      "type": "integer"
    },
    "offset": {
      "default": 0,
      "description": "Number of projects to skip (default: 0)",
      "minimum": 0,
      "type": "integer"
    }
  },
  "required": [
    "app_id"
  ],
  "type": "object"
}`

// Response Template for the ListProjectsByApp tool (Status: 200, Content-Type: application/json)
const ListProjectsByAppResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Projects retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **data** (Type: array):
      - **Items** (Type: object):
        - **appId**: App ID the project is associated with (Type: string, uuid, nullable):
            - Nullable: true
        - **source** (Type: object):
        - **createdAt** (Type: string, date-time):
        - **description** (Type: string, nullable):
            - Nullable: true
        - **settings**: Project settings (schema depends on project type) (Type: object):
        - **status** (Type: string):
            - Enum: ['active', 'deleting', 'deleted']
        - **displayName** (Type: string):
        - **enabledSecurityScan**: Whether security scanning is enabled for the project (Type: boolean):
        - **userId** (Type: string):
        - **expireAt**: Expiration time for temporary projects (e.g., GPT-created projects) (Type: string, date-time, nullable):
            - Nullable: true
        - **uniqueName** (Type: string):
        - **deletedAt** (Type: string, date-time, nullable):
            - Nullable: true
        - **type** (Type: string):
            - Enum: ['vcs', 'image', 'upload']
        - **id** (Type: string, uuid):
        - **updatedAt** (Type: string, date-time):
    - **pagination** (Type: object):
      - **offset**: Number of items skipped (Type: integer):
          - Example: '0'
      - **total**: Total number of items available (Type: integer):
          - Example: '100'
      - **count**: Number of items in the current page (Type: integer):
          - Example: '10'
      - **limit**: Maximum number of items per page (Type: integer):
          - Example: '10'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the ListProjectsByApp tool (Status: 400, Content-Type: application/json)
const ListProjectsByAppResponseTemplate_B = `# API Response Information

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

// Response Template for the ListProjectsByApp tool (Status: 401, Content-Type: application/json)
const ListProjectsByAppResponseTemplate_C = `# API Response Information

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

// Response Template for the ListProjectsByApp tool (Status: 403, Content-Type: application/json)
const ListProjectsByAppResponseTemplate_D = `# API Response Information

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

// Response Template for the ListProjectsByApp tool (Status: 404, Content-Type: application/json)
const ListProjectsByAppResponseTemplate_E = `# API Response Information

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

// Response Template for the ListProjectsByApp tool (Status: 500, Content-Type: application/json)
const ListProjectsByAppResponseTemplate_F = `# API Response Information

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

// NewListProjectsByAppMCPTool creates the MCP Tool instance for ListProjectsByApp
func NewListProjectsByAppMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListProjectsByApp",
		"List projects by app - Returns all projects associated with an app",
		[]byte(listProjectsByAppInputSchema),
	)
}

// ListProjectsByAppHandler is the handler function for the ListProjectsByApp tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListProjectsByAppHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListProjectsByApp")
}
