# CreateOS MCP Server

<a href="https://glama.ai/mcp/servers/NodeOps-app/createos-mcp">
  <img width="380" height="200" src="https://glama.ai/mcp/servers/NodeOps-app/createos-mcp/badge" alt="CreateOS MCP server" />
</a>

## Configuration

### Via API Key

```json
{
  "servers": {
    "autogen mcp": {
      "url": "https://api-createos.nodeops.network/mcp",
      "type": "http",
      "headers": {
        "X-Api-Key": "REPLACE-API-KEY"
      }
    }
  },
  "inputs": []
}
```