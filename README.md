[![GitHub Stars](https://img.shields.io/github/stars/NodeOps-app/createos-mcp?style=flat&logo=github)](https://github.com/NodeOps-app/createos-mcp/stargazers)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![Build Status](https://img.shields.io/github/actions/workflow/status/NodeOps-app/createos-mcp/ci.yaml?branch=main&logo=github)](https://github.com/NodeOps-app/createos-mcp/actions)
[![Last Commit](https://img.shields.io/github/last-commit/NodeOps-app/createos-mcp?logo=github)](https://github.com/NodeOps-app/createos-mcp/commits/main)

# CreateOS MCP

**Deploy, manage, and scale applications directly from your AI assistant.**

---

## Introduction

The [Model Context Protocol (MCP)](https://modelcontextprotocol.io) is an open standard originated by Anthropic that enables AI assistants — Claude, Cursor, GitHub Copilot, Windsurf, Gemini, and others — to securely connect to external tools, APIs, and services through a unified interface. The latest MCP Authorization and Streamable HTTP specifications are fully implemented.

**CreateOS MCP** is a production-grade MCP server that exposes **85+ tools** for full-stack application deployment and infrastructure management on the [CreateOS](https://createos.nodeops.network) platform. Connect it once to your AI coding tool, then deploy projects, manage environments, configure domains, run security scans, analyze deployment metrics, and more — all through natural language.

Instead of switching between dashboards, CLIs, and documentation, you stay in your editor and let your AI handle the infrastructure. CreateOS MCP turns prompts like *"deploy my app from this GitHub repo"* or *"scale the staging environment to 3 replicas"* into real actions, executed instantly.

Built in Go for performance and reliability, the server supports both **Streamable HTTP** and **stdio** transports, with **OAuth 2.0** and **API key** authentication out of the box.

---

## Features

- 🚀 **85+ MCP Tools** — Full coverage of the CreateOS platform API: projects, deployments, environments, domains, templates, and more
- 🔌 **9 Supported Clients** — Cursor, VS Code + Copilot, Claude Desktop, Claude Code, Windsurf, Gemini CLI, Gemini Code Assist, Opencode, Zapier, and ElevenLabs
- 🔐 **Secure Authentication** — API key and OAuth 2.0 with Dynamic Client Registration (RFC 7591), Protected Resource Metadata (RFC 9728)
- ⚡ **Dual Transport** — Streamable HTTP for remote access, stdio for local/embedded use
- 📦 **Deploy Anything** — GitHub repos, Docker images, file uploads (zip, base64, multipart), or marketplace templates
- 🌍 **Environment Management** — Create, configure, and monitor multiple environments per project with independent secrets and resources
- 🔒 **Security Scanning** — Trigger and review security scans on any deployment
- 📊 **Built-in Analytics** — Query request metrics, RPM, error paths, success rates, and traffic distribution per environment
- 🏪 **Template Marketplace** — Browse, purchase, and deploy from a library of production-ready project templates
- 🔑 **API Key Management** — Create, list, update, and revoke API keys programmatically
- 🌐 **Custom Domains** — Attach, refresh, and manage custom domains with automatic TLS
- 📋 **Build & Runtime Logs** — Stream build logs and runtime logs directly in your AI assistant

---

## Prerequisites

1. Create an account at [createos.nodeops.network](https://createos.nodeops.network)
2. Navigate to **Profile Settings** and generate your API key

> 📺 [Watch the API key retrieval tutorial](https://ik.imagekit.io/nodeops/No%20Sound/API%20Key_Retrieval.mp4)

---

## Usage Examples

Once connected, interact with CreateOS through natural language in your AI assistant:

### Deploying

```
"Deploy my current project to CreateOS"
"Create a new project called my-api using Node.js 20"
"Deploy the main branch of my GitHub repo naman485/my-app"
"Upload this zip file as a new deployment"
```

### Managing Environments

```
"List all environments for the my-api project"
"Set the DATABASE_URL environment variable on staging"
"Scale the production environment to 500 millicores and 1024MB memory"
"Show me the runtime logs for staging"
```

### Domains & Networking

```
"Add a custom domain api.example.com to my project"
"Refresh the SSL certificate for my domain"
"List all domains attached to my-api"
```

### Monitoring & Security

```
"Show me the request analytics for production this week"
"What are the top error paths in staging?"
"Run a security scan on my latest deployment"
"Show me the build logs for deployment abc123"
```

### Templates & Marketplace

```
"Browse available project templates"
"Deploy the FastAPI starter template"
"Create a project template from my current project"
```

---

## Configuration

### Hosted (Recommended)

The hosted MCP server is available at:

```
https://api-createos.nodeops.network/mcp
```

No setup required — just add your API key to your MCP client config as shown in the docs in the Integration folder.

### Self-Hosting

Clone and build the server:

```bash
git clone https://github.com/NodeOps-app/createos-mcp.git
cd createos-mcp
go build -o createos-mcp *.go
```

Create a `config.yaml`:

```yaml
port: 8080
base_url: https://your-domain.com
authorization_server_url: https://your-auth-server.com
api_base_url: https://api.createos.nodeops.network
transport: http       # "http" or "stdio"
log_level: debug

supported_scopes:
  - openid
  - offline_access
  - offline

response_types_supported:
  - code

grant_types_supported:
  - authorization_code
  - refresh_token

code_challenge_methods_supported:
  - S256

token_endpoint_auth_methods_supported:
  - none
```

Run the server:

```bash
./createos-mcp --config config.yaml
```

### Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `PORT` | Server port (overrides config) | No (default: `8080`) |

### Testing Locally

```bash
# Start the server
go run *.go --config config.yaml

# Open the MCP Inspector
npx @modelcontextprotocol/inspector

# In the inspector, set the URL to http://localhost:8080/mcp
# and add header X-Api-Key: YOUR_API_KEY
```

---

## Supported Tools

The server exposes **85+ tools** organized into the following categories:

| Category | Tools | Description |
|----------|-------|-------------|
| **Projects** | `CreateProject`, `GetProject`, `ListProjects`, `UpdateProject`, `DeleteProject`, `UpdateProjectSettings`, `CheckProjectUniqueName` | Create and manage projects with full lifecycle control |
| **Deployments** | `CreateDeployment`, `GetDeployment`, `ListDeployments`, `CancelDeployment`, `RetriggerDeployment`, `TriggerLatestDeployment`, `AssignDeploymentToProjectEnvironment` | Deploy code via GitHub, Docker, file upload, or zip |
| **File Uploads** | `UploadDeploymentFiles`, `UploadDeploymentBase64Files`, `UploadDeploymentZip`, `DownloadDeployment` | Multiple upload strategies for code deployment |
| **Environments** | `CreateProjectEnvironment`, `ListProjectEnvironments`, `UpdateProjectEnvironment`, `DeleteProjectEnvironment`, `UpdateProjectEnvironmentEnvironmentVariables`, `UpdateProjectEnvironmentResources` | Manage staging, production, and custom environments |
| **Domains** | `CreateDomain`, `ListDomains`, `DeleteDomain`, `RefreshDomain`, `UpdateDomainEnvironment` | Custom domains with automatic TLS |
| **Apps** | `CreateApp`, `ListApps`, `UpdateApp`, `DeleteApp`, `AddProjectsToApp`, `RemoveProjectsFromApp`, `AddServicesToApp`, `RemoveServicesFromApp`, `ListProjectsByApp`, `ListServicesByApp` | Group projects and services into logical applications |
| **Templates** | `CreateProjectTemplate`, `GetProjectTemplate`, `ListPublishedProjectTemplates`, `ListMyProjectTemplates`, `BuyProjectTemplate`, `DeployProjectTemplateViaGithub`, `UpdateProjectTemplate`, `DeleteProjectTemplate`, `DownloadProjectTemplate` | Template marketplace for reusable project blueprints |
| **API Keys** | `CreateAPIKey`, `ListAPIKeys`, `UpdateAPIKey`, `RevokeAPIKey`, `CheckAPIKeyUniqueName` | Programmatic API key management |
| **Security** | `TriggerSecurityScan`, `RetriggerSecurityScan`, `GetSecurityScan`, `GetSecurityScanDownloadUri` | Automated security scanning for deployments |
| **Analytics** | `GetProjectEnvironmentAnalytics`, `OverallRequests`, `RPM`, `RequestDistribution`, `RequestsOverTime`, `SuccessPercentage`, `TopErrorPaths`, `TopHitPaths` | Real-time request metrics and traffic analysis |
| **Logs** | `GetBuildLogs`, `GetDeploymentLogs`, `GetProjectEnvironmentLogs` | Build-time and runtime log streaming |
| **GitHub** | `InstallGithubApp`, `ListConnectedGithubAccounts`, `ListGithubRepositories`, `ListGithubRepositoryBranches`, `GetGithubRepositoryContent` | GitHub integration for repo-based deployments |
| **Account** | `GetCurrentUser`, `GetQuotas`, `GetSupportedProjectTypes` | User info, quotas, and platform capabilities |
| **Transfers** | `TransferProject`, `GetProjectTransferUri`, `ListProjectTransferHistory` | Transfer project ownership between accounts |

---

## Roadmap

- [ ] Database management tools (PostgreSQL, MySQL provisioning and management)
- [ ] Managed Kafka, Valkey/Redis integration tools
- [ ] Webhook and event-driven deployment triggers
- [ ] Team collaboration and role-based access tools
- [ ] Cost estimation and billing analytics tools
- [ ] SSE transport support
- [ ] Deployment rollback tools

---

## Contributing

Contributions are welcome! Here's how to get started:

1. **Fork** the repository
2. **Create** a feature branch: `git checkout -b feat/my-feature`
3. **Make** your changes and add tests
4. **Run** tests: `go test ./...`
5. **Commit** with a clear message: `git commit -m "feat: add my feature"`
6. **Push** and open a Pull Request

Please read our [Contributing Guidelines](CONTRIBUTING.md) if available.

### Development

The MCP tools and handlers are auto-generated from the CreateOS OpenAPI spec using `mcpgen`. See [DEVELOPMENT.md](DEVELOPMENT.md) for details.

---

## Code of Conduct

This project follows the [Contributor Covenant Code of Conduct](https://www.contributor-covenant.org/version/2/1/code_of_conduct/). By participating, you are expected to uphold this code.

---

## License

This project is licensed under the **MIT License**.

```
MIT License

Copyright (c) 2025 NodeOps

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

---

## Acknowledgments

- [Anthropic](https://anthropic.com) for creating the [Model Context Protocol](https://modelcontextprotocol.io) specification
- [mark3labs/mcp-go](https://github.com/mark3labs/mcp-go) for the excellent Go MCP SDK
- [Cursor](https://cursor.com), [GitHub Copilot](https://github.com/features/copilot), [Windsurf](https://windsurf.com), [Google Gemini](https://gemini.google.com), and the broader AI tooling ecosystem for driving MCP adoption
- The [CreateOS](https://createos.nodeops.network) team for building the deployment platform

---

<p align="center">
  <a href="https://nodeops.network/createos/docs/API-MCP/CreateOS-MCP">📚 Full Documentation</a> &nbsp;·&nbsp;
  <a href="https://createos.nodeops.network">🌐 CreateOS Platform</a> &nbsp;·&nbsp;
  <a href="https://github.com/NodeOps-app/createos-mcp/issues">🐛 Report Bug</a>
</p>

<p align="center">
  If you find this useful, give it a star ✨
  <br><br>
  <a href="https://github.com/NodeOps-app/createos-mcp">
    <img src="https://img.shields.io/github/stars/NodeOps-app/createos-mcp?style=social" alt="GitHub Stars">
  </a>
</p>
