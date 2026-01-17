### Development

We are using swagger to generate mcp server.

```bash
/Users/bhautik/.asdf/installs/golang/1.23.5/packages/bin/mcpgen \
  --input /Users/bhautik/workspace/src/autogen-backend-v2/pkg/swaggerdocs/openapi.yaml \
  -package mcptools \
  -validation \
  -output /Users/bhautik/workspace/src/autogen-v2-mcp
```

All MCP handlers are written in `handlers` directory.

### Testing

```
# Start MCP server
go run *.go --transport http --port 8080

# To test mcp server using client (without LLM)
npx @modelcontextprotocol/inspector

# This will open inspector in browser
# Make sure to set `X-Api-Key: TOKEN` header in inspector
```
