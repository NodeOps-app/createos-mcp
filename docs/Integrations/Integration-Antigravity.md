---
title: Antigravity
order: 10
---
# CreateOS MCP Integration with Antigravity

## Overview

Integrate CreateOS with Antigravity using the MCP (Model Context Protocol) to enable AI-powered deployment workflows directly from the Antigravity agent interface.

## Prerequisites

- Antigravity application installed
- CreateOS account with API key access

## Setup Instructions

### Step 1: Access MCP Server Settings

1. Open the **Agent chat window** in Antigravity
2. Click on the **"…"** menu in the top right corner
3. Select **MCP Servers** from the dropdown list

![Antigravity1](https://ik.imagekit.io/nodeops/Docs%20Images/Antigravity/1.png?updatedAt=1769668524785)

### Step 2: Open Configuration File

1. Click on **"Manage MCP Servers"**

![Antigravity2](https://ik.imagekit.io/nodeops/Docs%20Images/Antigravity/2.png?updatedAt=1769668524890)

2. In the menu that opens, click **"View raw config"**

![Antigravity3](https://ik.imagekit.io/nodeops/Docs%20Images/Antigravity/3.png?updatedAt=1769668524710)

3. This will open the `mcp_config.json` file

### Step 3: Configure CreateOS MCP

Paste the following configuration into `mcp_config.json`, replacing `CREATEOS_API_KEY` with your actual API key:

```json
{
    "mcpServers": {
        "createos": {
            "serverUrl": "https://api-createos.nodeops.network/mcp",
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

1. Save the `mcp_config.json` file
2. Restart the Antigravity application completely

## Verification

After restarting, Antigravity will automatically connect to the CreateOS MCP server. You can verify the connection by asking the agent to list your CreateOS projects or perform a deployment action.

## Usage

Once configured, simply tell the Antigravity agent what you want to deploy through CreateOS, and it will handle the integration automatically.