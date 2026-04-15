package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListPublishedProjectTemplates tool
const listPublishedProjectTemplatesInputSchema = `{
  "properties": {
    "categories": {
      "description": "Comma-separated category names to filter by",
      "type": "string"
    },
    "limit": {
      "default": 10,
      "description": "Maximum number of templates to return (default 10, max 20)",
      "maximum": 20,
      "minimum": 1,
      "type": "integer"
    },
    "name": {
      "description": "Search by template name",
      "type": "string"
    },
    "offset": {
      "default": 0,
      "description": "Number of templates to skip (default 0)",
      "minimum": 0,
      "type": "integer"
    }
  },
  "type": "object"
}`

// Response Template for the ListPublishedProjectTemplates tool (Status: 200, Content-Type: application/json)
const ListPublishedProjectTemplatesResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> List of published project templates

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **templates** (Type: array):
      - **Items** (Type: object):
        - **Combines All Of the following structures**:
          - **Part 1**: Project template model (Type: object):
            - **environmentConfiguration** (Type: object):
            - **createdAt** (Type: string, date-time):
            - **projectConfiguration** (Type: object):
            - **imageUri** (Type: string):
            - **categories** (Type: array):
              - **Items** (Type: string):
            - **id** (Type: string, uuid):
            - **buildResources** (Type: object):
            - **name** (Type: string):
            - **userId** (Type: string):
            - **useCases** (Type: string):
            - **description** (Type: string):
            - **projectId** (Type: string, uuid):
            - **updatedAt** (Type: string, date-time):
            - **environmentId** (Type: string, uuid):
            - **status** (Type: string):
                - Enum: ['draft', 'underReview', 'publishing', 'listed', 'rejected', 'deleted']
            - **overview** (Type: string):
            - **videoUrl** (Type: string):
            - **amount** (Type: number):
            - **deploymentResources** (Type: object):
          - **Part 2** (Type: object):
            - **hasUserPurchased** (Type: boolean):
            - **templatePurchasedId** (Type: string, uuid, nullable):
                - Nullable: true
            - **author** (Type: object):
              - **id**: User ID of the author (Type: string):
              - **name**: Display name or email local part (Type: string, nullable):
                  - Nullable: true
    - **pagination** (Type: object):
      - **total**: Total number of items available (Type: integer):
          - Example: '100'
      - **count**: Number of items in the current page (Type: integer):
          - Example: '10'
      - **limit**: Maximum number of items per page (Type: integer):
          - Example: '10'
      - **offset**: Number of items skipped (Type: integer):
          - Example: '0'
  - **status** (Type: string):
      - Example: 'success'
`

// Response Template for the ListPublishedProjectTemplates tool (Status: 500, Content-Type: application/json)
const ListPublishedProjectTemplatesResponseTemplate_B = `# API Response Information

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

// NewListPublishedProjectTemplatesMCPTool creates the MCP Tool instance for ListPublishedProjectTemplates
func NewListPublishedProjectTemplatesMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListPublishedProjectTemplates",
		"List published project templates - Returns a paginated list of published (listed) project templates. Optional auth; when authenticated, includes whether the user has purchased each template. Can filter by categories (comma-separated) and search by name.",
		[]byte(listPublishedProjectTemplatesInputSchema),
	)
}

// ListPublishedProjectTemplatesHandler is the handler function for the ListPublishedProjectTemplates tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListPublishedProjectTemplatesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListPublishedProjectTemplates")
}
