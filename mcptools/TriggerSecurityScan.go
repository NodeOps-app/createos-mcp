package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the TriggerSecurityScan tool
const triggerSecurityScanInputSchema = `{
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
    }
  },
  "required": [
    "project_id",
    "deployment_id"
  ],
  "type": "object"
}`

// Response Template for the TriggerSecurityScan tool (Status: 200, Content-Type: application/json)
const TriggerSecurityScanResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Security scan triggered successfully

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'true'
  - **data** (Type: string):
      - Example: 'security scan report created and triggered'
`

// Response Template for the TriggerSecurityScan tool (Status: 400, Content-Type: application/json)
const TriggerSecurityScanResponseTemplate_B = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 400

**Content-Type:** application/json

> Bad request - security scan not enabled for project

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the TriggerSecurityScan tool (Status: 401, Content-Type: application/json)
const TriggerSecurityScanResponseTemplate_C = `# API Response Information

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

// Response Template for the TriggerSecurityScan tool (Status: 403, Content-Type: application/json)
const TriggerSecurityScanResponseTemplate_D = `# API Response Information

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

// Response Template for the TriggerSecurityScan tool (Status: 404, Content-Type: application/json)
const TriggerSecurityScanResponseTemplate_E = `# API Response Information

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

// Response Template for the TriggerSecurityScan tool (Status: 500, Content-Type: application/json)
const TriggerSecurityScanResponseTemplate_F = `# API Response Information

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

// NewTriggerSecurityScanMCPTool creates the MCP Tool instance for TriggerSecurityScan
func NewTriggerSecurityScanMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"TriggerSecurityScan",
		"Trigger security scan - Creates and triggers a new security scan for a deployment. Security scanning must be enabled for the project.",
		[]byte(triggerSecurityScanInputSchema),
	)
}

// TriggerSecurityScanHandler is the handler function for the TriggerSecurityScan tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func TriggerSecurityScanHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "TriggerSecurityScan")
}
