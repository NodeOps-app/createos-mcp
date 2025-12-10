package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the AssignDeploymentToProjectEnvironment tool
const assignDeploymentToProjectEnvironmentInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "deploymentId": {
          "example": "550e8400-e29b-41d4-a716-446655440000",
          "format": "uuid",
          "type": "string"
        }
      },
      "required": [
        "deploymentId"
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

// Response Template for the AssignDeploymentToProjectEnvironment tool (Status: 200, Content-Type: application/json)
const AssignDeploymentToProjectEnvironmentResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Deployment assigned successfully

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'true'
  - **data** (Type: string):
      - Example: 'deployment assigned to project environment successfully'
`

// Response Template for the AssignDeploymentToProjectEnvironment tool (Status: 400, Content-Type: application/json)
const AssignDeploymentToProjectEnvironmentResponseTemplate_B = `# API Response Information

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

// Response Template for the AssignDeploymentToProjectEnvironment tool (Status: 401, Content-Type: application/json)
const AssignDeploymentToProjectEnvironmentResponseTemplate_C = `# API Response Information

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

// Response Template for the AssignDeploymentToProjectEnvironment tool (Status: 403, Content-Type: application/json)
const AssignDeploymentToProjectEnvironmentResponseTemplate_D = `# API Response Information

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

// Response Template for the AssignDeploymentToProjectEnvironment tool (Status: 404, Content-Type: application/json)
const AssignDeploymentToProjectEnvironmentResponseTemplate_E = `# API Response Information

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

// Response Template for the AssignDeploymentToProjectEnvironment tool (Status: 422, Content-Type: application/json)
const AssignDeploymentToProjectEnvironmentResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 422

**Content-Type:** application/json

> Unprocessable Entity - deployment not eligible or project not active

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the AssignDeploymentToProjectEnvironment tool (Status: 500, Content-Type: application/json)
const AssignDeploymentToProjectEnvironmentResponseTemplate_G = `# API Response Information

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

// NewAssignDeploymentToProjectEnvironmentMCPTool creates the MCP Tool instance for AssignDeploymentToProjectEnvironment
func NewAssignDeploymentToProjectEnvironmentMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"AssignDeploymentToProjectEnvironment",
		"Assign deployment to environment - Assigns a deployment to a project environment",
		[]byte(assignDeploymentToProjectEnvironmentInputSchema),
	)
}

// AssignDeploymentToProjectEnvironmentHandler is the handler function for the AssignDeploymentToProjectEnvironment tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func AssignDeploymentToProjectEnvironmentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "AssignDeploymentToProjectEnvironment")
}
