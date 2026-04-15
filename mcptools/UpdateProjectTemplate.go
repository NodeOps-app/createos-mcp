package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the UpdateProjectTemplate tool
const updateProjectTemplateInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "amount": {
          "maximum": 5000,
          "minimum": 0,
          "type": "number"
        },
        "categories": {
          "items": {
            "maxLength": 128,
            "type": "string"
          },
          "maxItems": 5,
          "type": "array"
        },
        "description": {
          "maxLength": 10000,
          "minLength": 100,
          "type": "string"
        },
        "imageUri": {
          "maxLength": 100,
          "type": "string"
        },
        "name": {
          "maxLength": 100,
          "minLength": 10,
          "type": "string"
        },
        "overview": {
          "maxLength": 10000,
          "minLength": 100,
          "type": "string"
        },
        "useCases": {
          "maxLength": 10000,
          "minLength": 100,
          "type": "string"
        },
        "videoUrl": {
          "format": "uri",
          "maxLength": 100,
          "type": "string"
        }
      },
      "required": [
        "name",
        "description",
        "overview",
        "useCases",
        "videoUrl",
        "imageUri"
      ],
      "type": "object"
    },
    "template_id": {
      "format": "uuid",
      "type": "string"
    }
  },
  "required": [
    "template_id",
    "body"
  ],
  "type": "object"
}`

// Response Template for the UpdateProjectTemplate tool (Status: 200, Content-Type: application/json)
const UpdateProjectTemplateResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Template updated

## Response Structure

- Structure (Type: object):
  - **status** (Type: string):
      - Example: 'success'
  - **data** (Type: object):
    - **id**: Updated template ID (Type: string, uuid):
`

// Response Template for the UpdateProjectTemplate tool (Status: 400, Content-Type: application/json)
const UpdateProjectTemplateResponseTemplate_B = `# API Response Information

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

// Response Template for the UpdateProjectTemplate tool (Status: 401, Content-Type: application/json)
const UpdateProjectTemplateResponseTemplate_C = `# API Response Information

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

// Response Template for the UpdateProjectTemplate tool (Status: 403, Content-Type: application/json)
const UpdateProjectTemplateResponseTemplate_D = `# API Response Information

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

// Response Template for the UpdateProjectTemplate tool (Status: 404, Content-Type: application/json)
const UpdateProjectTemplateResponseTemplate_E = `# API Response Information

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

// Response Template for the UpdateProjectTemplate tool (Status: 500, Content-Type: application/json)
const UpdateProjectTemplateResponseTemplate_F = `# API Response Information

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

// NewUpdateProjectTemplateMCPTool creates the MCP Tool instance for UpdateProjectTemplate
func NewUpdateProjectTemplateMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"UpdateProjectTemplate",
		"Update project template - Updates the template metadata (name, description, overview, amount, use cases, video URL, image URI, categories). User must own the template.",
		[]byte(updateProjectTemplateInputSchema),
	)
}

// UpdateProjectTemplateHandler is the handler function for the UpdateProjectTemplate tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func UpdateProjectTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "UpdateProjectTemplate")
}
