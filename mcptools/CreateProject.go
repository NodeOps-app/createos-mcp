package mcptools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

// Input Schema for the CreateProject tool
const createProjectInputSchema = `{
  "properties": {
    "body": {
      "properties": {
        "appId": {
          "description": "Optional app ID to associate the project with",
          "example": "550e8400-e29b-41d4-a716-446655440000",
          "format": "uuid",
          "type": [
            "string",
            "null"
          ]
        },
        "description": {
          "description": "Detailed description of the project",
          "example": "A full-stack application built with Next js and TypeScript",
          "maxLength": 2048,
          "minLength": 4,
          "type": [
            "string",
            "null"
          ]
        },
        "displayName": {
          "description": "Human-readable display name for the project",
          "example": "My Awesome Application",
          "maxLength": 48,
          "minLength": 4,
          "pattern": "^[a-zA-Z0-9 _-]+$",
          "type": "string"
        },
        "enabledSecurityScan": {
          "default": false,
          "description": "Enable security scanning for the project",
          "example": false,
          "type": "boolean"
        },
        "settings": {
          "description": "Build and runtime settings for the project. For VCS projects: use VCSSettings schema. For image projects: use ImageSettings schema. For upload projects: use UploadSettings schema.",
          "oneOf": [
            {
              "description": "Build and runtime settings for VCS projects",
              "properties": {
                "buildCommand": {
                  "description": "Command to build the project",
                  "example": "npm run build",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "buildDir": {
                  "description": "Directory where build output is located",
                  "example": "dist",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "buildFlag": {
                  "description": "Additional flags to pass to the build command",
                  "example": "-trimpath",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "buildFlags": {
                  "additionalProperties": {
                    "type": "string"
                  },
                  "description": "Build-time environment variables (also accepts buildVars)",
                  "example": {
                    "NEXT_PUBLIC_API_URL": "https://api.example.com",
                    "NODE_ENV": "production"
                  },
                  "type": "object"
                },
                "buildVars": {
                  "additionalProperties": {
                    "type": "string"
                  },
                  "description": "Build-time environment variables (alternative to buildFlags)",
                  "example": {
                    "NEXT_PUBLIC_API_URL": "https://api.example.com",
                    "NODE_ENV": "production"
                  },
                  "type": "object"
                },
                "directoryPath": {
                  "description": "Project directory path within the repository (defaults to root)",
                  "example": ".",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "framework": {
                  "description": "Framework name (e.g., nextjs, reactjs-spa, reactjs-ssr). If specified, runtime must be compatible.",
                  "example": "nextjs",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "hasDockerfile": {
                  "default": false,
                  "description": "Whether the project has a Dockerfile (if true, build settings may be ignored)",
                  "example": false,
                  "type": "boolean"
                },
                "ignoreBranches": {
                  "description": "List of exact branch names to ignore for automatic deployments (not regex patterns)",
                  "example": [
                    "develop",
                    "staging"
                  ],
                  "items": {
                    "maxLength": 255,
                    "type": "string"
                  },
                  "type": "array"
                },
                "installCommand": {
                  "description": "Command to install dependencies",
                  "example": "npm install",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "port": {
                  "description": "Port number the application listens on",
                  "example": 3000,
                  "maximum": 65535,
                  "minimum": 1,
                  "type": "integer"
                },
                "runCommand": {
                  "description": "Command to run the application",
                  "example": "npm run start",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "runEnvs": {
                  "additionalProperties": {
                    "type": "string"
                  },
                  "description": "Runtime environment variables",
                  "example": {
                    "API_KEY": "secret-key",
                    "DATABASE_URL": "postgresql://localhost:5432/mydb"
                  },
                  "type": "object"
                },
                "runFlag": {
                  "description": "Additional flags to pass to the run command",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "runtime": {
                  "description": "Runtime environment (e.g., node:20, golang:1.25, bun:1.3). Required if framework is not specified.",
                  "example": "node:20",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "useBuildAI": {
                  "default": false,
                  "description": "Whether to use Build AI for automated build configuration",
                  "example": false,
                  "type": "boolean"
                }
              },
              "type": "object"
            },
            {
              "additionalProperties": false,
              "description": "Settings for image type projects",
              "not": {
                "properties": {
                  "buildCommand": {
                    "type": "string"
                  },
                  "directoryPath": {
                    "type": "string"
                  },
                  "framework": {
                    "type": "string"
                  },
                  "hasDockerfile": {
                    "type": "boolean"
                  },
                  "ignoreBranches": {
                    "items": {
                      "type": "string"
                    },
                    "type": "array"
                  },
                  "installCommand": {
                    "type": "string"
                  },
                  "runtime": {
                    "type": "string"
                  },
                  "useBuildAI": {
                    "type": "boolean"
                  }
                },
                "type": "object"
              },
              "properties": {
                "port": {
                  "description": "Port number the application listens on",
                  "example": 8080,
                  "maximum": 65535,
                  "minimum": 1,
                  "type": "integer"
                },
                "runEnvs": {
                  "additionalProperties": {
                    "type": "string"
                  },
                  "description": "Runtime environment variables",
                  "example": {
                    "API_KEY": "secret-key",
                    "DATABASE_URL": "postgresql://localhost:5432/mydb"
                  },
                  "type": "object"
                }
              },
              "required": [
                "port"
              ],
              "type": "object"
            },
            {
              "additionalProperties": false,
              "allOf": [
                {
                  "not": {
                    "additionalProperties": false,
                    "properties": {
                      "port": {
                        "maximum": 65535,
                        "minimum": 1,
                        "type": "integer"
                      },
                      "runEnvs": {
                        "additionalProperties": {
                          "type": "string"
                        },
                        "type": "object"
                      }
                    },
                    "required": [
                      "port"
                    ],
                    "type": "object"
                  }
                },
                {
                  "not": {
                    "anyOf": [
                      {
                        "properties": {
                          "ignoreBranches": {
                            "items": {
                              "type": "string"
                            },
                            "type": "array"
                          }
                        },
                        "type": "object"
                      },
                      {
                        "properties": {
                          "buildFlags": {
                            "additionalProperties": {
                              "type": "string"
                            },
                            "type": "object"
                          }
                        },
                        "type": "object"
                      }
                    ]
                  }
                }
              ],
              "description": "Build and runtime settings for upload type projects",
              "properties": {
                "buildCommand": {
                  "description": "Command to build the project",
                  "example": "npm run build",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "buildDir": {
                  "description": "Directory where build output is located",
                  "example": "dist",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "buildFlag": {
                  "description": "Additional flags to pass to the build command",
                  "example": "-trimpath",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "buildVars": {
                  "additionalProperties": {
                    "type": "string"
                  },
                  "description": "Build-time environment variables",
                  "example": {
                    "NEXT_PUBLIC_API_URL": "https://api.example.com",
                    "NODE_ENV": "production"
                  },
                  "type": "object"
                },
                "directoryPath": {
                  "description": "Project directory path within the uploaded files (defaults to root)",
                  "example": ".",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "framework": {
                  "description": "Framework name (e.g., nextjs, reactjs-spa, reactjs-ssr). Required if useBuildAI is false and hasDockerfile is false.",
                  "example": "nextjs",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "hasDockerfile": {
                  "default": false,
                  "description": "Whether the project has a Dockerfile (if true, build settings may be ignored)",
                  "example": false,
                  "type": "boolean"
                },
                "installCommand": {
                  "description": "Command to install dependencies",
                  "example": "npm install",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "port": {
                  "description": "Port number the application listens on",
                  "example": 80,
                  "maximum": 65535,
                  "minimum": 1,
                  "type": [
                    "integer",
                    "null"
                  ]
                },
                "runCommand": {
                  "description": "Command to run the application",
                  "example": "npm run start",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "runEnvs": {
                  "additionalProperties": {
                    "type": "string"
                  },
                  "description": "Runtime environment variables",
                  "example": {
                    "API_KEY": "secret-key",
                    "DATABASE_URL": "postgresql://localhost:5432/mydb"
                  },
                  "type": "object"
                },
                "runFlag": {
                  "description": "Additional flags to pass to the run command",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "runtime": {
                  "description": "Runtime environment (e.g., node:20, golang:1.25, bun:1.3, build-ai). Use \"build-ai\" to enable Build AI.",
                  "example": "build-ai",
                  "maxLength": 255,
                  "type": [
                    "string",
                    "null"
                  ]
                },
                "useBuildAI": {
                  "default": false,
                  "description": "Whether to use Build AI for automated build configuration (should be true when runtime is \"build-ai\")",
                  "example": true,
                  "type": "boolean"
                }
              },
              "type": "object"
            }
          ]
        },
        "source": {
          "description": "Source configuration for the project. For VCS projects: must include vcsName, vcsInstallationId, and vcsRepoId. For image and upload projects: should be an empty object {}.",
          "oneOf": [
            {
              "additionalProperties": false,
              "description": "Source configuration for VCS projects",
              "properties": {
                "vcsInstallationId": {
                  "description": "GitHub App installation ID (must be connected to user's account)",
                  "example": "12345678",
                  "type": "string"
                },
                "vcsName": {
                  "description": "Version control system name",
                  "enum": [
                    "github"
                  ],
                  "example": "github",
                  "type": "string"
                },
                "vcsRepoId": {
                  "description": "GitHub repository ID",
                  "example": "98765432",
                  "type": "string"
                }
              },
              "required": [
                "vcsName",
                "vcsInstallationId",
                "vcsRepoId"
              ],
              "type": "object"
            },
            {
              "additionalProperties": false,
              "description": "Source configuration for image and upload projects (empty object)",
              "example": {},
              "type": "object"
            }
          ]
        },
        "type": {
          "description": "Type of project source",
          "enum": [
            "vcs",
            "image",
            "upload"
          ],
          "example": "vcs",
          "type": "string"
        },
        "uniqueName": {
          "description": "Unique identifier for the project (per user). Must be unique across all projects for the user.",
          "example": "my-awesome-app",
          "maxLength": 32,
          "minLength": 4,
          "pattern": "^[a-zA-Z0-9-]+$",
          "type": "string"
        }
      },
      "required": [
        "uniqueName",
        "displayName",
        "type",
        "settings"
      ],
      "type": "object"
    }
  },
  "required": [
    "body"
  ],
  "type": "object"
}`

