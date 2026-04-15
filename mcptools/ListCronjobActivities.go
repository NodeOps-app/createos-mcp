package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListCronjobActivities tool
const listCronjobActivitiesInputSchema = `{
  "properties": {
    "cronjob_id": {
      "description": "Cronjob unique identifier",
      "format": "uuid",
      "type": "string"
    },
    "project_id": {
      "description": "Project unique identifier",
      "format": "uuid",
      "type": "string"
    }
  },
  "required": [
    "project_id",
    "cronjob_id"
  ],
  "type": "object"
}`

// Response Template for the ListCronjobActivities tool (Status: 200, Content-Type: application/json)
const ListCronjobActivitiesResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Cronjob activities retrieved successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: array):
    - **Items** (Type: object):
      - **success** (Type: boolean, nullable):
          - Nullable: true
      - **createdAt** (Type: string, date-time):
      - **cronJobId** (Type: string, uuid):
      - **id** (Type: string, uuid):
      - **log** (Type: string, nullable):
          - Nullable: true
      - **scheduledAt** (Type: string, date-time):
      - **statusCode** (Type: integer, nullable):
          - Nullable: true
  - **success** (Type: boolean):
      - Example: 'true'
`

// Response Template for the ListCronjobActivities tool (Status: 401, Content-Type: application/json)
const ListCronjobActivitiesResponseTemplate_B = `# API Response Information

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

// Response Template for the ListCronjobActivities tool (Status: 403, Content-Type: application/json)
const ListCronjobActivitiesResponseTemplate_C = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 403

**Content-Type:** application/json

> Forbidden - user does not have permission

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the ListCronjobActivities tool (Status: 404, Content-Type: application/json)
const ListCronjobActivitiesResponseTemplate_D = `# API Response Information

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

// Response Template for the ListCronjobActivities tool (Status: 500, Content-Type: application/json)
const ListCronjobActivitiesResponseTemplate_E = `# API Response Information

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

// NewListCronjobActivitiesMCPTool creates the MCP Tool instance for ListCronjobActivities
func NewListCronjobActivitiesMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListCronjobActivities",
		"List cronjob activities - Retrieves the last 10 execution activities for a cronjob, including status codes, success/failure, and logs.",
		[]byte(listCronjobActivitiesInputSchema),
	)
}

// ListCronjobActivitiesHandler is the handler function for the ListCronjobActivities tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListCronjobActivitiesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListCronjobActivities")
}
