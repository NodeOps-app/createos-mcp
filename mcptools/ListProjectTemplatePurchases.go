package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListProjectTemplatePurchases tool
const listProjectTemplatePurchasesInputSchema = `{
  "properties": {
    "limit": {
      "default": 10,
      "maximum": 50,
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

// Response Template for the ListProjectTemplatePurchases tool (Status: 200, Content-Type: application/json)
const ListProjectTemplatePurchasesResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> List of purchases with template info

## Response Structure

- Structure (Type: object):
  - **status** (Type: string):
      - Example: 'success'
  - **data** (Type: object):
    - **pagination** (Type: object):
      - **total**: Total number of items available (Type: integer):
          - Example: '100'
      - **count**: Number of items in the current page (Type: integer):
          - Example: '10'
      - **limit**: Maximum number of items per page (Type: integer):
          - Example: '10'
      - **offset**: Number of items skipped (Type: integer):
          - Example: '0'
    - **purchases** (Type: array):
      - **Items** (Type: object):
        - **projectTemplateId** (Type: string, uuid):
        - **purchasedAmount** (Type: number):
        - **template**: Project template model (Type: object):
          - **useCases** (Type: string):
          - **createdAt** (Type: string, date-time):
          - **name** (Type: string):
          - **deploymentResources** (Type: object):
          - **id** (Type: string, uuid):
          - **imageUri** (Type: string):
          - **updatedAt** (Type: string, date-time):
          - **categories** (Type: array):
            - **Items** (Type: string):
          - **status** (Type: string):
              - Enum: ['draft', 'underReview', 'publishing', 'listed', 'rejected', 'deleted']
          - **projectConfiguration** (Type: object):
          - **environmentId** (Type: string, uuid):
          - **projectId** (Type: string, uuid):
          - **amount** (Type: number):
          - **videoUrl** (Type: string):
          - **environmentConfiguration** (Type: object):
          - **userId** (Type: string):
          - **buildResources** (Type: object):
          - **overview** (Type: string):
          - **description** (Type: string):
        - **updatedAt** (Type: string, date-time):
        - **userId** (Type: string):
        - **createdAt** (Type: string, date-time):
        - **id** (Type: string, uuid):
`

// Response Template for the ListProjectTemplatePurchases tool (Status: 401, Content-Type: application/json)
const ListProjectTemplatePurchasesResponseTemplate_B = `# API Response Information

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

// Response Template for the ListProjectTemplatePurchases tool (Status: 500, Content-Type: application/json)
const ListProjectTemplatePurchasesResponseTemplate_C = `# API Response Information

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

// NewListProjectTemplatePurchasesMCPTool creates the MCP Tool instance for ListProjectTemplatePurchases
func NewListProjectTemplatePurchasesMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListProjectTemplatePurchases",
		"List project template purchases - Returns the authenticated user's project template purchases (paginated).",
		[]byte(listProjectTemplatePurchasesInputSchema),
	)
}

// ListProjectTemplatePurchasesHandler is the handler function for the ListProjectTemplatePurchases tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListProjectTemplatePurchasesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListProjectTemplatePurchases")
}
