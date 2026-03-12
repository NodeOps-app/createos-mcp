---
title: Claude Code
order: 1
---

# Claude Code Integration Guide

## Overview

Integrate CreateOS with Claude Code using the MCP (Model Context Protocol) to enable AI-powered deployment workflows directly from your command line.

---

## Prerequisites

- Claude Code installed
- CreateOS account with API key access
- Command line access (PowerShell, Terminal, or Command Prompt)

---

## Setup Instructions

### Step 1: Verify Claude Code Installation

1. Open your terminal or command prompt
2. Run the following command to check if Claude Code is installed:
```powershell
claude --version
```

If not installed, install Claude Code according to its <a href="https://code.claude.com/docs/en/mcp" target="_blank">official installation instructions</a>.

### Step 2: Authenticate Claude Code

1. Log in to Claude Code by running:
```powershell
claude auth login
```

2. Verify login:
```powershell
claude auth status
```

Ensure you're successfully authenticated before proceeding.

### Step 3: Configure CreateOS MCP Connection

Run the following command, replacing `"CREATEOS_API_KEY"` with your actual CreateOS API key:
```powershell
claude mcp add createos-mcp --transport http https://api-createos.nodeops.network/mcp --header "X-API-Key: CREATEOS_API_KEY"
```

**Getting Your API Key:**

1. Log in to CreateOS
2. Navigate to **Profile Settings**
3. Copy your API key from the **API Key** section

<br/>

<a href="https://ik.imagekit.io/nodeops/No%20Sound/API%20Key%20Retrieval.mp4" target="_blank">Watch Tutorial →</a>

<MDXVideo
  src="https://ik.imagekit.io/nodeops/No%20Sound/API%20Key%20Retrieval.mp4"
  caption="API key retrieval tutorial"
/>

**Security Note**: Keep your API key confidential and never share it publicly.

### Step 4: Apply Configuration

1. After adding the MCP server, restart Claude Code
2. Close any active Claude sessions and reopen

### Step 5: Verify Connection

Run the following command to verify the MCP connection:
```powershell
claude mcp list
```

**In Claude Desktop**, if you've already initiated the `claude` command, use:
```
/mcp
```

**Expected Output:**
```
✓ createos-mcp connected
```

---

## Demo Video

Vibecode and deploy through CreateOS MCP: **<a href="https://ik.imagekit.io/nodeops/No%20Sound/Claude%20Code%20MCP%20Demo.mp4" target="_blank">Demo</a>**
<MDXVideo
  src="https://ik.imagekit.io/nodeops/No%20Sound/Claude%20Code%20MCP%20Demo.mp4"
  caption="Claude Code Vibecode Demo"
/>
---

## Troubleshooting

If the connection fails, confirm:

- Claude Code is properly installed and authenticated
- API key is valid and has proper permissions
- Command syntax is correct (no extra spaces or characters)
- Claude Code has been fully restarted after adding the MCP server

---

## Next Steps

Once connected, you can:

- Use AI prompts to deploy services directly from the command line
- Access CreateOS infrastructure through Claude Code
- Manage deployments with natural language commands
