package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the DeployProjectTemplateViaGithub tool
const deployProjectTemplateViaGithubInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "installationId": {
          "description": "GitHub App installation ID",
          "format": "int64",
          "type": "integer"
        }
      },
      "required": [
        "installationId"
      ],
      "type": "object"
    },
    "purchase_id": {
      "format": "uuid",
      "type": "string"
    }
  },
  "required": [
    "purchase_id",
    "body"
  ],
  "type": "object"
}`

// Response Template for the DeployProjectTemplateViaGithub tool (Status: 200, Content-Type: application/json)
const DeployProjectTemplateViaGithubResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Job created; use jobId to poll status

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **jobId**: ID to poll for job status (Type: string, uuid):
  - **status** (Type: string):
      - Example: 'success'
`

// Response Template for the DeployProjectTemplateViaGithub tool (Status: 400, Content-Type: application/json)
const DeployProjectTemplateViaGithubResponseTemplate_B = `# API Response Information

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

// Response Template for the DeployProjectTemplateViaGithub tool (Status: 401, Content-Type: application/json)
const DeployProjectTemplateViaGithubResponseTemplate_C = `# API Response Information

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

// Response Template for the DeployProjectTemplateViaGithub tool (Status: 403, Content-Type: application/json)
const DeployProjectTemplateViaGithubResponseTemplate_D = `# API Response Information

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

// Response Template for the DeployProjectTemplateViaGithub tool (Status: 404, Content-Type: application/json)
const DeployProjectTemplateViaGithubResponseTemplate_E = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 404

**Content-Type:** application/json

> Purchase or template not found

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the DeployProjectTemplateViaGithub tool (Status: 422, Content-Type: application/json)
const DeployProjectTemplateViaGithubResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 422

**Content-Type:** application/json

> Connect a GitHub account or installation not found

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the DeployProjectTemplateViaGithub tool (Status: 429, Content-Type: application/json)
const DeployProjectTemplateViaGithubResponseTemplate_G = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 429

**Content-Type:** application/json

> Max projects per user reached

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the DeployProjectTemplateViaGithub tool (Status: 500, Content-Type: application/json)
const DeployProjectTemplateViaGithubResponseTemplate_H = `# API Response Information

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

// NewDeployProjectTemplateViaGithubMCPTool creates the MCP Tool instance for DeployProjectTemplateViaGithub
func NewDeployProjectTemplateViaGithubMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"DeployProjectTemplateViaGithub",
		"Deploy purchased template via GitHub - Starts an async job to deploy the purchased template to a GitHub repository using the given installation. User must have the GitHub account connected and have capacity to create a new project.",
		[]byte(deployProjectTemplateViaGithubInputSchema),
	)
}

// DeployProjectTemplateViaGithubHandler is the handler function for the DeployProjectTemplateViaGithub tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func DeployProjectTemplateViaGithubHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "DeployProjectTemplateViaGithub")
}
