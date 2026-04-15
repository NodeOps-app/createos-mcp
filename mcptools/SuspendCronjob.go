package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the SuspendCronjob tool
const suspendCronjobInputSchema = `{
  "properties": {
    "cronjob_id": {
      "description": "Cronjob unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "project_id": {
      "description": "Project unique identifier",
      "format": "uuid",
      "type": "string"
    }
  },
  "required": [
    "project_id",
    "cronjob_id"
  ],
  "type": "object"
}`

// Response Template for the SuspendCronjob tool (Status: 200, Content-Type: application/json)
const SuspendCronjobResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Cronjob suspended successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: string):
      - Example: 'cronjob suspended'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the SuspendCronjob tool (Status: 401, Content-Type: application/json)
const SuspendCronjobResponseTemplate_B = `# API Response Information

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

// Response Template for the SuspendCronjob tool (Status: 403, Content-Type: application/json)
const SuspendCronjobResponseTemplate_C = `# API Response Information

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

// Response Template for the SuspendCronjob tool (Status: 404, Content-Type: application/json)
const SuspendCronjobResponseTemplate_D = `# API Response Information

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

// Response Template for the SuspendCronjob tool (Status: 422, Content-Type: application/json)
const SuspendCronjobResponseTemplate_E = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 422

**Content-Type:** application/json

> Unprocessable Entity - cronjob is already suspended

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the SuspendCronjob tool (Status: 500, Content-Type: application/json)
const SuspendCronjobResponseTemplate_F = `# API Response Information

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

// NewSuspendCronjobMCPTool creates the MCP Tool instance for SuspendCronjob
func NewSuspendCronjobMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"SuspendCronjob",
		"Suspend a cronjob - Suspends a cronjob, pausing its scheduled executions.",
		[]byte(suspendCronjobInputSchema),
	)
}

// SuspendCronjobHandler is the handler function for the SuspendCronjob tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func SuspendCronjobHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "SuspendCronjob")
}
