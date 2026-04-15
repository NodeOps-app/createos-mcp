package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListGithubRepositoryBranches tool
const listGithubRepositoryBranchesInputSchema = `{
  "properties": {
    "installation_id": {
      "description": "GitHub App installation ID",
      "type": "string"
    },
    "page": {
      "default": 1,
      "description": "Page number for pagination",
      "minimum": 1,
      "type": "integer"
    },
    "per-page": {
      "default": 30,
      "description": "Number of items per page",
      "maximum": 100,
      "minimum": 1,
      "type": "integer"
    },
    "protected": {
      "description": "Filter by protected status",
      "enum": [
        "true",
        "false"
      ],
      "type": "string"
    },
    "repository": {
      "description": "Repository full name (owner/repo)",
      "example": "owner/repo",
      "type": "string"
    }
  },
  "required": [
    "installation_id",
    "repository"
  ],
  "type": "object"
}`

// Response Template for the ListGithubRepositoryBranches tool (Status: 200, Content-Type: application/json)
const ListGithubRepositoryBranchesResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Branches retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: array):
    - **Items** (Type: object):
      - **name** (Type: string):
      - **protected** (Type: boolean):
      - **commit** (Type: object):
        - **sha** (Type: string):
        - **url** (Type: string):
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the ListGithubRepositoryBranches tool (Status: 400, Content-Type: application/json)
const ListGithubRepositoryBranchesResponseTemplate_B = `# API Response Information

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

// Response Template for the ListGithubRepositoryBranches tool (Status: 401, Content-Type: application/json)
const ListGithubRepositoryBranchesResponseTemplate_C = `# API Response Information

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

// Response Template for the ListGithubRepositoryBranches tool (Status: 404, Content-Type: application/json)
const ListGithubRepositoryBranchesResponseTemplate_D = `# API Response Information

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

// Response Template for the ListGithubRepositoryBranches tool (Status: 500, Content-Type: application/json)
const ListGithubRepositoryBranchesResponseTemplate_E = `# API Response Information

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

// NewListGithubRepositoryBranchesMCPTool creates the MCP Tool instance for ListGithubRepositoryBranches
func NewListGithubRepositoryBranchesMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListGithubRepositoryBranches",
		"List repository branches - Lists all branches for a GitHub repository",
		[]byte(listGithubRepositoryBranchesInputSchema),
	)
}

// ListGithubRepositoryBranchesHandler is the handler function for the ListGithubRepositoryBranches tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListGithubRepositoryBranchesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListGithubRepositoryBranches")
}
