---
title: Cursor
order: 4
---

# Cursor MCP Integration

[Click to Add CreateOS MCP](cursor://anysphere.cursor-deeplink/mcp/install?name=createos&config=eyJ1cmwiOiJodHRwczovL2FwaS1jcmVhdGVvcy5ub2Rlb3BzLm5ldHdvcmsvbWNwIiwidHlwZSI6Imh0dHAiLCJoZWFkZXJzIjp7IlgtQXBpLUtleSI6IkNSRUFURU9TX0FQSV9LRVkifSwiaW5wdXRzIjpbXX0=)

## Overview

Below is a comprehensive guide to add CreateOS MCP. If required, see the <a href="https://docs.cursor.com" target="_blank">Cursor documentation</a> for more details. Integrate CreateOS with Cursor IDE using the MCP (Model Context Protocol) to enable AI-powered deployment workflows directly from your development environment.

---

## Prerequisites

- Active Cursor IDE installation
- CreateOS account with API key access
- Git configured and connected to Cursor
- Latest code changes committed and pushed

---

## Setup Instructions

### Step 1: Access MCP Settings

1. Open Cursor IDE
2. Navigate to **Settings → MCP**

![Cursor1](https://ik.imagekit.io/nodeops/Docs%20Images/Cursor/1.png?updatedAt=1769668492725)

### Step 2: Add New MCP Server

1. Click **"New MCP Server"**
2. Cursor will open your `mcp.json` configuration file

### Step 3: Configure Connection Settings

Paste the following configuration into your `mcp.json` file:

```json
{
  "mcpServers": {
    "createos": {
      "url": "https://api-createos.nodeops.network/mcp",
      "type": "http",
      "headers": {
        "X-Api-Key": "CREATEOS_API_KEY"
      },
      "inputs": []
    }
  }
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

### Step 4: Apply Configuration

1. Save the `mcp.json` file
2. Restart Cursor IDE to apply the configuration changes

### Step 5: Verify Installation

1. Go to **Settings → Tools & MCP**
![Cursor2](https://ik.imagekit.io/nodeops/Docs%20Images/Cursor/2.png?updatedAt=1769668492733)
2. Under "Installed MCP Servers", locate **"createos"**
3. If not enabled, click on it to enable the server

**If the server doesn't appear, confirm:**
- `mcp.json` syntax is correct
- API key is valid and active
- Cursor IDE was fully restarted
- Git is properly configured and connected

### Step 6: Start Using CreateOS

Once configured, you can prompt Cursor to deploy using CreateOS MCP or access its tools directly through the AI assistant.

For more details, see the <a href="https://docs.cursor.com" target="_blank">Cursor documentation</a>.

---

## Demo Video

Vibecode and deploy through CreateOS MCP: **<a href="https://ik.imagekit.io/nodeops/No%20Sound/Cursor%20MCP%20Demo.mp4" target="_blank">Demo</a>**
<MDXVideo
  src="https://ik.imagekit.io/nodeops/No%20Sound/Cursor%20MCP%20Demo.mp4"
  caption="Cursor Vibecode Demo"
/>

