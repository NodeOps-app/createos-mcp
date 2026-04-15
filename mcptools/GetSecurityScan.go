package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the GetSecurityScan tool
const getSecurityScanInputSchema = `{
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

// Response Template for the GetSecurityScan tool (Status: 200, Content-Type: application/json)
const GetSecurityScanResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Security scan report retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **type**: Type of security scan (Type: string):
        - Example: 'vulnerability'
    - **deploymentId**: Deployment ID this report belongs to (Type: string, uuid):
        - Example: '550e8400-e29b-41d4-a716-446655440001'
    - **parsedReport**: Parsed security scan results (Type: object, nullable):
        - Nullable: true
      - **Allows Additional Properties**
    - **deletedAt**: When the report was deleted (Type: string, date-time, nullable):
        - Nullable: true
    - **failedReason**: Reason for failure if status is failed (Type: string, nullable):
        - Nullable: true
    - **id**: Security scan report ID (Type: string, uuid):
        - Example: '550e8400-e29b-41d4-a716-446655440000'
    - **tool**: Security scanning tool used (Type: string):
        - Example: 'trivy'
    - **expiresAt**: When the scan job times out (Type: string, date-time):
        - Example: '2025-01-18T12:00:00Z'
    - **executor**: Worker that executed the scan (Type: string, nullable):
        - Nullable: true
        - Example: 'worker-1'
    - **status**: Current status of the security scan (Type: string):
        - Example: 'success'
        - Enum: ['pending', 'in_progress', 'success', 'failed']
    - **updatedAt**: When the report was last updated (Type: string, date-time):
        - Example: '2025-01-17T12:30:00Z'
    - **createdAt**: When the report was created (Type: string, date-time):
        - Example: '2025-01-17T12:00:00Z'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the GetSecurityScan tool (Status: 400, Content-Type: application/json)
const GetSecurityScanResponseTemplate_B = `# API Response Information

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

// Response Template for the GetSecurityScan tool (Status: 401, Content-Type: application/json)
const GetSecurityScanResponseTemplate_C = `# API Response Information

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

// Response Template for the GetSecurityScan tool (Status: 403, Content-Type: application/json)
const GetSecurityScanResponseTemplate_D = `# API Response Information

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

// Response Template for the GetSecurityScan tool (Status: 404, Content-Type: application/json)
const GetSecurityScanResponseTemplate_E = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 404

**Content-Type:** application/json

> Not found - project, deployment, or security scan report not found

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the GetSecurityScan tool (Status: 500, Content-Type: application/json)
const GetSecurityScanResponseTemplate_F = `# API Response Information

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

// NewGetSecurityScanMCPTool creates the MCP Tool instance for GetSecurityScan
func NewGetSecurityScanMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetSecurityScan",
		"Get security scan report - Retrieves the security scan report for a deployment",
		[]byte(getSecurityScanInputSchema),
	)
}

// GetSecurityScanHandler is the handler function for the GetSecurityScan tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func GetSecurityScanHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "GetSecurityScan")
}
