package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the AddProjectsToApp tool
const addProjectsToAppInputSchema = `{
  "properties": {
    "app_id": {
      "description": "App unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "body": {
      "properties": {
        "projectIds": {
          "description": "List of project IDs to add to or remove from the app",
          "example": [
            "550e8400-e29b-41d4-a716-446655440000",
            "550e8400-e29b-41d4-a716-446655440001"
          ],
          "items": {
            "format": "uuid",
            "type": "string"
          },
          "type": "array"
        }
      },
      "required": [
        "projectIds"
      ],
      "type": "object"
    }
  },
  "required": [
    "app_id",
    "body"
  ],
  "type": "object"
}`

// Response Template for the AddProjectsToApp tool (Status: 200, Content-Type: application/json)
const AddProjectsToAppResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Projects added to app successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: string):
      - Example: 'projects added to app successfully'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the AddProjectsToApp tool (Status: 400, Content-Type: application/json)
const AddProjectsToAppResponseTemplate_B = `# API Response Information

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

// Response Template for the AddProjectsToApp tool (Status: 401, Content-Type: application/json)
const AddProjectsToAppResponseTemplate_C = `# API Response Information

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

// Response Template for the AddProjectsToApp tool (Status: 403, Content-Type: application/json)
const AddProjectsToAppResponseTemplate_D = `# API Response Information

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

// Response Template for the AddProjectsToApp tool (Status: 404, Content-Type: application/json)
const AddProjectsToAppResponseTemplate_E = `# API Response Information

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

// Response Template for the AddProjectsToApp tool (Status: 500, Content-Type: application/json)
const AddProjectsToAppResponseTemplate_F = `# API Response Information

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

// NewAddProjectsToAppMCPTool creates the MCP Tool instance for AddProjectsToApp
func NewAddProjectsToAppMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"AddProjectsToApp",
		"Add projects to app - Associates one or more projects with an app",
		[]byte(addProjectsToAppInputSchema),
	)
}

// AddProjectsToAppHandler is the handler function for the AddProjectsToApp tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func AddProjectsToAppHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "AddProjectsToApp")
}
