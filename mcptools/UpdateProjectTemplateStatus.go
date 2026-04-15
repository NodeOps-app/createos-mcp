package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the UpdateProjectTemplateStatus tool
const updateProjectTemplateStatusInputSchema = `{
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

// Response Template for the UpdateProjectTemplateStatus tool (Status: 200, Content-Type: application/json)
const UpdateProjectTemplateStatusResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Status updated to underReview

## Response Structure

- Structure (Type: object):
  - **status** (Type: string):
      - Example: 'success'
  - **data** (Type: object):
    - **status** (Type: string):
        - Example: 'underReview'
`

// Response Template for the UpdateProjectTemplateStatus tool (Status: 400, Content-Type: application/json)
const UpdateProjectTemplateStatusResponseTemplate_B = `# API Response Information

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

// Response Template for the UpdateProjectTemplateStatus tool (Status: 401, Content-Type: application/json)
const UpdateProjectTemplateStatusResponseTemplate_C = `# API Response Information

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

// Response Template for the UpdateProjectTemplateStatus tool (Status: 403, Content-Type: application/json)
const UpdateProjectTemplateStatusResponseTemplate_D = `# API Response Information

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

// Response Template for the UpdateProjectTemplateStatus tool (Status: 404, Content-Type: application/json)
const UpdateProjectTemplateStatusResponseTemplate_E = `# API Response Information

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

// Response Template for the UpdateProjectTemplateStatus tool (Status: 422, Content-Type: application/json)
const UpdateProjectTemplateStatusResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 422

**Content-Type:** application/json

> Can only submit from draft or rejected

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the UpdateProjectTemplateStatus tool (Status: 500, Content-Type: application/json)
const UpdateProjectTemplateStatusResponseTemplate_G = `# API Response Information

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

// NewUpdateProjectTemplateStatusMCPTool creates the MCP Tool instance for UpdateProjectTemplateStatus
func NewUpdateProjectTemplateStatusMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"UpdateProjectTemplateStatus",
		"Send project template for review - Transitions template status from draft or rejected to underReview. No request body. User must own the template.",
		[]byte(updateProjectTemplateStatusInputSchema),
	)
}

// UpdateProjectTemplateStatusHandler is the handler function for the UpdateProjectTemplateStatus tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func UpdateProjectTemplateStatusHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "UpdateProjectTemplateStatus")
}
