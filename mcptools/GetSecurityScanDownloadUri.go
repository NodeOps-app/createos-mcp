package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the GetSecurityScanDownloadUri tool
const getSecurityScanDownloadUriInputSchema = `{
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

// Response Template for the GetSecurityScanDownloadUri tool (Status: 200, Content-Type: application/json)
const GetSecurityScanDownloadUriResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Download URL retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **downloadUri**: Signed URL to download the security scan report (valid for 1 hour) (Type: string):
        - Example: 'https://storage.example.com/reports/scan.json?signature=...'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the GetSecurityScanDownloadUri tool (Status: 400, Content-Type: application/json)
const GetSecurityScanDownloadUriResponseTemplate_B = `# API Response Information

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

// Response Template for the GetSecurityScanDownloadUri tool (Status: 401, Content-Type: application/json)
const GetSecurityScanDownloadUriResponseTemplate_C = `# API Response Information

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

// Response Template for the GetSecurityScanDownloadUri tool (Status: 403, Content-Type: application/json)
const GetSecurityScanDownloadUriResponseTemplate_D = `# API Response Information

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

// Response Template for the GetSecurityScanDownloadUri tool (Status: 404, Content-Type: application/json)
const GetSecurityScanDownloadUriResponseTemplate_E = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 404

**Content-Type:** application/json

> Not found - project, deployment, security scan report not found, or report not successful

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the GetSecurityScanDownloadUri tool (Status: 500, Content-Type: application/json)
const GetSecurityScanDownloadUriResponseTemplate_F = `# API Response Information

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

// NewGetSecurityScanDownloadUriMCPTool creates the MCP Tool instance for GetSecurityScanDownloadUri
func NewGetSecurityScanDownloadUriMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetSecurityScanDownloadUri",
		"Download security scan report - Gets a signed URL to download the full security scan report. Only available when the scan status is successful.",
		[]byte(getSecurityScanDownloadUriInputSchema),
	)
}

// GetSecurityScanDownloadUriHandler is the handler function for the GetSecurityScanDownloadUri tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func GetSecurityScanDownloadUriHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "GetSecurityScanDownloadUri")
}
