package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListConnectedGithubAccounts tool
const listConnectedGithubAccountsInputSchema = `{
  "type": "object"
}`

// Response Template for the ListConnectedGithubAccounts tool (Status: 200, Content-Type: application/json)
const ListConnectedGithubAccountsResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Connected accounts retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: array):
    - **Items** (Type: object):
      - **avatarUrl** (Type: string):
      - **installationId** (Type: integer, int64):
      - **ownerId** (Type: integer, int64):
      - **username** (Type: string):
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the ListConnectedGithubAccounts tool (Status: 401, Content-Type: application/json)
const ListConnectedGithubAccountsResponseTemplate_B = `# API Response Information

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

// Response Template for the ListConnectedGithubAccounts tool (Status: 500, Content-Type: application/json)
const ListConnectedGithubAccountsResponseTemplate_C = `# API Response Information

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

// NewListConnectedGithubAccountsMCPTool creates the MCP Tool instance for ListConnectedGithubAccounts
func NewListConnectedGithubAccountsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListConnectedGithubAccounts",
		"List connected GitHub accounts - Returns a list of all GitHub accounts connected to the user",
		[]byte(listConnectedGithubAccountsInputSchema),
	)
}

// ListConnectedGithubAccountsHandler is the handler function for the ListConnectedGithubAccounts tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListConnectedGithubAccountsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListConnectedGithubAccounts")
}
