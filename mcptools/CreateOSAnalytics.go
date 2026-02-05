package mcptools

import "github.com/mark3labs/mcp-go/mcp"

const createOSProjectEnvironmentAnalyticsInputSchema = `{
  "type": "object",
  "required": ["project_id", "environment_id", "metric"],
  "properties": {
    "project_id": {"type": "string", "description": "Project ID (UUID)."},
    "environment_id": {"type": "string", "description": "Environment ID (UUID)."},
    "metric": {
      "type": "string",
      "enum": [
        "requests_over_time",
        "rpm",
        "success_percentage",
        "request_distribution",
        "top_error_paths",
        "top_hit_paths",
        "overall_requests"
      ],
      "description": "Analytics metric to fetch."
    },
    "start": {"type": "integer", "description": "Start timestamp (unix)."},
    "end": {"type": "integer", "description": "End timestamp (unix)."}
  }
}`

func NewCreateOSProjectEnvironmentAnalyticsMCPTool() mcp.Tool {
	return mcp.NewToolWithRawSchema(
		"createos.project_env.analytics",
		"Fetch project environment analytics by metric. Use when the user asks for traffic or error stats.",
		[]byte(createOSProjectEnvironmentAnalyticsInputSchema),
	)
}
