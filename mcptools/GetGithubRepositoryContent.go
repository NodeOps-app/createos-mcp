package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the GetGithubRepositoryContent tool
const getGithubRepositoryContentInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "branch": {
          "description": "Branch name",
          "example": "main",
          "type": "string"
        },
        "repository": {
          "description": "Repository full name (owner/repo)",
          "example": "owner/repo",
          "type": "string"
        },
        "treeSha": {
          "description": "SHA of the tree to get (optional, defaults to branch)",
          "example": "abc123def456",
          "type": "string"
        }
      },
      "required": [
        "repository",
        "branch"
      ],
      "type": "object"
    },
    "installation_id": {
      "description": "GitHub App installation ID",
      "type": "string"
    }
  },
  "required": [
    "installation_id",
    "body"
  ],
  "type": "object"
}`

// Response Template for the GetGithubRepositoryContent tool (Status: 200, Content-Type: application/json)
const GetGithubRepositoryContentResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Repository content retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data**: GitHub repository tree structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the GetGithubRepositoryContent tool (Status: 400, Content-Type: application/json)
const GetGithubRepositoryContentResponseTemplate_B = `# API Response Information

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

// Response Template for the GetGithubRepositoryContent tool (Status: 401, Content-Type: application/json)
const GetGithubRepositoryContentResponseTemplate_C = `# API Response Information

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

// Response Template for the GetGithubRepositoryContent tool (Status: 404, Content-Type: application/json)
const GetGithubRepositoryContentResponseTemplate_D = `# API Response Information

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

// Response Template for the GetGithubRepositoryContent tool (Status: 500, Content-Type: application/json)
const GetGithubRepositoryContentResponseTemplate_E = `# API Response Information

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

// NewGetGithubRepositoryContentMCPTool creates the MCP Tool instance for GetGithubRepositoryContent
func NewGetGithubRepositoryContentMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetGithubRepositoryContent",
		"Get repository content - Retrieves the file tree structure of a repository",
		[]byte(getGithubRepositoryContentInputSchema),
	)
}

// GetGithubRepositoryContentHandler is the handler function for the GetGithubRepositoryContent tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func GetGithubRepositoryContentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "GetGithubRepositoryContent")
}
