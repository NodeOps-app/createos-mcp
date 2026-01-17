package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the GetProjectEnvironmentAnalyticsTopErrorPaths tool
const getProjectEnvironmentAnalyticsTopErrorPathsInputSchema = `{
  "properties": {
    "end": {
      "description": "End timestamp (Unix timestamp in seconds). Defaults to current time.",
      "example": 1704070800,
      "format": "int64",
      "type": "integer"
    },
    "environment_id": {
      "description": "Environment unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "project_id": {
      "description": "Project unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "start": {
      "description": "Start timestamp (Unix timestamp in seconds). Defaults to 1 hour ago.",
      "example": 1704067200,
      "format": "int64",
      "type": "integer"
    }
  },
  "required": [
    "project_id",
    "environment_id"
  ],
  "type": "object"
}`

// Response Template for the GetProjectEnvironmentAnalyticsTopErrorPaths tool (Status: 200, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsTopErrorPathsResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Top error paths retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: array):
    - **Items** (Type: object):
      - **status**: HTTP status code (Type: string):
          - Example: '404'
      - **count**: Number of error requests to this path (Type: number, float):
          - Example: '20'
      - **path**: Request path (Type: string):
          - Example: '/api/invalid'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the GetProjectEnvironmentAnalyticsTopErrorPaths tool (Status: 400, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsTopErrorPathsResponseTemplate_B = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalyticsTopErrorPaths tool (Status: 401, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsTopErrorPathsResponseTemplate_C = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalyticsTopErrorPaths tool (Status: 403, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsTopErrorPathsResponseTemplate_D = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalyticsTopErrorPaths tool (Status: 404, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsTopErrorPathsResponseTemplate_E = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalyticsTopErrorPaths tool (Status: 500, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsTopErrorPathsResponseTemplate_F = `# API Response Information

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

// NewGetProjectEnvironmentAnalyticsTopErrorPathsMCPTool creates the MCP Tool instance for GetProjectEnvironmentAnalyticsTopErrorPaths
func NewGetProjectEnvironmentAnalyticsTopErrorPathsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetProjectEnvironmentAnalyticsTopErrorPaths",
		"Get top 10 error paths - Retrieves the top 10 paths with the most errors (4xx and 5xx status codes) for a project environment.",
		[]byte(getProjectEnvironmentAnalyticsTopErrorPathsInputSchema),
	)
}

// GetProjectEnvironmentAnalyticsTopErrorPathsHandler is the handler function for the GetProjectEnvironmentAnalyticsTopErrorPaths tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func GetProjectEnvironmentAnalyticsTopErrorPathsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "GetProjectEnvironmentAnalyticsTopErrorPaths")
}
