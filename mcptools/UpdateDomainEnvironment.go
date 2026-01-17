package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the UpdateDomainEnvironment tool
const updateDomainEnvironmentInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "environmentId": {
          "description": "Environment ID to assign the domain to",
          "example": "550e8400-e29b-41d4-a716-446655440000",
          "format": "uuid",
          "type": [
            "string",
            "null"
          ]
        }
      },
      "type": "object"
    },
    "domain_id": {
      "description": "Domain unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "project_id": {
      "description": "Project unique identifier",
      "format": "uuid",
      "type": "string"
    }
  },
  "required": [
    "project_id",
    "domain_id",
    "body"
  ],
  "type": "object"
}`

// Response Template for the UpdateDomainEnvironment tool (Status: 200, Content-Type: application/json)
const UpdateDomainEnvironmentResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Domain environment updated successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: string):
      - Example: 'domain environment updated successfully'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the UpdateDomainEnvironment tool (Status: 400, Content-Type: application/json)
const UpdateDomainEnvironmentResponseTemplate_B = `# API Response Information

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

// Response Template for the UpdateDomainEnvironment tool (Status: 401, Content-Type: application/json)
const UpdateDomainEnvironmentResponseTemplate_C = `# API Response Information

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

// Response Template for the UpdateDomainEnvironment tool (Status: 403, Content-Type: application/json)
const UpdateDomainEnvironmentResponseTemplate_D = `# API Response Information

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

// Response Template for the UpdateDomainEnvironment tool (Status: 404, Content-Type: application/json)
const UpdateDomainEnvironmentResponseTemplate_E = `# API Response Information

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

// Response Template for the UpdateDomainEnvironment tool (Status: 422, Content-Type: application/json)
const UpdateDomainEnvironmentResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 422

**Content-Type:** application/json

> Unprocessable Entity - domain not active, project not active, or environment not deployed

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the UpdateDomainEnvironment tool (Status: 500, Content-Type: application/json)
const UpdateDomainEnvironmentResponseTemplate_G = `# API Response Information

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

// NewUpdateDomainEnvironmentMCPTool creates the MCP Tool instance for UpdateDomainEnvironment
func NewUpdateDomainEnvironmentMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"UpdateDomainEnvironment",
		"Update domain environment - Updates the environment assignment for a domain. The domain must be in active status.",
		[]byte(updateDomainEnvironmentInputSchema),
	)
}

// UpdateDomainEnvironmentHandler is the handler function for the UpdateDomainEnvironment tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func UpdateDomainEnvironmentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "UpdateDomainEnvironment")
}
