package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListMyProjectTemplates tool
const listMyProjectTemplatesInputSchema = `{
  "properties": {
    "limit": {
      "default": 10,
      "maximum": 20,
      "minimum": 1,
      "type": "integer"
    },
    "offset": {
      "default": 0,
      "minimum": 0,
      "type": "integer"
    }
  },
  "type": "object"
}`

// Response Template for the ListMyProjectTemplates tool (Status: 200, Content-Type: application/json)
const ListMyProjectTemplatesResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> List of user's project templates

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **pagination** (Type: object):
      - **offset**: Number of items skipped (Type: integer):
          - Example: '0'
      - **total**: Total number of items available (Type: integer):
          - Example: '100'
      - **count**: Number of items in the current page (Type: integer):
          - Example: '10'
      - **limit**: Maximum number of items per page (Type: integer):
          - Example: '10'
    - **templates** (Type: array):
      - **Items** (Type: object):
        - **Combines All Of the following structures**:
          - **Part 1**: Project template model (Type: object):
            - **overview** (Type: string):
            - **description** (Type: string):
            - **userId** (Type: string):
            - **buildResources** (Type: object):
            - **name** (Type: string):
            - **deploymentResources** (Type: object):
            - **useCases** (Type: string):
            - **createdAt** (Type: string, date-time):
            - **status** (Type: string):
                - Enum: ['draft', 'underReview', 'publishing', 'listed', 'rejected', 'deleted']
            - **projectConfiguration** (Type: object):
            - **id** (Type: string, uuid):
            - **imageUri** (Type: string):
            - **updatedAt** (Type: string, date-time):
            - **categories** (Type: array):
              - **Items** (Type: string):
            - **videoUrl** (Type: string):
            - **environmentConfiguration** (Type: object):
            - **environmentId** (Type: string, uuid):
            - **projectId** (Type: string, uuid):
            - **amount** (Type: number):
          - **Part 2** (Type: object):
            - **hasUserPurchased** (Type: boolean):
            - **templatePurchasedId** (Type: string, uuid, nullable):
                - Nullable: true
            - **author** (Type: object):
              - **id**: User ID of the author (Type: string):
              - **name**: Display name or email local part (Type: string, nullable):
                  - Nullable: true
  - **status** (Type: string):
      - Example: 'success'
`

// Response Template for the ListMyProjectTemplates tool (Status: 401, Content-Type: application/json)
const ListMyProjectTemplatesResponseTemplate_B = `# API Response Information

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

// Response Template for the ListMyProjectTemplates tool (Status: 500, Content-Type: application/json)
const ListMyProjectTemplatesResponseTemplate_C = `# API Response Information

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

// NewListMyProjectTemplatesMCPTool creates the MCP Tool instance for ListMyProjectTemplates
func NewListMyProjectTemplatesMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListMyProjectTemplates",
		"List my project templates - Returns the authenticated user's project templates (all statuses), paginated.",
		[]byte(listMyProjectTemplatesInputSchema),
	)
}

// ListMyProjectTemplatesHandler is the handler function for the ListMyProjectTemplates tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListMyProjectTemplatesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListMyProjectTemplates")
}
