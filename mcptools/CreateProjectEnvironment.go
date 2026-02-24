package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the CreateProjectEnvironment tool
const createProjectEnvironmentInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "branch": {
          "description": "Git branch name for VCS projects (required for VCS projects, must exist in repository). For image projects, this field should be omitted or null.",
          "example": "main",
          "maxLength": 255,
          "minLength": 1,
          "type": [
            "string",
            "null"
          ]
        },
        "description": {
          "example": "Production environment",
          "maxLength": 2048,
          "minLength": 4,
          "type": [
            "string",
            "null"
          ]
        },
        "displayName": {
          "description": "Display name for the environment",
          "example": "Production",
          "maxLength": 48,
          "minLength": 4,
          "pattern": "^[a-zA-Z0-9 _-]+$",
          "type": "string"
        },
        "isAutoPromoteEnabled": {
          "default": false,
          "description": "Enable automatic promotion of deployments",
          "example": false,
          "type": "boolean"
        },
        "resources": {
          "properties": {
            "cpu": {
              "description": "CPU in millicores",
              "example": 500,
              "maximum": 500,
              "minimum": 200,
              "type": "integer"
            },
            "memory": {
              "description": "Memory in MB",
              "example": 1024,
              "maximum": 1024,
              "minimum": 500,
              "type": "integer"
            },
            "replicas": {
              "description": "Number of replicas",
              "example": 2,
              "maximum": 3,
              "minimum": 1,
              "type": "integer"
            }
          },
          "required": [
            "cpu",
            "memory",
            "replicas"
          ],
          "type": "object"
        },
        "settings": {
          "description": "Environment settings (runtime environment variables)",
          "properties": {
            "runEnvs": {
              "additionalProperties": {
                "type": "string"
              },
              "example": {
                "API_KEY": "prod-key",
                "DATABASE_URL": "postgresql://prod-db:5432/mydb"
              },
              "type": "object"
            }
          },
          "type": "object"
        },
        "uniqueName": {
          "example": "production",
          "maxLength": 32,
          "minLength": 4,
          "pattern": "^[a-zA-Z0-9-]+$",
          "type": "string"
        }
      },
      "required": [
        "displayName",
        "uniqueName",
        "description",
        "settings",
        "resources"
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

// Response Template for the CreateProjectEnvironment tool (Status: 201, Content-Type: application/json)
const CreateProjectEnvironmentResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 201

**Content-Type:** application/json

> Environment created successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **id**: Unique identifier of the created project (Type: string, uuid):
        - Example: '550e8400-e29b-41d4-a716-446655440000'
  - **status** (Type: string):
      - Example: 'success'
`

// Response Template for the CreateProjectEnvironment tool (Status: 400, Content-Type: application/json)
const CreateProjectEnvironmentResponseTemplate_B = `# API Response Information

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

// Response Template for the CreateProjectEnvironment tool (Status: 401, Content-Type: application/json)
const CreateProjectEnvironmentResponseTemplate_C = `# API Response Information

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

// Response Template for the CreateProjectEnvironment tool (Status: 403, Content-Type: application/json)
const CreateProjectEnvironmentResponseTemplate_D = `# API Response Information

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

// Response Template for the CreateProjectEnvironment tool (Status: 404, Content-Type: application/json)
const CreateProjectEnvironmentResponseTemplate_E = `# API Response Information

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

// Response Template for the CreateProjectEnvironment tool (Status: 409, Content-Type: application/json)
const CreateProjectEnvironmentResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 409

**Content-Type:** application/json

> Conflict - environment name already exists

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the CreateProjectEnvironment tool (Status: 422, Content-Type: application/json)
const CreateProjectEnvironmentResponseTemplate_G = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 422

**Content-Type:** application/json

> Unprocessable Entity - project not active

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the CreateProjectEnvironment tool (Status: 429, Content-Type: application/json)
const CreateProjectEnvironmentResponseTemplate_H = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 429

**Content-Type:** application/json

> Too many requests - maximum number of environments reached for the project

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the CreateProjectEnvironment tool (Status: 500, Content-Type: application/json)
const CreateProjectEnvironmentResponseTemplate_I = `# API Response Information

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

// NewCreateProjectEnvironmentMCPTool creates the MCP Tool instance for CreateProjectEnvironment
func NewCreateProjectEnvironmentMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"CreateProjectEnvironment",
		"Create project environment - Creates a new environment for a project. For VCS projects: branch is required and must exist in the repository. For image projects: branch should be omitted or null.",
		[]byte(createProjectEnvironmentInputSchema),
	)
}

// CreateProjectEnvironmentHandler is the handler function for the CreateProjectEnvironment tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func CreateProjectEnvironmentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "CreateProjectEnvironment")
}
