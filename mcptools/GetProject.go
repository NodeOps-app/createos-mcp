package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the GetProject tool
const getProjectInputSchema = `{
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

// Response Template for the GetProject tool (Status: 200, Content-Type: application/json)
const GetProjectResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Project retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **description** (Type: string, nullable):
        - Nullable: true
    - **settings**: Project settings (schema depends on project type) (Type: object):
    - **status** (Type: string):
        - Enum: ['active', 'deleting', 'deleted']
    - **displayName** (Type: string):
    - **enabledSecurityScan**: Whether security scanning is enabled for the project (Type: boolean):
    - **userId** (Type: string):
    - **updatedAt** (Type: string, date-time):
    - **expireAt**: Expiration time for temporary projects (e.g., GPT-created projects) (Type: string, date-time, nullable):
        - Nullable: true
    - **uniqueName** (Type: string):
    - **deletedAt** (Type: string, date-time, nullable):
        - Nullable: true
    - **type** (Type: string):
        - Enum: ['vcs', 'image', 'upload']
    - **id** (Type: string, uuid):
    - **appId**: App ID the project is associated with (Type: string, uuid, nullable):
        - Nullable: true
    - **source** (Type: object):
    - **createdAt** (Type: string, date-time):
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the GetProject tool (Status: 400, Content-Type: application/json)
const GetProjectResponseTemplate_B = `# API Response Information

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

// Response Template for the GetProject tool (Status: 401, Content-Type: application/json)
const GetProjectResponseTemplate_C = `# API Response Information

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

// Response Template for the GetProject tool (Status: 403, Content-Type: application/json)
const GetProjectResponseTemplate_D = `# API Response Information

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

// Response Template for the GetProject tool (Status: 404, Content-Type: application/json)
const GetProjectResponseTemplate_E = `# API Response Information

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

// Response Template for the GetProject tool (Status: 500, Content-Type: application/json)
const GetProjectResponseTemplate_F = `# API Response Information

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

// NewGetProjectMCPTool creates the MCP Tool instance for GetProject
func NewGetProjectMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetProject",
		"Get project by ID - Retrieves detailed information about a specific project",
		[]byte(getProjectInputSchema),
	)
}

// GetProjectHandler is the handler function for the GetProject tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func GetProjectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "GetProject")
}
