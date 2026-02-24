package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the GetSupportedProjectTypes tool
const getSupportedProjectTypesInputSchema = `{
  "type": "object"
}`

// Response Template for the GetSupportedProjectTypes tool (Status: 200, Content-Type: application/json)
const GetSupportedProjectTypesResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Supported project types retrieved successfully

## Response Structure

- Structure (Type: array):
  - **Items** (Type: object):
    - **editables**: Editable configuration options (Type: object):
    - **name**: Name of the framework or runtime (Type: string):
    - **runtimes**: List of compatible runtimes (for frameworks) (Type: array):
      - **Items** (Type: string):
    - **type** (Type: string):
        - Enum: ['framework', 'runtime']
`

// Response Template for the GetSupportedProjectTypes tool (Status: 401, Content-Type: application/json)
const GetSupportedProjectTypesResponseTemplate_B = `# API Response Information

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

// Response Template for the GetSupportedProjectTypes tool (Status: 500, Content-Type: application/json)
const GetSupportedProjectTypesResponseTemplate_C = `# API Response Information

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

// NewGetSupportedProjectTypesMCPTool creates the MCP Tool instance for GetSupportedProjectTypes
func NewGetSupportedProjectTypesMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetSupportedProjectTypes",
		"Get supported project types - Returns the list of supported frameworks and runtimes",
		[]byte(getSupportedProjectTypesInputSchema),
	)
}

// GetSupportedProjectTypesHandler is the handler function for the GetSupportedProjectTypes tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func GetSupportedProjectTypesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "GetSupportedProjectTypes")
}
