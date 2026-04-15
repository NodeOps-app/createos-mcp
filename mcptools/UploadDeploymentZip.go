package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the UploadDeploymentZip tool
const uploadDeploymentZipInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "file": {
          "description": "ZIP file containing deployment files",
          "format": "binary",
          "type": "string"
        }
      },
      "type": "object"
    },
    "project_id": {
      "description": "Project unique identifier",
      "format": "uuid",
      "type": "string"
    }
  },
  "required": [
    "project_id",
    "body"
  ],
  "type": "object"
}`

// Response Template for the UploadDeploymentZip tool (Status: 200, Content-Type: application/json)
const UploadDeploymentZipResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Deployment created successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **id** (Type: string, uuid):
        - Example: '550e8400-e29b-41d4-a716-446655440000'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the UploadDeploymentZip tool (Status: 400, Content-Type: application/json)
const UploadDeploymentZipResponseTemplate_B = `# API Response Information

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

// Response Template for the UploadDeploymentZip tool (Status: 401, Content-Type: application/json)
const UploadDeploymentZipResponseTemplate_C = `# API Response Information

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

// Response Template for the UploadDeploymentZip tool (Status: 403, Content-Type: application/json)
const UploadDeploymentZipResponseTemplate_D = `# API Response Information

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

// Response Template for the UploadDeploymentZip tool (Status: 404, Content-Type: application/json)
const UploadDeploymentZipResponseTemplate_E = `# API Response Information

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

// Response Template for the UploadDeploymentZip tool (Status: 422, Content-Type: application/json)
const UploadDeploymentZipResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 422

**Content-Type:** application/json

> Unprocessable Entity - project is not an upload project or project not active

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the UploadDeploymentZip tool (Status: 429, Content-Type: application/json)
const UploadDeploymentZipResponseTemplate_G = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 429

**Content-Type:** application/json

> Too many requests - max manual deployments creation limit exceeded

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the UploadDeploymentZip tool (Status: 500, Content-Type: application/json)
const UploadDeploymentZipResponseTemplate_H = `# API Response Information

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

// NewUploadDeploymentZipMCPTool creates the MCP Tool instance for UploadDeploymentZip
func NewUploadDeploymentZipMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"UploadDeploymentZip",
		"Upload deployment as ZIP - Creates a new deployment by uploading a ZIP file. Only available for upload type projects.",
		[]byte(uploadDeploymentZipInputSchema),
	)
}

// UploadDeploymentZipHandler is the handler function for the UploadDeploymentZip tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func UploadDeploymentZipHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "UploadDeploymentZip")
}
