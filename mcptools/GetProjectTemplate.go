package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the GetProjectTemplate tool
const getProjectTemplateInputSchema = `{
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

// Response Template for the GetProjectTemplate tool (Status: 200, Content-Type: application/json)
const GetProjectTemplateResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Project template details

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **Combines All Of the following structures**:
      - **Part 1**: Project template model (Type: object):
        - **createdAt** (Type: string, date-time):
        - **name** (Type: string):
        - **deploymentResources** (Type: object):
        - **useCases** (Type: string):
        - **imageUri** (Type: string):
        - **updatedAt** (Type: string, date-time):
        - **categories** (Type: array):
          - **Items** (Type: string):
        - **status** (Type: string):
            - Enum: ['draft', 'underReview', 'publishing', 'listed', 'rejected', 'deleted']
        - **projectConfiguration** (Type: object):
        - **id** (Type: string, uuid):
        - **projectId** (Type: string, uuid):
        - **amount** (Type: number):
        - **videoUrl** (Type: string):
        - **environmentConfiguration** (Type: object):
        - **environmentId** (Type: string, uuid):
        - **buildResources** (Type: object):
        - **overview** (Type: string):
        - **description** (Type: string):
        - **userId** (Type: string):
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

// Response Template for the GetProjectTemplate tool (Status: 400, Content-Type: application/json)
const GetProjectTemplateResponseTemplate_B = `# API Response Information

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

// Response Template for the GetProjectTemplate tool (Status: 403, Content-Type: application/json)
const GetProjectTemplateResponseTemplate_C = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 403

**Content-Type:** application/json

> Forbidden - template not listed and user is not owner

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the GetProjectTemplate tool (Status: 404, Content-Type: application/json)
const GetProjectTemplateResponseTemplate_D = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 404

**Content-Type:** application/json

> Not found - resource does not exist

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the GetProjectTemplate tool (Status: 500, Content-Type: application/json)
const GetProjectTemplateResponseTemplate_E = `# API Response Information

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

// NewGetProjectTemplateMCPTool creates the MCP Tool instance for GetProjectTemplate
func NewGetProjectTemplateMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"GetProjectTemplate",
		"Get project template by ID - Returns a single project template by ID. Listed templates are public; draft/under review/rejected are visible only to the owner. Optional auth; when authenticated, includes purchase info.",
		[]byte(getProjectTemplateInputSchema),
	)
}

// GetProjectTemplateHandler is the handler function for the GetProjectTemplate tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func GetProjectTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "GetProjectTemplate")
}
