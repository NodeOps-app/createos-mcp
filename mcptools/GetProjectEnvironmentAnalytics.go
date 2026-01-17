package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the GetProjectEnvironmentAnalytics tool
const getProjectEnvironmentAnalyticsInputSchema = `{
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

// Response Template for the GetProjectEnvironmentAnalytics tool (Status: 200, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Analytics retrieved successfully

## Response Structure

- Structure (Type: object):
  - **overallRequests** (Type: object):
    - **total**: Total number of requests (Type: number, float):
        - Example: '1250'
    - **client**: Number of client error requests (4xx status codes) (Type: number, float):
        - Example: '100'
    - **server**: Number of server error requests (5xx status codes) (Type: number, float):
        - Example: '50'
    - **success**: Number of successful requests (2xx status codes) (Type: number, float):
        - Example: '1100'
  - **requestDistribution** (Type: array):
    - **Items** (Type: object):
      - **count**: Number of requests with this status code (Type: number, float):
          - Example: '1000'
      - **status**: HTTP status code (Type: string):
          - Example: '200'
  - **rpm** (Type: object):
    - **average**: Average requests per minute (Type: number, float):
        - Example: '20.8'
    - **peak**: Peak requests per minute (Type: number, float):
        - Example: '25.5'
  - **successPercentage**: Success percentage (2xx + 3xx responses) (Type: number, float):
      - Example: '88'
  - **topErrorPaths** (Type: array):
    - **Items** (Type: object):
      - **count**: Number of error requests to this path (Type: number, float):
          - Example: '20'
      - **path**: Request path (Type: string):
          - Example: '/api/invalid'
      - **status**: HTTP status code (Type: string):
          - Example: '404'
  - **topHitPaths** (Type: array):
    - **Items** (Type: object):
      - **count**: Number of requests to this path (Type: number, float):
          - Example: '500'
      - **path**: Request path (Type: string):
          - Example: '/api/users'
`

// Response Template for the GetProjectEnvironmentAnalytics tool (Status: 400, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsResponseTemplate_B = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalytics tool (Status: 401, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsResponseTemplate_C = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalytics tool (Status: 403, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsResponseTemplate_D = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalytics tool (Status: 404, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsResponseTemplate_E = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalytics tool (Status: 500, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsResponseTemplate_F = `# API Response Information

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

// NewGetProjectEnvironmentAnalyticsMCPTool creates the MCP Tool instance for GetProjectEnvironmentAnalytics
func NewGetProjectEnvironmentAnalyticsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetProjectEnvironmentAnalytics",
		"Get environment analytics - Retrieves comprehensive analytics data for a project environment including overall requests, request distribution, RPM, success percentage, top hit paths, and top error paths. Supports custom time ranges via start and end parameters.",
		[]byte(getProjectEnvironmentAnalyticsInputSchema),
	)
}

// GetProjectEnvironmentAnalyticsHandler is the handler function for the GetProjectEnvironmentAnalytics tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func GetProjectEnvironmentAnalyticsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "GetProjectEnvironmentAnalytics")
}
