package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the UpdateProjectEnvironment tool
const updateProjectEnvironmentInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "branch": {
          "description": "For VCS projects: The branch name to associate with the environment. Must exist in the repository. For image projects: This field should be omitted or null.",
          "example": "main",
          "maxLength": 255,
          "minLength": 1,
          "type": [
            "string",
            "null"
          ]
        },
        "description": {
          "example": "Updated production environment",
          "maxLength": 2048,
          "minLength": 4,
          "pattern": "^[a-zA-Z0-9 _-]+$",
          "type": "string"
        },
        "displayName": {
          "description": "Display name for the environment",
          "example": "Production",
          "maxLength": 48,
          "minLength": 4,
          "pattern": "^[a-zA-Z0-9 _-]+$",
          "type": "string"
        },
        "isAutoPromoteEnabled": {
          "example": false,
          "type": "boolean"
        }
      },
      "required": [
        "displayName",
        "description"
      ],
      "type": "object"
    },
    "environment_id": {
      "description": "Environment unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "project_id": {
      "description": "Project unique identifier",
      "format": "uuid",
      "type": "string"
    }
  },
  "required": [
    "project_id",
    "environment_id",
    "body"
  ],
  "type": "object"
}`

// Response Template for the UpdateProjectEnvironment tool (Status: 200, Content-Type: application/json)
const UpdateProjectEnvironmentResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Environment updated successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: string):
      - Example: 'project environment updated successfully'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the UpdateProjectEnvironment tool (Status: 400, Content-Type: application/json)
const UpdateProjectEnvironmentResponseTemplate_B = `# API Response Information

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

// Response Template for the UpdateProjectEnvironment tool (Status: 401, Content-Type: application/json)
const UpdateProjectEnvironmentResponseTemplate_C = `# API Response Information

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

// Response Template for the UpdateProjectEnvironment tool (Status: 403, Content-Type: application/json)
const UpdateProjectEnvironmentResponseTemplate_D = `# API Response Information

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

// Response Template for the UpdateProjectEnvironment tool (Status: 404, Content-Type: application/json)
const UpdateProjectEnvironmentResponseTemplate_E = `# API Response Information

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

// Response Template for the UpdateProjectEnvironment tool (Status: 422, Content-Type: application/json)
const UpdateProjectEnvironmentResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 422

**Content-Type:** application/json

> Unprocessable Entity - environment being deleted or project not active

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the UpdateProjectEnvironment tool (Status: 500, Content-Type: application/json)
const UpdateProjectEnvironmentResponseTemplate_G = `# API Response Information

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

// NewUpdateProjectEnvironmentMCPTool creates the MCP Tool instance for UpdateProjectEnvironment
func NewUpdateProjectEnvironmentMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"UpdateProjectEnvironment",
		"Update project environment - Updates environment description, branch (for VCS projects), and auto-promote settings. For VCS projects: branch is required and must exist in the repository. For image projects: branch should be omitted or null.",
		[]byte(updateProjectEnvironmentInputSchema),
	)
}

// UpdateProjectEnvironmentHandler is the handler function for the UpdateProjectEnvironment tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func UpdateProjectEnvironmentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "UpdateProjectEnvironment")
}
