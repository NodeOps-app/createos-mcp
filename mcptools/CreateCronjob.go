package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the CreateCronjob tool
const createCronjobInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "environmentId": {
          "description": "Environment to run the cronjob in",
          "format": "uuid",
          "type": "string"
        },
        "name": {
          "description": "Name of the cronjob",
          "example": "health-check",
          "maxLength": 255,
          "minLength": 1,
          "type": "string"
        },
        "schedule": {
          "description": "Cron schedule expression",
          "example": "*/5 * * * *",
          "maxLength": 256,
          "type": "string"
        },
        "settings": {
          "description": "Cronjob type-specific settings",
          "oneOf": [
            {
              "properties": {
                "body": {
                  "description": "Request body",
                  "maxLength": 65536,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "headers": {
                  "additionalProperties": {
                    "type": "string"
                  },
                  "description": "HTTP headers to include in the request",
                  "example": {
                    "Authorization": "Bearer token"
                  },
                  "type": "object"
                },
                "method": {
                  "description": "HTTP method",
                  "enum": [
                    "GET",
                    "POST",
                    "PUT",
                    "DELETE",
                    "PATCH",
                    "HEAD"
                  ],
                  "example": "GET",
                  "type": "string"
                },
                "path": {
                  "description": "Request path (must start with /)",
                  "example": "/api/health",
                  "maxLength": 2048,
                  "type": "string"
                },
                "timeoutInSeconds": {
                  "description": "Request timeout in seconds",
                  "example": 5,
                  "maximum": 5,
                  "minimum": 1,
                  "type": "integer"
                }
              },
              "required": [
                "path",
                "method"
              ],
              "type": "object"
            }
          ]
        },
        "type": {
          "description": "Type of the cronjob",
          "enum": [
            "http",
            "artifact"
          ],
          "example": "http",
          "type": "string"
        }
      },
      "required": [
        "name",
        "schedule",
        "type",
        "environmentId"
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

// Response Template for the CreateCronjob tool (Status: 200, Content-Type: application/json)
const CreateCronjobResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Cronjob created successfully

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'true'
  - **data** (Type: object):
    - **id**: Created resource ID (e.g. purchase ID) (Type: string, uuid):
`

// Response Template for the CreateCronjob tool (Status: 400, Content-Type: application/json)
const CreateCronjobResponseTemplate_B = `# API Response Information

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

// Response Template for the CreateCronjob tool (Status: 401, Content-Type: application/json)
const CreateCronjobResponseTemplate_C = `# API Response Information

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

// Response Template for the CreateCronjob tool (Status: 403, Content-Type: application/json)
const CreateCronjobResponseTemplate_D = `# API Response Information

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

// Response Template for the CreateCronjob tool (Status: 404, Content-Type: application/json)
const CreateCronjobResponseTemplate_E = `# API Response Information

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

// Response Template for the CreateCronjob tool (Status: 422, Content-Type: application/json)
const CreateCronjobResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 422

**Content-Type:** application/json

> Unprocessable Entity - cronjob limit reached or environment not active

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the CreateCronjob tool (Status: 500, Content-Type: application/json)
const CreateCronjobResponseTemplate_G = `# API Response Information

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

// NewCreateCronjobMCPTool creates the MCP Tool instance for CreateCronjob
func NewCreateCronjobMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"CreateCronjob",
		"Create a cronjob - Creates a new cronjob for a project. The cronjob can be of type http or artifact. HTTP type cronjobs call a configured HTTP endpoint on the specified schedule. Maximum of 4 cronjobs per project.",
		[]byte(createCronjobInputSchema),
	)
}

// CreateCronjobHandler is the handler function for the CreateCronjob tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func CreateCronjobHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "CreateCronjob")
}