// Response Template for the CreateProject tool (Status: 201, Content-Type: application/json)
const CreateProjectResponseTemplate_A = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 201

**Content-Type:** application/json

> Project created successfully

## Response Structure

- Structure (Type: object):
  - **data** (Type: object):
    - **id**: Unique identifier of the created project (Type: string, uuid):
        - Example: '550e8400-e29b-41d4-a716-446655440000'
  - **status** (Type: string):
      - Example: 'success'
`

// Response Template for the CreateProject tool (Status: 400, Content-Type: application/json)
const CreateProjectResponseTemplate_B = `# API Response Information

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

// Response Template for the CreateProject tool (Status: 401, Content-Type: application/json)
const CreateProjectResponseTemplate_C = `# API Response Information

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

// Response Template for the CreateProject tool (Status: 404, Content-Type: application/json)
const CreateProjectResponseTemplate_D = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 404

**Content-Type:** application/json

> Not found - GitHub installation not found

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the CreateProject tool (Status: 409, Content-Type: application/json)
const CreateProjectResponseTemplate_E = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 409

**Content-Type:** application/json

> Conflict - project with unique name already exists

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the CreateProject tool (Status: 429, Content-Type: application/json)
const CreateProjectResponseTemplate_F = `# API Response Information

Below is the response template for this API endpoint.

