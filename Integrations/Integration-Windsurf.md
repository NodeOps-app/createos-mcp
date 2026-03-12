---
title: Windsurf
order: 9
---

# Windsurf MCP Integration

## Overview

Integrate CreateOS with Windsurf IDE using the MCP (Model Context Protocol) to enable AI-powered deployment workflows directly from your development environment.

---

## Prerequisites

- Active Windsurf IDE installation
- CreateOS account with API key access

---

## Setup Instructions

### Step 1: Access MCP Settings

1. Open Windsurf IDE
2. Navigate to **Windsurf Settings**
3. Go to **Cascade → MCP Servers → Open Marketplace**
![Windsurf1](https://ik.imagekit.io/nodeops/Docs%20Images/Windsurf/1.png?updatedAt=1769668492739)

### Step 2: Open MCP Configuration

1. In the MCP Marketplace, locate the settings icon for Installed MCPs
2. Click on the settings icon
3. Windsurf will open the `mcp.json` configuration file

### Step 3: Configure Connection Settings

Paste the following configuration into your `mcp.json` file:
```json
{
  "mcpServers": {
    "createos": {
      "url": "https://api-createos.nodeops.network/mcp",
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

### Step 4: Save Configuration

1. Save the `mcp.json` file
2. Close the configuration editor

### Step 5: Verify Connection

1. Navigate back to **Cascade → MCP Servers → Open Marketplace**
2. Verify that "createos" appears in the list of installed and connected MCP servers

**If the server doesn't appear as connected, confirm:**
- `mcp.json` syntax is correct
- API key is valid and active
- Configuration file was properly saved

### Step 6: Start Using CreateOS

1. Click on the Windsurf logo to open the chat interface
![Windsurf2](https://ik.imagekit.io/nodeops/Docs%20Images/Windsurf/2.png?updatedAt=1769668492752)

2. You're now ready to use CreateOS MCP
3. Prompt Cascade to deploy your applications using CreateOS
