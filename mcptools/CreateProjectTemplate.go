package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the CreateProjectTemplate tool
const createProjectTemplateInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "amount": {
          "default": 0,
          "maximum": 5000,
          "minimum": 0,
          "type": "number"
        },
        "buildResources": {
          "description": "CPU/memory/timeout for build",
          "type": "object"
        },
        "categories": {
          "items": {
            "maxLength": 128,
            "type": "string"
          },
          "maxItems": 5,
          "type": "array"
        },
        "deploymentResources": {
          "description": "CPU/memory/replicas for deployment",
          "type": "object"
        },
        "description": {
          "maxLength": 10000,
          "minLength": 100,
          "type": "string"
        },
        "environmentConfiguration": {
          "description": "Environment settings and resources (settings, resources)",
          "type": "object"
        },
        "environmentId": {
          "format": "uuid",
          "type": "string"
        },
        "imageUri": {
          "maxLength": 100,
          "type": "string"
        },
        "name": {
          "description": "Display name for the template",
          "maxLength": 100,
          "minLength": 10,
          "type": "string"
        },
        "overview": {
          "maxLength": 10000,
          "minLength": 100,
          "type": "string"
        },
        "projectConfiguration": {
          "description": "VCS project settings (e.g. runtime, framework, buildCommand)",
          "type": "object"
        },
        "projectId": {
          "format": "uuid",
          "type": "string"
        },
        "useCases": {
          "maxLength": 10000,
          "minLength": 100,
          "type": "string"
        },
        "videoUrl": {
          "type": "string"
        }
      },
      "required": [
        "name",
        "description",
        "overview",
        "projectId",
        "environmentId",
        "useCases",
        "imageUri",
        "projectConfiguration",
        "environmentConfiguration",
        "deploymentResources",
        "buildResources"
      ],
      "type": "object"
    }
  },
  "required": [
    "body"
  ],
  "type": "object"
}`

// Response Template for the CreateProjectTemplate tool (Status: 201, Content-Type: application/json)
const CreateProjectTemplateResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 201

**Content-Type:** application/json

> Project template created

## Response Structure

- Structure (Type: object):
  - **status** (Type: string):
      - Example: 'success'
  - **data** (Type: object):
    - **id**: Created project template ID (Type: string, uuid):
`

// Response Template for the CreateProjectTemplate tool (Status: 400, Content-Type: application/json)
const CreateProjectTemplateResponseTemplate_B = `# API Response Information

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

// Response Template for the CreateProjectTemplate tool (Status: 401, Content-Type: application/json)
const CreateProjectTemplateResponseTemplate_C = `# API Response Information

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

// Response Template for the CreateProjectTemplate tool (Status: 403, Content-Type: application/json)
const CreateProjectTemplateResponseTemplate_D = `# API Response Information

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

// Response Template for the CreateProjectTemplate tool (Status: 404, Content-Type: application/json)
const CreateProjectTemplateResponseTemplate_E = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 404

**Content-Type:** application/json

> Project or environment not found

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the CreateProjectTemplate tool (Status: 422, Content-Type: application/json)
const CreateProjectTemplateResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 422

**Content-Type:** application/json

> Project not active or environment does not belong to project

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the CreateProjectTemplate tool (Status: 429, Content-Type: application/json)
const CreateProjectTemplateResponseTemplate_G = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 429

**Content-Type:** application/json

> Too many draft/under-review templates (max 10)

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the CreateProjectTemplate tool (Status: 500, Content-Type: application/json)
const CreateProjectTemplateResponseTemplate_H = `# API Response Information

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

// NewCreateProjectTemplateMCPTool creates the MCP Tool instance for CreateProjectTemplate
func NewCreateProjectTemplateMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"CreateProjectTemplate",
		"Create project template - Creates a new project template in draft status from an existing project and environment. User can have at most 10 templates in draft or under review.",
		[]byte(createProjectTemplateInputSchema),
	)
}

// CreateProjectTemplateHandler is the handler function for the CreateProjectTemplate tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func CreateProjectTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "CreateProjectTemplate")
}