The template shows a possible response, including its status code and content type, to help you understand and generate correct outputs.

**Status Code:** 429

**Content-Type:** application/json

> Too many requests - maximum projects per user reached

## Response Structure

- Structure (Type: object):
  - **message**: Error message describing what went wrong (Type: string):
      - Example: 'invalid uniqueName'
  - **success** (Type: boolean):
      - Example: 'false'
`

// Response Template for the CreateProject tool (Status: 500, Content-Type: application/json)
const CreateProjectResponseTemplate_G = `# API Response Information

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

// NewCreateProjectMCPTool creates the MCP Tool instance for CreateProject
func NewCreateProjectMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"CreateProject",
		"Create a new project - Creates a new project with source code configuration and build settings. The project can be linked to a VCS repository (e.g., GitHub) and configured with framework/runtime settings for automated builds and deployments.",
		[]byte(createProjectInputSchema),
	)
}

// CreateProjectHandler is the handler function for the CreateProject tool.
// This function is automatically generated. Users should implement the actual
// logic within this function body to integrate with backend APIs.
// You can generate types, http client and helpers for parsing request params to facilitate the implementation.
func CreateProjectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// IMPORTANT: Replace the following placeholder implementation with your actual logic.
	// Use the 'request' parameter to access tool call arguments.
	// Make HTTP calls or interact with services as needed.
	// Return an *mcp.CallToolResult with the response payload, or an error.

	// Example placeholder implementation:
	// Extract the parameters from the request and parse them.
	// Call your backend API or perform the necessary operations using 'params'.
	// Handle the response and errors accordingly.
	return nil, fmt.Errorf("%s not implemented", "CreateProject")
}
