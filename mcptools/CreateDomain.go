package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the CreateDomain tool
const createDomainInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "environmentId": {
          "description": "Optional environment ID to assign the domain to",
          "example": "550e8400-e29b-41d4-a716-446655440000",
          "format": "uuid",
          "type": [
            "string",
            "null"
          ]
        },
        "name": {
          "description": "Domain name (must be a valid domain)",
          "example": "example.com",
          "maxLength": 255,
          "minLength": 3,
          "type": "string"
        }
      },
      "required": [
        "name"
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

// Response Template for the CreateDomain tool (Status: 201, Content-Type: application/json)
const CreateDomainResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 201

**Content-Type:** application/json

> Domain created successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **id** (Type: string, uuid):
        - Example: '550e8400-e29b-41d4-a716-446655440000'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the CreateDomain tool (Status: 400, Content-Type: application/json)
const CreateDomainResponseTemplate_B = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 400

**Content-Type:** application/json

> Bad request - invalid domain name or validation error

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the CreateDomain tool (Status: 401, Content-Type: application/json)
const CreateDomainResponseTemplate_C = `# API Response Information

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

// Response Template for the CreateDomain tool (Status: 403, Content-Type: application/json)
const CreateDomainResponseTemplate_D = `# API Response Information

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

// Response Template for the CreateDomain tool (Status: 404, Content-Type: application/json)
const CreateDomainResponseTemplate_E = `# API Response Information

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

// Response Template for the CreateDomain tool (Status: 422, Content-Type: application/json)
const CreateDomainResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 422

**Content-Type:** application/json

> Unprocessable Entity - domain limit reached, duplicate domain, or project not active

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the CreateDomain tool (Status: 500, Content-Type: application/json)
const CreateDomainResponseTemplate_G = `# API Response Information

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

// NewCreateDomainMCPTool creates the MCP Tool instance for CreateDomain
func NewCreateDomainMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"CreateDomain",
		"Create domain - Adds a custom domain to a project. The domain must be unique and not already in use by another project.",
		[]byte(createDomainInputSchema),
	)
}

// CreateDomainHandler is the handler function for the CreateDomain tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func CreateDomainHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "CreateDomain")
}
