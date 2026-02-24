package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListDeployments tool
const listDeploymentsInputSchema = `{
  "properties": {
    "limit": {
      "default": 10,
      "description": "Maximum number of deployments to return (default: 10, max: 20)",
      "maximum": 20,
      "minimum": 1,
      "type": "integer"
    },
    "offset": {
      "default": 0,
      "description": "Number of deployments to skip (default: 0)",
      "minimum": 0,
      "type": "integer"
    },
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

// Response Template for the ListDeployments tool (Status: 200, Content-Type: application/json)
const ListDeploymentsResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Deployments retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **data** (Type: array):
      - **Items** (Type: object):
        - **extra**: Additional deployment information (URLs, etc.) (Type: object):
          - **Additional Properties**:
            - **property value** (Type: string):
        - **settings**: Build and runtime settings for VCS projects (Type: object):
          - **useBuildAI**: Whether to use Build AI for automated build configuration (Type: boolean):
              - Default: 'false'
              - Example: 'false'
          - **buildVars**: Build-time environment variables (alternative to buildFlags) (Type: object):
              - Max Length: 255
              - Example: '{"NEXT_PUBLIC_API_URL":"https://api.example.com","NODE_ENV":"production"}'
            - **Additional Properties**:
              - **property value** (Type: string):
          - **directoryPath**: Project directory path within the repository (defaults to root) (Type: string, nullable):
              - Max Length: 255
              - Nullable: true
              - Example: '.'
          - **framework**: Framework name (e.g., nextjs, reactjs-spa, reactjs-ssr). If specified, runtime must be compatible. (Type: string, nullable):
              - Max Length: 255
              - Nullable: true
              - Example: 'nextjs'
          - **runtime**: Runtime environment (e.g., node:20, golang:1.25, bun:1.3). Required if framework is not specified. (Type: string, nullable):
              - Max Length: 255
              - Nullable: true
              - Example: 'node:20'
          - **port**: Port number the application listens on (Type: integer):
              - Minimum: 1
              - Maximum: 65535
              - Example: '3000'
          - **installCommand**: Command to install dependencies (Type: string, nullable):
              - Max Length: 255
              - Nullable: true
              - Example: 'npm install'
          - **ignoreBranches**: List of exact branch names to ignore for automatic deployments (not regex patterns) (Type: array):
              - Example: '["develop","staging"]'
            - **Items** (Type: string):
                - Max Length: 255
          - **runEnvs**: Runtime environment variables (Type: object):
              - Max Length: 255
              - Example: '{"API_KEY":"secret-key","DATABASE_URL":"postgresql://localhost:5432/mydb"}'
            - **Additional Properties**:
              - **property value** (Type: string):
          - **buildDir**: Directory where build output is located (Type: string, nullable):
              - Max Length: 255
              - Nullable: true
              - Example: 'dist'
          - **hasDockerfile**: Whether the project has a Dockerfile (if true, build settings may be ignored) (Type: boolean):
              - Default: 'false'
              - Example: 'false'
          - **buildCommand**: Command to build the project (Type: string, nullable):
              - Max Length: 255
              - Nullable: true
              - Example: 'npm run build'
          - **buildFlag**: Additional flags to pass to the build command (Type: string, nullable):
              - Max Length: 255
              - Nullable: true
              - Example: '-trimpath'
          - **runFlag**: Additional flags to pass to the run command (Type: string, nullable):
              - Max Length: 255
              - Nullable: true
          - **buildFlags**: Build-time environment variables (also accepts buildVars) (Type: object):
              - Max Length: 255
              - Example: '{"NEXT_PUBLIC_API_URL":"https://api.example.com","NODE_ENV":"production"}'
            - **Additional Properties**:
              - **property value** (Type: string):
          - **runCommand**: Command to run the application (Type: string, nullable):
              - Max Length: 255
              - Nullable: true
              - Example: 'npm run start'
        - **source**: Deployment source information. For VCS projects: contains branch, commit, and commitMessage. For image projects: contains the image reference. (Type: Combinator):
          - **One Of the following structures**:
            - **Option 1**: Source for VCS deployments (Type: object):
              - **branch**: Git branch name (Type: string):
                  - Example: 'main'
              - **commit**: Git commit SHA (Type: string):
                  - Example: 'abc123def456'
              - **commitMessage**: Git commit message (Type: string):
                  - Example: 'Fix bug in authentication'
            - **Option 2**: Source for image deployments (Type: object):
              - **image**: Docker image reference (Type: string):
                  - Example: 'nginx:latest'
        - **status** (Type: string):
            - Enum: ['queue', 'building', 'deploying', 'deployed', 'build_error', 'crashing', 'cancelled', 'error', 'sleeping', 'terminating']
        - **cancelledAt** (Type: string, date-time, nullable):
            - Nullable: true
        - **createdAt** (Type: string, date-time):
        - **cancelReason** (Type: string, nullable):
            - Nullable: true
        - **projectId** (Type: string, uuid):
        - **updatedAt** (Type: string, date-time):
        - **id** (Type: string, uuid):
        - **artifactImage** (Type: string, nullable):
            - Nullable: true
    - **pagination** (Type: object):
      - **offset**: Number of items skipped (Type: integer):
          - Example: '0'
      - **total**: Total number of items available (Type: integer):
          - Example: '100'
      - **count**: Number of items in the current page (Type: integer):
          - Example: '10'
      - **limit**: Maximum number of items per page (Type: integer):
          - Example: '10'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the ListDeployments tool (Status: 400, Content-Type: application/json)
const ListDeploymentsResponseTemplate_B = `# API Response Information

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

// Response Template for the ListDeployments tool (Status: 401, Content-Type: application/json)
const ListDeploymentsResponseTemplate_C = `# API Response Information

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

// Response Template for the ListDeployments tool (Status: 403, Content-Type: application/json)
const ListDeploymentsResponseTemplate_D = `# API Response Information

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

// Response Template for the ListDeployments tool (Status: 404, Content-Type: application/json)
const ListDeploymentsResponseTemplate_E = `# API Response Information

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

// Response Template for the ListDeployments tool (Status: 422, Content-Type: application/json)
const ListDeploymentsResponseTemplate_F = `# API Response Information

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

// Response Template for the ListDeployments tool (Status: 500, Content-Type: application/json)
const ListDeploymentsResponseTemplate_G = `# API Response Information

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

// NewListDeploymentsMCPTool creates the MCP Tool instance for ListDeployments
func NewListDeploymentsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListDeployments",
		"List deployments - Returns all deployments for a project",
		[]byte(listDeploymentsInputSchema),
	)
}

// ListDeploymentsHandler is the handler function for the ListDeployments tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListDeploymentsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListDeployments")
}
