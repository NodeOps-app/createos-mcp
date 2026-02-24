package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the BuyProjectTemplate tool
const buyProjectTemplateInputSchema = `{
  "properties": {
    "template_id": {
      "format": "uuid",
      "type": "string"
    }
  },
  "required": [
    "template_id"
  ],
  "type": "object"
}`

// Response Template for the BuyProjectTemplate tool (Status: 200, Content-Type: application/json)
const BuyProjectTemplateResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Purchase created; returns purchase ID

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **id**: Created resource ID (e.g. purchase ID) (Type: string, uuid):
  - **status** (Type: string):
      - Example: 'success'
`

// Response Template for the BuyProjectTemplate tool (Status: 400, Content-Type: application/json)
const BuyProjectTemplateResponseTemplate_B = `# API Response Information

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

// Response Template for the BuyProjectTemplate tool (Status: 401, Content-Type: application/json)
const BuyProjectTemplateResponseTemplate_C = `# API Response Information

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

// Response Template for the BuyProjectTemplate tool (Status: 404, Content-Type: application/json)
const BuyProjectTemplateResponseTemplate_D = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 404

**Content-Type:** application/json

> Template not found or deleted

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the BuyProjectTemplate tool (Status: 422, Content-Type: application/json)
const BuyProjectTemplateResponseTemplate_E = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 422

**Content-Type:** application/json

> Already purchased or insufficient credit

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the BuyProjectTemplate tool (Status: 500, Content-Type: application/json)
const BuyProjectTemplateResponseTemplate_F = `# API Response Information

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

// NewBuyProjectTemplateMCPTool creates the MCP Tool instance for BuyProjectTemplate
func NewBuyProjectTemplateMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"BuyProjectTemplate",
		"Buy project template - Purchases the template using the user's credits. Deducts template amount from user balance and creates a purchase record. User must not have already purchased this template.",
		[]byte(buyProjectTemplateInputSchema),
	)
}

// BuyProjectTemplateHandler is the handler function for the BuyProjectTemplate tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func BuyProjectTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "BuyProjectTemplate")
}
