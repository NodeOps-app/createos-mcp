package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListProjectTemplateCounts tool
const listProjectTemplateCountsInputSchema = `{
  "properties": {
    "name": {
      "description": "Filter counts by template name search",
      "type": "string"
    }
  },
  "type": "object"
}`

// Response Template for the ListProjectTemplateCounts tool (Status: 200, Content-Type: application/json)
const ListProjectTemplateCountsResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Total and category counts

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **total**: Total number of published templates (Type: integer):
    - **categoryCounts** (Type: array):
      - **Items** (Type: object):
        - **count** (Type: integer):
        - **name** (Type: string):
  - **status** (Type: string):
      - Example: 'success'
`

// Response Template for the ListProjectTemplateCounts tool (Status: 500, Content-Type: application/json)
const ListProjectTemplateCountsResponseTemplate_B = `# API Response Information

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

// NewListProjectTemplateCountsMCPTool creates the MCP Tool instance for ListProjectTemplateCounts
func NewListProjectTemplateCountsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListProjectTemplateCounts",
		"List project template counts by category - Returns total count of published templates and per-category counts. Optional name search.",
		[]byte(listProjectTemplateCountsInputSchema),
	)
}

// ListProjectTemplateCountsHandler is the handler function for the ListProjectTemplateCounts tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListProjectTemplateCountsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListProjectTemplateCounts")
}
