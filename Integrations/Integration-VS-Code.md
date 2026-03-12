---
title: VS Code with Copilot
order: 7
---

# VS Code with Copilot MCP Integration

[Click to Add CreateOS MCP](vscode:mcp/install?%7B%22name%22%3A%22CreateOS%22%2C%22url%22%3A%22https%3A%2F%2Fapi-createos.nodeops.network%2Fmcp%22%7D)

Through the button:
1. Click on Install
2. Allow when prompted, it will take you to OAuth page
3. Login using your CreateOS credentials

## Overview

Integrate CreateOS with VS Code using the MCP (Model Context Protocol) to enable AI-powered deployment workflows directly from your development environment.

---

## Prerequisites

- Active VS Code installation
- CreateOS account with API key access
- GitHub connected to VS Code

---

## Setup Instructions

### Step 1: Create MCP Configuration File

1. Open your workspace in VS Code
2. Create a new folder named `.vscode` in your workspace root (if it doesn't already exist)
3. Inside the `.vscode` folder, create a new file named `mcp.json`

### Step 2: Configure Connection Settings

Paste the following configuration into your `mcp.json` file:

```json
{
  "servers": {
    "createos": {
      "url": "https://api-createos.nodeops.network/mcp",
      "type": "http",
      "headers": {
        "X-Api-Key": "CREATEOS_API_KEY"
      }
    }
  },
  "inputs": []
}
```

**Getting Your API Key:**
1. Log in to CreateOS
2. Navigate to Profile Settings
3. Copy your API key from the API Key section

<br/>

<a href="https://ik.imagekit.io/nodeops/No%20Sound/API%20Key%20Retrieval.mp4" target="_blank">Watch Tutorial →</a>

<MDXVideo
  src="https://ik.imagekit.io/nodeops/No%20Sound/API%20Key%20Retrieval.mp4"
  caption="API key retrieval tutorial"
/>

**Security Note:** Keep your API key confidential and never share it publicly.

Replace `"CREATEOS_API_KEY"` with your actual CreateOS API key in the configuration.

### Step 3: Apply Configuration

1. Save the `mcp.json` file
2. Restart VS Code to apply the configuration changes

---

## Demo Video

Vibecode and deploy through CreateOS MCP: **<a href="https://ik.imagekit.io/nodeops/No%20Sound/VS%20Code%20Copilot%20Demo.mp4" target="_blank">Demo</a>**
<MDXVideo
  src="https://ik.imagekit.io/nodeops/No%20Sound/VS%20Code%20Copilot%20Demo.mp4"
  caption="VS Code with Copilot Vibecode Demo"
/>


