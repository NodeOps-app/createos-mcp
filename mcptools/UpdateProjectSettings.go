package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the UpdateProjectSettings tool
const updateProjectSettingsInputSchema = `{
  "properties": {
    "body": {
      "oneOf": [
        {
          "description": "Build and runtime settings for VCS projects",
          "properties": {
            "buildCommand": {
              "description": "Command to build the project",
              "example": "npm run build",
              "maxLength": 255,
              "type": [
                "string",
                "null"
              ]
            },
            "buildDir": {
              "description": "Directory where build output is located",
              "example": "dist",
              "maxLength": 255,
              "type": [
                "string",
                "null"
              ]
            },
            "buildFlag": {
              "description": "Additional flags to pass to the build command",
              "example": "-trimpath",
              "maxLength": 255,
              "type": [
                "string",
                "null"
              ]
            },
            "buildFlags": {
              "additionalProperties": {
                "type": "string"
              },
              "description": "Build-time environment variables (also accepts buildVars)",
              "example": {
                "NEXT_PUBLIC_API_URL": "https://api.example.com",
                "NODE_ENV": "production"
              },
              "type": "object"
            },
            "buildVars": {
              "additionalProperties": {
                "type": "string"
              },
              "description": "Build-time environment variables (alternative to buildFlags)",
              "example": {
                "NEXT_PUBLIC_API_URL": "https://api.example.com",
                "NODE_ENV": "production"
              },
              "type": "object"
            },
            "directoryPath": {
              "description": "Project directory path within the repository (defaults to root)",
              "example": ".",
              "maxLength": 255,
              "type": [
                "string",
                "null"
              ]
            },
            "framework": {
              "description": "Framework name (e.g., nextjs, reactjs-spa, reactjs-ssr). If specified, runtime must be compatible.",
              "example": "nextjs",
              "maxLength": 255,
              "type": [
                "string",
                "null"
              ]
            },
            "hasDockerfile": {
              "default": false,
              "description": "Whether the project has a Dockerfile (if true, build settings may be ignored)",
              "example": false,
              "type": "boolean"
            },
            "ignoreBranches": {
              "description": "List of exact branch names to ignore for automatic deployments (not regex patterns)",
              "example": [
                "develop",
                "staging"
              ],
              "items": {
                "maxLength": 255,
                "type": "string"
              },
              "type": "array"
            },
            "installCommand": {
              "description": "Command to install dependencies",
              "example": "npm install",
              "maxLength": 255,
              "type": [
                "string",
                "null"
              ]
            },
            "port": {
              "description": "Port number the application listens on",
              "example": 3000,
              "maximum": 65535,
              "minimum": 1,
              "type": "integer"
            },
            "runCommand": {
              "description": "Command to run the application",
              "example": "npm run start",
              "maxLength": 255,
              "type": [
                "string",
                "null"
              ]
            },
            "runEnvs": {
              "additionalProperties": {
                "type": "string"
              },
              "description": "Runtime environment variables",
              "example": {
                "API_KEY": "secret-key",
                "DATABASE_URL": "postgresql://localhost:5432/mydb"
              },
              "type": "object"
            },
            "runFlag": {
              "description": "Additional flags to pass to the run command",
              "maxLength": 255,
              "type": [
                "string",
                "null"
              ]
            },
            "runtime": {
              "description": "Runtime environment (e.g., node:20, golang:1.25, bun:1.3). Required if framework is not specified.",
              "example": "node:20",
              "maxLength": 255,
              "type": [
                "string",
                "null"
              ]
            },
            "useBuildAI": {
              "default": false,
              "description": "Whether to use Build AI for automated build configuration",
              "example": false,
              "type": "boolean"
            }
          },
          "type": "object"
        },
        {
          "additionalProperties": false,
          "description": "Settings for image type projects",
          "not": {
            "properties": {
              "buildCommand": {
                "type": "string"
              },
              "directoryPath": {
                "type": "string"
              },
              "framework": {
                "type": "string"
              },
              "hasDockerfile": {
                "type": "boolean"
              },
              "ignoreBranches": {
                "items": {
                  "type": "string"
                },
                "type": "array"
              },
              "installCommand": {
                "type": "string"
              },
              "runtime": {
                "type": "string"
              },
              "useBuildAI": {
                "type": "boolean"
              }
            },
            "type": "object"
          },
          "properties": {
            "port": {
              "description": "Port number the application listens on",
              "example": 8080,
              "maximum": 65535,
              "minimum": 1,
              "type": "integer"
            },
            "runEnvs": {
              "additionalProperties": {
                "type": "string"
              },
              "description": "Runtime environment variables",
              "example": {
                "API_KEY": "secret-key",
                "DATABASE_URL": "postgresql://localhost:5432/mydb"
              },
              "type": "object"
            }
          },
          "required": [
            "port"
          ],
          "type": "object"
        }
      ]
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

// Response Template for the UpdateProjectSettings tool (Status: 200, Content-Type: application/json)
const UpdateProjectSettingsResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Project settings updated successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: string):
      - Example: 'project settings updated successfully'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the UpdateProjectSettings tool (Status: 400, Content-Type: application/json)
const UpdateProjectSettingsResponseTemplate_B = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 400

**Content-Type:** application/json

> Bad request - invalid settings, validation error, or project not in active state

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the UpdateProjectSettings tool (Status: 401, Content-Type: application/json)
const UpdateProjectSettingsResponseTemplate_C = `# API Response Information

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

// Response Template for the UpdateProjectSettings tool (Status: 403, Content-Type: application/json)
const UpdateProjectSettingsResponseTemplate_D = `# API Response Information

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

// Response Template for the UpdateProjectSettings tool (Status: 404, Content-Type: application/json)
const UpdateProjectSettingsResponseTemplate_E = `# API Response Information

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

// Response Template for the UpdateProjectSettings tool (Status: 500, Content-Type: application/json)
const UpdateProjectSettingsResponseTemplate_F = `# API Response Information

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

// NewUpdateProjectSettingsMCPTool creates the MCP Tool instance for UpdateProjectSettings
func NewUpdateProjectSettingsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"UpdateProjectSettings",
		"Update project settings - Updates build and runtime settings for a project. For VCS projects: accepts VCSSettings schema with all build/runtime configuration options. For image projects: accepts ImageSettings schema with port and runtime environment variables.",
		[]byte(updateProjectSettingsInputSchema),
	)
}

// UpdateProjectSettingsHandler is the handler function for the UpdateProjectSettings tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func UpdateProjectSettingsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "UpdateProjectSettings")
}
