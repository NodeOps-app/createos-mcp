package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the RetriggerDeployment tool
const retriggerDeploymentInputSchema = `{
  "properties": {
    "deployment_id": {
      "description": "Deployment unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "project_id": {
      "description": "Project unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "settings": {
      "default": "deployment",
      "description": "Which settings to use for the new deployment. \"project\": Use current project settings. \"deployment\": Use settings from the original deployment (default).",
      "enum": [
        "project",
        "deployment"
      ],
      "type": "string"
    }
  },
  "required": [
    "project_id",
    "deployment_id"
  ],
  "type": "object"
}`

// Response Template for the RetriggerDeployment tool (Status: 200, Content-Type: application/json)
const RetriggerDeploymentResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Deployment triggered successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: string):
      - Example: 'deployment triggered'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the RetriggerDeployment tool (Status: 400, Content-Type: application/json)
const RetriggerDeploymentResponseTemplate_B = `# API Response Information

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

// Response Template for the RetriggerDeployment tool (Status: 401, Content-Type: application/json)
const RetriggerDeploymentResponseTemplate_C = `# API Response Information

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

// Response Template for the RetriggerDeployment tool (Status: 403, Content-Type: application/json)
const RetriggerDeploymentResponseTemplate_D = `# API Response Information

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

// Response Template for the RetriggerDeployment tool (Status: 404, Content-Type: application/json)
const RetriggerDeploymentResponseTemplate_E = `# API Response Information

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

// Response Template for the RetriggerDeployment tool (Status: 422, Content-Type: application/json)
const RetriggerDeploymentResponseTemplate_F = `# API Response Information

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

// Response Template for the RetriggerDeployment tool (Status: 500, Content-Type: application/json)
const RetriggerDeploymentResponseTemplate_G = `# API Response Information

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

// NewRetriggerDeploymentMCPTool creates the MCP Tool instance for RetriggerDeployment
func NewRetriggerDeploymentMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"RetriggerDeployment",
		"Retrigger deployment - Creates a new deployment with the same source as an existing deployment. The settings can be taken from either the project (current project settings) or the original deployment (deployment settings) only works when VCS.",
		[]byte(retriggerDeploymentInputSchema),
	)
}

// RetriggerDeploymentHandler is the handler function for the RetriggerDeployment tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func RetriggerDeploymentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "RetriggerDeployment")
}
