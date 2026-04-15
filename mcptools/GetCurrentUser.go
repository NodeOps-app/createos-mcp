package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the GetCurrentUser tool
const getCurrentUserInputSchema = `{
  "type": "object"
}`

// Response Template for the GetCurrentUser tool (Status: 200, Content-Type: application/json)
const GetCurrentUserResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> User information retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data**: User information from OneClick authentication service (Type: object):
    - **Allows Additional Properties**
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the GetCurrentUser tool (Status: 401, Content-Type: application/json)
const GetCurrentUserResponseTemplate_B = `# API Response Information

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

// Response Template for the GetCurrentUser tool (Status: 403, Content-Type: application/json)
const GetCurrentUserResponseTemplate_C = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 403

**Content-Type:** application/json

> Forbidden - user is suspended

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the GetCurrentUser tool (Status: 500, Content-Type: application/json)
const GetCurrentUserResponseTemplate_D = `# API Response Information

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

// NewGetCurrentUserMCPTool creates the MCP Tool instance for GetCurrentUser
func NewGetCurrentUserMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetCurrentUser",
		"Get current user - Returns information about the currently authenticated user",
		[]byte(getCurrentUserInputSchema),
	)
}

// GetCurrentUserHandler is the handler function for the GetCurrentUser tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func GetCurrentUserHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "GetCurrentUser")
}
