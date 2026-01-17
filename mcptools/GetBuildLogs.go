package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the GetBuildLogs tool
const getBuildLogsInputSchema = `{
  "properties": {
    "deployment_id": {
      "description": "Deployment unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "project_id": {
      "description": "Project unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "skip": {
      "default": 0,
      "description": "Number of log lines to skip",
      "minimum": 0,
      "type": "integer"
    }
  },
  "required": [
    "project_id",
    "deployment_id"
  ],
  "type": "object"
}`

// Response Template for the GetBuildLogs tool (Status: 200, Content-Type: application/json)
const GetBuildLogsResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Build logs retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: array):
    - **Items** (Type: object):
      - **createdAt** (Type: string, date-time):
      - **deploymentId** (Type: string, uuid):
      - **lineNumber** (Type: integer, uint32):
      - **log** (Type: string):
      - **stage**: Stage of the deployment (e.g., build, deploy) (Type: string):
      - **ts** (Type: string, date-time):
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the GetBuildLogs tool (Status: 400, Content-Type: application/json)
const GetBuildLogsResponseTemplate_B = `# API Response Information

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

// Response Template for the GetBuildLogs tool (Status: 401, Content-Type: application/json)
const GetBuildLogsResponseTemplate_C = `# API Response Information

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

// Response Template for the GetBuildLogs tool (Status: 403, Content-Type: application/json)
const GetBuildLogsResponseTemplate_D = `# API Response Information

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

// Response Template for the GetBuildLogs tool (Status: 404, Content-Type: application/json)
const GetBuildLogsResponseTemplate_E = `# API Response Information

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

// Response Template for the GetBuildLogs tool (Status: 422, Content-Type: application/json)
const GetBuildLogsResponseTemplate_F = `# API Response Information

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

// Response Template for the GetBuildLogs tool (Status: 500, Content-Type: application/json)
const GetBuildLogsResponseTemplate_G = `# API Response Information

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

// NewGetBuildLogsMCPTool creates the MCP Tool instance for GetBuildLogs
func NewGetBuildLogsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetBuildLogs",
		"Get build logs - Retrieves build logs for a deployment",
		[]byte(getBuildLogsInputSchema),
	)
}

// GetBuildLogsHandler is the handler function for the GetBuildLogs tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func GetBuildLogsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "GetBuildLogs")
}
