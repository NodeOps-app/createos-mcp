package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListProjectEnvironments tool
const listProjectEnvironmentsInputSchema = `{
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

// Response Template for the ListProjectEnvironments tool (Status: 200, Content-Type: application/json)
const ListProjectEnvironmentsResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Environments retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: array):
    - **Items** (Type: object):
      - **projectId** (Type: string, uuid):
      - **displayName**: Display name for the environment (Type: string):
      - **resources** (Type: object):
      - **updatedAt** (Type: string, date-time):
      - **extra**: Additional deployment information (URLs, etc.) (Type: object):
        - **Additional Properties**:
          - **property value** (Type: string):
      - **settings** (Type: object):
      - **isAutoPromoteEnabled** (Type: boolean):
      - **status** (Type: string):
          - Enum: ['active', 'updating', 'promoting', 'resizing', 'deleting']
      - **id** (Type: string, uuid):
      - **uniqueName** (Type: string):
      - **createdAt** (Type: string, date-time):
      - **description** (Type: string):
      - **projectDeploymentId** (Type: string, uuid, nullable):
          - Nullable: true
      - **branch** (Type: string):
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the ListProjectEnvironments tool (Status: 400, Content-Type: application/json)
const ListProjectEnvironmentsResponseTemplate_B = `# API Response Information

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

// Response Template for the ListProjectEnvironments tool (Status: 401, Content-Type: application/json)
const ListProjectEnvironmentsResponseTemplate_C = `# API Response Information

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

// Response Template for the ListProjectEnvironments tool (Status: 403, Content-Type: application/json)
const ListProjectEnvironmentsResponseTemplate_D = `# API Response Information

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

// Response Template for the ListProjectEnvironments tool (Status: 404, Content-Type: application/json)
const ListProjectEnvironmentsResponseTemplate_E = `# API Response Information

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

// Response Template for the ListProjectEnvironments tool (Status: 422, Content-Type: application/json)
const ListProjectEnvironmentsResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 422

**Content-Type:** application/json

> Unprocessable Entity - project not active

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the ListProjectEnvironments tool (Status: 500, Content-Type: application/json)
const ListProjectEnvironmentsResponseTemplate_G = `# API Response Information

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

// NewListProjectEnvironmentsMCPTool creates the MCP Tool instance for ListProjectEnvironments
func NewListProjectEnvironmentsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListProjectEnvironments",
		"List project environments - Returns all environments for a project",
		[]byte(listProjectEnvironmentsInputSchema),
	)
}

// ListProjectEnvironmentsHandler is the handler function for the ListProjectEnvironments tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListProjectEnvironmentsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListProjectEnvironments")
}
