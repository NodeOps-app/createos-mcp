package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the UploadDeploymentBase64Files tool
const uploadDeploymentBase64FilesInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "files": {
          "description": "Array of files to upload (maximum 100 files)",
          "items": {
            "properties": {
              "content": {
                "description": "File content (plain text for /files endpoint, base64-encoded for /files/base64 endpoint)",
                "example": "\u003c!DOCTYPE html\u003e\u003chtml\u003e\u003cbody\u003eHello\u003c/body\u003e\u003c/html\u003e",
                "type": "string"
              },
              "path": {
                "description": "File path relative to project root",
                "example": "index.html",
                "type": "string"
              }
            },
            "required": [
              "path",
              "content"
            ],
            "type": "object"
          },
          "maxItems": 100,
          "type": "array"
        }
      },
      "required": [
        "files"
      ],
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

// Response Template for the UploadDeploymentBase64Files tool (Status: 200, Content-Type: application/json)
const UploadDeploymentBase64FilesResponseTemplate_A = `# API Response Information

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

// Response Template for the UploadDeploymentBase64Files tool (Status: 400, Content-Type: application/json)
const UploadDeploymentBase64FilesResponseTemplate_B = `# API Response Information

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

// Response Template for the UploadDeploymentBase64Files tool (Status: 401, Content-Type: application/json)
const UploadDeploymentBase64FilesResponseTemplate_C = `# API Response Information

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

// Response Template for the UploadDeploymentBase64Files tool (Status: 403, Content-Type: application/json)
const UploadDeploymentBase64FilesResponseTemplate_D = `# API Response Information

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

// Response Template for the UploadDeploymentBase64Files tool (Status: 404, Content-Type: application/json)
const UploadDeploymentBase64FilesResponseTemplate_E = `# API Response Information

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

// Response Template for the UploadDeploymentBase64Files tool (Status: 422, Content-Type: application/json)
const UploadDeploymentBase64FilesResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 422

**Content-Type:** application/json

> Unprocessable Entity - project is not an upload project or project not active

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the UploadDeploymentBase64Files tool (Status: 429, Content-Type: application/json)
const UploadDeploymentBase64FilesResponseTemplate_G = `# API Response Information

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

// Response Template for the UploadDeploymentBase64Files tool (Status: 500, Content-Type: application/json)
const UploadDeploymentBase64FilesResponseTemplate_H = `# API Response Information

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

// NewUploadDeploymentBase64FilesMCPTool creates the MCP Tool instance for UploadDeploymentBase64Files
func NewUploadDeploymentBase64FilesMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"UploadDeploymentBase64Files",
		"Upload deployment files (Base64) - Creates a new deployment by uploading base64-encoded files. Only available for upload type projects. File contents must be base64-encoded.",
		[]byte(uploadDeploymentBase64FilesInputSchema),
	)
}

// UploadDeploymentBase64FilesHandler is the handler function for the UploadDeploymentBase64Files tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func UploadDeploymentBase64FilesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "UploadDeploymentBase64Files")
}
