package mcptools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListProjects tool
const listProjectsInputSchema = `{
  "properties": {
    "limit": {
      "default": 10,
      "description": "Maximum number of projects to return (default: 10, max: 20)",
      "maximum": 20,
      "minimum": 1,
      "type": "integer"
    },
    "offset": {
      "default": 0,
      "description": "Number of projects to skip (default: 0)",
      "minimum": 0,
      "type": "integer"
    },
    "show-deleted": {
      "default": false,
      "description": "Include deleted projects in the results",
      "type": "boolean"
    }
  },
  "type": "object"
}`

// Response Template for the ListProjects tool (Status: 200, Content-Type: application/json)
const ListProjectsResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> List of projects retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **data** (Type: array):
      - **Items** (Type: object):
        - **id** (Type: string, uuid):
        - **createdAt** (Type: string, date-time):
        - **description** (Type: string, nullable):
            - Nullable: true
        - **source** (Type: object):
        - **deletedAt** (Type: string, date-time, nullable):
            - Nullable: true
        - **settings**: Build and runtime settings for VCS projects (Type: object):
          - **installCommand**: Command to install dependencies (Type: string, nullable):
              - Max Length: 255
              - Nullable: true
              - Example: 'npm install'
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
          - **runFlag**: Additional flags to pass to the run command (Type: string, nullable):
              - Max Length: 255
              - Nullable: true
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
          - **framework**: Framework name (e.g., nextjs, reactjs-spa, reactjs-ssr). If specified, runtime must be compatible. (Type: string, nullable):
              - Max Length: 255
              - Nullable: true
              - Example: 'nextjs'
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
          - **buildCommand**: Command to build the project (Type: string, nullable):
              - Max Length: 255
              - Nullable: true
              - Example: 'npm run build'
        - **type** (Type: string):
            - Enum: ['vcs', 'image']
        - **uniqueName** (Type: string):
        - **displayName** (Type: string):
        - **status** (Type: string):
            - Enum: ['active', 'deleting', 'deleted']
        - **updatedAt** (Type: string, date-time):
        - **userId** (Type: string):
    - **pagination** (Type: object):
      - **count**: Number of items in the current page (Type: integer):
          - Example: '10'
      - **limit**: Maximum number of items per page (Type: integer):
          - Example: '10'
      - **offset**: Number of items skipped (Type: integer):
          - Example: '0'
      - **total**: Total number of items available (Type: integer):
          - Example: '100'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the ListProjects tool (Status: 401, Content-Type: application/json)
const ListProjectsResponseTemplate_B = `# API Response Information

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

// Response Template for the ListProjects tool (Status: 500, Content-Type: application/json)
const ListProjectsResponseTemplate_C = `# API Response Information

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

// NewListProjectsMCPTool creates the MCP Tool instance for ListProjects
func NewListProjectsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListProjects",
		"List all projects - Returns a list of all projects for the authenticated user",
		[]byte(listProjectsInputSchema),
	)
}

// ListProjectsHandler is the handler function for the ListProjects tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListProjectsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListProjects")
}
