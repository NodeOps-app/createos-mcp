package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the CreateApp tool
const createAppInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "color": {
          "description": "Hex color code for the app (defaults to",
          "example": "#3B82F6",
          "pattern": "^#[0-9A-Fa-f]{6}$",
          "type": "string"
        },
        "description": {
          "description": "Description of the app",
          "example": "A collection of related projects and services",
          "maxLength": 2048,
          "minLength": 1,
          "type": [
            "string",
            "null"
          ]
        },
        "name": {
          "description": "Name of the app",
          "example": "My Application",
          "maxLength": 255,
          "minLength": 2,
          "type": "string"
        }
      },
      "required": [
        "name"
      ],
      "type": "object"
    }
  },
  "required": [
    "body"
  ],
  "type": "object"
}`

// Response Template for the CreateApp tool (Status: 200, Content-Type: application/json)
const CreateAppResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> App created successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **id** (Type: string, uuid):
        - Example: '550e8400-e29b-41d4-a716-446655440000'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the CreateApp tool (Status: 400, Content-Type: application/json)
const CreateAppResponseTemplate_B = `# API Response Information

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

// Response Template for the CreateApp tool (Status: 401, Content-Type: application/json)
const CreateAppResponseTemplate_C = `# API Response Information

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

// Response Template for the CreateApp tool (Status: 429, Content-Type: application/json)
const CreateAppResponseTemplate_D = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 429

**Content-Type:** application/json

> Too many requests - maximum apps per user reached

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the CreateApp tool (Status: 500, Content-Type: application/json)
const CreateAppResponseTemplate_E = `# API Response Information

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

// NewCreateAppMCPTool creates the MCP Tool instance for CreateApp
func NewCreateAppMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"CreateApp",
		"Create app - Creates a new app to organize projects and services",
		[]byte(createAppInputSchema),
	)
}

// CreateAppHandler is the handler function for the CreateApp tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func CreateAppHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "CreateApp")
}
