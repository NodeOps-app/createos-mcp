package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListServicesByApp tool
const listServicesByAppInputSchema = `{
  "properties": {
    "app_id": {
      "description": "App unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "limit": {
      "default": 10,
      "description": "Maximum number of services to return (default: 10, max: 20)",
      "maximum": 20,
      "minimum": 1,
      "type": "integer"
    },
    "offset": {
      "default": 0,
      "description": "Number of services to skip (default: 0)",
      "minimum": 0,
      "type": "integer"
    }
  },
  "required": [
    "app_id"
  ],
  "type": "object"
}`

// Response Template for the ListServicesByApp tool (Status: 200, Content-Type: application/json)
const ListServicesByAppResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Services retrieved successfully

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'true'
  - **data** (Type: object):
    - **data** (Type: array):
      - **Items**: Service instance with optional plan information (Type: object):
    - **pagination** (Type: object):
      - **offset**: Number of items skipped (Type: integer):
          - Example: '0'
      - **total**: Total number of items available (Type: integer):
          - Example: '100'
      - **count**: Number of items in the current page (Type: integer):
          - Example: '10'
      - **limit**: Maximum number of items per page (Type: integer):
          - Example: '10'
`

// Response Template for the ListServicesByApp tool (Status: 400, Content-Type: application/json)
const ListServicesByAppResponseTemplate_B = `# API Response Information

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

// Response Template for the ListServicesByApp tool (Status: 401, Content-Type: application/json)
const ListServicesByAppResponseTemplate_C = `# API Response Information

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

// Response Template for the ListServicesByApp tool (Status: 403, Content-Type: application/json)
const ListServicesByAppResponseTemplate_D = `# API Response Information

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

// Response Template for the ListServicesByApp tool (Status: 404, Content-Type: application/json)
const ListServicesByAppResponseTemplate_E = `# API Response Information

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

// Response Template for the ListServicesByApp tool (Status: 500, Content-Type: application/json)
const ListServicesByAppResponseTemplate_F = `# API Response Information

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

// NewListServicesByAppMCPTool creates the MCP Tool instance for ListServicesByApp
func NewListServicesByAppMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListServicesByApp",
		"List services by app - Returns all services associated with an app",
		[]byte(listServicesByAppInputSchema),
	)
}

// ListServicesByAppHandler is the handler function for the ListServicesByApp tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListServicesByAppHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListServicesByApp")
}
