package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the GetProjectEnvironmentAnalyticsRequestsOverTime tool
const getProjectEnvironmentAnalyticsRequestsOverTimeInputSchema = `{
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

// Response Template for the GetProjectEnvironmentAnalyticsRequestsOverTime tool (Status: 200, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsRequestsOverTimeResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Requests over time data retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: array):
    - **Items** (Type: object):
      - **timestamp**: Unix timestamp in seconds (Type: integer, int64):
          - Example: '1704067200'
      - **count**: Total number of requests at this timestamp (Type: integer):
          - Example: '150'
      - **status2xx**: Number of 2xx status code responses (Type: integer):
          - Example: '120'
      - **status3xx**: Number of 3xx status code responses (Type: integer):
          - Example: '10'
      - **status4xx**: Number of 4xx status code responses (Type: integer):
          - Example: '15'
      - **status5xx**: Number of 5xx status code responses (Type: integer):
          - Example: '5'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the GetProjectEnvironmentAnalyticsRequestsOverTime tool (Status: 400, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsRequestsOverTimeResponseTemplate_B = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalyticsRequestsOverTime tool (Status: 401, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsRequestsOverTimeResponseTemplate_C = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalyticsRequestsOverTime tool (Status: 403, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsRequestsOverTimeResponseTemplate_D = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalyticsRequestsOverTime tool (Status: 404, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsRequestsOverTimeResponseTemplate_E = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalyticsRequestsOverTime tool (Status: 422, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsRequestsOverTimeResponseTemplate_F = `# API Response Information

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

// Response Template for the GetProjectEnvironmentAnalyticsRequestsOverTime tool (Status: 500, Content-Type: application/json)
const GetProjectEnvironmentAnalyticsRequestsOverTimeResponseTemplate_G = `# API Response Information

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

// NewGetProjectEnvironmentAnalyticsRequestsOverTimeMCPTool creates the MCP Tool instance for GetProjectEnvironmentAnalyticsRequestsOverTime
func NewGetProjectEnvironmentAnalyticsRequestsOverTimeMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetProjectEnvironmentAnalyticsRequestsOverTime",
		"Get requests over time - Retrieves time series data for requests broken down by status code ranges (2xx, 3xx, 4xx, 5xx) for a project environment.",
		[]byte(getProjectEnvironmentAnalyticsRequestsOverTimeInputSchema),
	)
}

// GetProjectEnvironmentAnalyticsRequestsOverTimeHandler is the handler function for the GetProjectEnvironmentAnalyticsRequestsOverTime tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func GetProjectEnvironmentAnalyticsRequestsOverTimeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "GetProjectEnvironmentAnalyticsRequestsOverTime")
}
