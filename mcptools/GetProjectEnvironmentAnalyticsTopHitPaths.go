package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the GetProjectEnvironmentAnalyticsTopHitPaths tool
const getProjectEnvironmentAnalyticsTopHitPathsInputSchema = `{
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

// Response Template for the GetProjectEnvironmentAnalyticsTopHitPaths tool (Status: 200, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsTopHitPathsResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Top hit paths retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: array):
    - **Items** (Type: object):
      - **path**: Request path (Type: string):
          - Example: '/api/users'
      - **count**: Number of requests to this path (Type: number, float):
          - Example: '500'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the GetProjectEnvironmentAnalyticsTopHitPaths tool (Status: 400, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsTopHitPathsResponseTemplate_B = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalyticsTopHitPaths tool (Status: 401, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsTopHitPathsResponseTemplate_C = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalyticsTopHitPaths tool (Status: 403, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsTopHitPathsResponseTemplate_D = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalyticsTopHitPaths tool (Status: 404, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsTopHitPathsResponseTemplate_E = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalyticsTopHitPaths tool (Status: 500, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsTopHitPathsResponseTemplate_F = `# API Response Information

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

// NewGetProjectEnvironmentAnalyticsTopHitPathsMCPTool creates the MCP Tool instance for GetProjectEnvironmentAnalyticsTopHitPaths
func NewGetProjectEnvironmentAnalyticsTopHitPathsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetProjectEnvironmentAnalyticsTopHitPaths",
		"Get top 10 hit paths - Retrieves the top 10 most frequently accessed paths for a project environment.",
		[]byte(getProjectEnvironmentAnalyticsTopHitPathsInputSchema),
	)
}

// GetProjectEnvironmentAnalyticsTopHitPathsHandler is the handler function for the GetProjectEnvironmentAnalyticsTopHitPaths tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func GetProjectEnvironmentAnalyticsTopHitPathsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "GetProjectEnvironmentAnalyticsTopHitPaths")
}
