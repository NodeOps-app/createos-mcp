package mcptools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the GetProject tool
const getProjectInputSchema = `{
  "properties": {
    "project_id": {
      "description": "Project unique identifier",
      "format": "uuid",
      "type": "string"
    }
  },
  "required": [
    "project_id"
  ],
  "type": "object"
}`

// Response Template for the GetProject tool (Status: 200, Content-Type: application/json)
const GetProjectResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Project retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **id** (Type: string, uuid):
    - **createdAt** (Type: string, date-time):
    - **description** (Type: string, nullable):
        - Nullable: true
    - **source** (Type: object):
    - **settings**: Build and runtime settings for VCS projects (Type: object):
      - **runtime**: Runtime environment (e.g., node:20, golang:1.25, bun:1.3). Required if framework is not specified. (Type: string, nullable):
          - Max Length: 255
          - Nullable: true
          - Example: 'node:20'
      - **directoryPath**: Project directory path within the repository (defaults to root) (Type: string, nullable):
          - Max Length: 255
          - Nullable: true
          - Example: '.'
      - **ignoreBranches**: List of exact branch names to ignore for automatic deployments (not regex patterns) (Type: array):
          - Example: '["develop","staging"]'
        - **Items** (Type: string):
            - Max Length: 255
      - **buildVars**: Build-time environment variables (alternative to buildFlags) (Type: object):
          - Max Length: 255
          - Example: '{"NEXT_PUBLIC_API_URL":"https://api.example.com","NODE_ENV":"production"}'
        - **Additional Properties**:
          - **property value** (Type: string):
      - **runCommand**: Command to run the application (Type: string, nullable):
          - Max Length: 255
          - Nullable: true
          - Example: 'npm run start'
      - **framework**: Framework name (e.g., nextjs, reactjs-spa, reactjs-ssr). If specified, runtime must be compatible. (Type: string, nullable):
          - Max Length: 255
          - Nullable: true
          - Example: 'nextjs'
      - **buildCommand**: Command to build the project (Type: string, nullable):
          - Max Length: 255
          - Nullable: true
          - Example: 'npm run build'
      - **buildFlags**: Build-time environment variables (also accepts buildVars) (Type: object):
          - Max Length: 255
          - Example: '{"NEXT_PUBLIC_API_URL":"https://api.example.com","NODE_ENV":"production"}'
        - **Additional Properties**:
          - **property value** (Type: string):
      - **runEnvs**: Runtime environment variables (Type: object):
          - Max Length: 255
          - Example: '{"API_KEY":"secret-key","DATABASE_URL":"postgresql://localhost:5432/mydb"}'
        - **Additional Properties**:
          - **property value** (Type: string):
      - **installCommand**: Command to install dependencies (Type: string, nullable):
          - Max Length: 255
          - Nullable: true
          - Example: 'npm install'
      - **useBuildAI**: Whether to use Build AI for automated build configuration (Type: boolean):
          - Default: 'false'
          - Example: 'false'
      - **hasDockerfile**: Whether the project has a Dockerfile (if true, build settings may be ignored) (Type: boolean):
          - Default: 'false'
          - Example: 'false'
      - **buildFlag**: Additional flags to pass to the build command (Type: string, nullable):
          - Max Length: 255
          - Nullable: true
          - Example: '-trimpath'
      - **port**: Port number the application listens on (Type: integer):
          - Minimum: 1
          - Maximum: 65535
          - Example: '3000'
      - **buildDir**: Directory where build output is located (Type: string, nullable):
          - Max Length: 255
          - Nullable: true
          - Example: 'dist'
      - **runFlag**: Additional flags to pass to the run command (Type: string, nullable):
          - Max Length: 255
          - Nullable: true
    - **type** (Type: string):
        - Enum: ['vcs', 'image']
    - **uniqueName** (Type: string):
    - **displayName** (Type: string):
    - **status** (Type: string):
        - Enum: ['active', 'deleting', 'deleted']
    - **updatedAt** (Type: string, date-time):
    - **userId** (Type: string):
    - **deletedAt** (Type: string, date-time, nullable):
        - Nullable: true
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the GetProject tool (Status: 400, Content-Type: application/json)
const GetProjectResponseTemplate_B = `# API Response Information

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

// Response Template for the GetProject tool (Status: 401, Content-Type: application/json)
const GetProjectResponseTemplate_C = `# API Response Information

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

// Response Template for the GetProject tool (Status: 403, Content-Type: application/json)
const GetProjectResponseTemplate_D = `# API Response Information

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

// Response Template for the GetProject tool (Status: 404, Content-Type: application/json)
const GetProjectResponseTemplate_E = `# API Response Information

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

// Response Template for the GetProject tool (Status: 500, Content-Type: application/json)
const GetProjectResponseTemplate_F = `# API Response Information

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

// NewGetProjectMCPTool creates the MCP Tool instance for GetProject
func NewGetProjectMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetProject",
		"Get project by ID - Retrieves detailed information about a specific project",
		[]byte(getProjectInputSchema),
	)
}

// GetProjectHandler is the handler function for the GetProject tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func GetProjectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "GetProject")
}
