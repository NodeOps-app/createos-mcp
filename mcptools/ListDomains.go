package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the ListDomains tool
const listDomainsInputSchema = `{
  "properties": {
    "project_id": {
      "description": "Project unique identifier",
      "format": "uuid",
      "type": "string"
    }
  },
  "required": [
    "project_id"
  ],
  "type": "object"
}`

// Response Template for the ListDomains tool (Status: 200, Content-Type: application/json)
const ListDomainsResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 200

**Content-Type:** application/json

> Domains retrieved successfully

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'true'
  - **data** (Type: array):
    - **Items** (Type: object):
      - **lastFailedVerifiedAt**: Timestamp of last failed verification attempt (Type: string, date-time, nullable):
          - Nullable: true
      - **updatedAt**: When the domain was last updated (Type: string, date-time):
          - Example: '2025-01-17T12:00:00Z'
      - **projectId**: Project ID this domain belongs to (Type: string, uuid):
          - Example: '550e8400-e29b-41d4-a716-446655440001'
      - **deletingAt**: When the domain deletion was initiated (Type: string, date-time, nullable):
          - Nullable: true
      - **environmentId**: Desired environment ID for this domain (Type: string, uuid, nullable):
          - Nullable: true
          - Example: '550e8400-e29b-41d4-a716-446655440002'
      - **name**: Domain name (Type: string):
          - Example: 'example.com'
      - **records**: DNS records required for domain verification (Type: object):
        - **Allows Additional Properties**
      - **status**: Current status of the domain (Type: string):
          - Example: 'active'
          - Enum: ['pending', 'issuing', 'updating', 'active', 'error', 'deleting']
      - **createdAt**: When the domain was created (Type: string, date-time):
          - Example: '2025-01-17T12:00:00Z'
      - **domainCertificateId**: Associated SSL certificate ID (Type: string, uuid, nullable):
          - Nullable: true
          - Example: '550e8400-e29b-41d4-a716-446655440003'
      - **failedVerificationCount**: Number of failed DNS verification attempts (Type: integer):
          - Example: '0'
      - **id**: Domain unique identifier (Type: string, uuid):
          - Example: '550e8400-e29b-41d4-a716-446655440000'
      - **message**: Error or status message (Type: string, nullable):
          - Nullable: true
`

// Response Template for the ListDomains tool (Status: 400, Content-Type: application/json)
const ListDomainsResponseTemplate_B = `# API Response Information

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

// Response Template for the ListDomains tool (Status: 401, Content-Type: application/json)
const ListDomainsResponseTemplate_C = `# API Response Information

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

// Response Template for the ListDomains tool (Status: 403, Content-Type: application/json)
const ListDomainsResponseTemplate_D = `# API Response Information

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

// Response Template for the ListDomains tool (Status: 404, Content-Type: application/json)
const ListDomainsResponseTemplate_E = `# API Response Information

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

// Response Template for the ListDomains tool (Status: 422, Content-Type: application/json)
const ListDomainsResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 422

**Content-Type:** application/json

> Unprocessable Entity - project is not active

## Response Structure

- Structure (Type: object):
  - **success** (Type: boolean):
      - Example: 'false'
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
`

// Response Template for the ListDomains tool (Status: 500, Content-Type: application/json)
const ListDomainsResponseTemplate_G = `# API Response Information

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

// NewListDomainsMCPTool creates the MCP Tool instance for ListDomains
func NewListDomainsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"ListDomains",
		"List domains - Returns all custom domains for a project",
		[]byte(listDomainsInputSchema),
	)
}

// ListDomainsHandler is the handler function for the ListDomains tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func ListDomainsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "ListDomains")
}
