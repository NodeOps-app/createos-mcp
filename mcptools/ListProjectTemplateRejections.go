package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListProjectTemplateRejections tool
const listProjectTemplateRejectionsInputSchema = `{
  "properties": {
    "template_id": {
      "format": "uuid",
      "type": "string"
    }
  },
  "required": [
    "template_id"
  ],
  "type": "object"
}`

// Response Template for the ListProjectTemplateRejections tool (Status: 200, Content-Type: application/json)
const ListProjectTemplateRejectionsResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> List of rejections

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **rejections** (Type: array):
      - **Items** (Type: object):
        - **id** (Type: integer):
        - **projectTemplateId** (Type: string, uuid):
        - **rejectionReason** (Type: string):
        - **updatedAt** (Type: string, date-time):
        - **createdAt** (Type: string, date-time):
  - **status** (Type: string):
      - Example: 'success'
`

// Response Template for the ListProjectTemplateRejections tool (Status: 400, Content-Type: application/json)
const ListProjectTemplateRejectionsResponseTemplate_B = `# API Response Information

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

// Response Template for the ListProjectTemplateRejections tool (Status: 401, Content-Type: application/json)
const ListProjectTemplateRejectionsResponseTemplate_C = `# API Response Information

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

// Response Template for the ListProjectTemplateRejections tool (Status: 404, Content-Type: application/json)
const ListProjectTemplateRejectionsResponseTemplate_D = `# API Response Information

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

// Response Template for the ListProjectTemplateRejections tool (Status: 500, Content-Type: application/json)
const ListProjectTemplateRejectionsResponseTemplate_E = `# API Response Information

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

// NewListProjectTemplateRejectionsMCPTool creates the MCP Tool instance for ListProjectTemplateRejections
func NewListProjectTemplateRejectionsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListProjectTemplateRejections",
		"List project template rejections - Returns rejection history for the template. Only the template owner can list rejections.",
		[]byte(listProjectTemplateRejectionsInputSchema),
	)
}

// ListProjectTemplateRejectionsHandler is the handler function for the ListProjectTemplateRejections tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListProjectTemplateRejectionsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListProjectTemplateRejections")
}
