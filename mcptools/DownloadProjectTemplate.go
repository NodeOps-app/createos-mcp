package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the DownloadProjectTemplate tool
const downloadProjectTemplateInputSchema = `{
  "properties": {
    "purchase_id": {
      "format": "uuid",
      "type": "string"
    }
  },
  "required": [
    "purchase_id"
  ],
  "type": "object"
}`

// Response Template for the DownloadProjectTemplate tool (Status: 200, Content-Type: application/json)
const DownloadProjectTemplateResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Signed download URL

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **downloadUri**: Signed URL to download the resource (Type: string, uri):
  - **status** (Type: string):
      - Example: 'success'
`

// Response Template for the DownloadProjectTemplate tool (Status: 400, Content-Type: application/json)
const DownloadProjectTemplateResponseTemplate_B = `# API Response Information

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

// Response Template for the DownloadProjectTemplate tool (Status: 401, Content-Type: application/json)
const DownloadProjectTemplateResponseTemplate_C = `# API Response Information

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

// Response Template for the DownloadProjectTemplate tool (Status: 403, Content-Type: application/json)
const DownloadProjectTemplateResponseTemplate_D = `# API Response Information

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

// Response Template for the DownloadProjectTemplate tool (Status: 404, Content-Type: application/json)
const DownloadProjectTemplateResponseTemplate_E = `# API Response Information

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

// Response Template for the DownloadProjectTemplate tool (Status: 500, Content-Type: application/json)
const DownloadProjectTemplateResponseTemplate_F = `# API Response Information

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

// NewDownloadProjectTemplateMCPTool creates the MCP Tool instance for DownloadProjectTemplate
func NewDownloadProjectTemplateMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"DownloadProjectTemplate",
		"Get download URL for purchased template - Returns a signed URL to download the purchased template zip. Purchase must belong to the authenticated user.",
		[]byte(downloadProjectTemplateInputSchema),
	)
}

// DownloadProjectTemplateHandler is the handler function for the DownloadProjectTemplate tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func DownloadProjectTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "DownloadProjectTemplate")
}
