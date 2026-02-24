package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the TransferProject tool
const transferProjectInputSchema = `{
  "properties": {
    "project_id": {
      "description": "Project unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "token": {
      "description": "Transfer token obtained from the transfer-uri endpoint",
      "type": "string"
    }
  },
  "required": [
    "project_id",
    "token"
  ],
  "type": "object"
}`

// Response Template for the TransferProject tool (Status: 200, Content-Type: application/json)
const TransferProjectResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Project transferred successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: string):
      - Example: 'project transferred successfully'
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the TransferProject tool (Status: 400, Content-Type: application/json)
const TransferProjectResponseTemplate_B = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 400

**Content-Type:** application/json

> Bad request - invalid project ID or missing token

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the TransferProject tool (Status: 401, Content-Type: application/json)
const TransferProjectResponseTemplate_C = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 401

**Content-Type:** application/json

> Unauthorized - invalid or expired token

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the TransferProject tool (Status: 403, Content-Type: application/json)
const TransferProjectResponseTemplate_D = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 403

**Content-Type:** application/json

> Forbidden - token does not match project or owner

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the TransferProject tool (Status: 404, Content-Type: application/json)
const TransferProjectResponseTemplate_E = `# API Response Information

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

// Response Template for the TransferProject tool (Status: 429, Content-Type: application/json)
const TransferProjectResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 429

**Content-Type:** application/json

> Too many requests - maximum projects per user reached

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the TransferProject tool (Status: 500, Content-Type: application/json)
const TransferProjectResponseTemplate_G = `# API Response Information

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

// NewTransferProjectMCPTool creates the MCP Tool instance for TransferProject
func NewTransferProjectMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"TransferProject",
		"Transfer project ownership - Transfers project ownership to the authenticated user using a valid transfer token. The token must be obtained from the transfer-uri endpoint by the current project owner.",
		[]byte(transferProjectInputSchema),
	)
}

// TransferProjectHandler is the handler function for the TransferProject tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func TransferProjectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "TransferProject")
}
