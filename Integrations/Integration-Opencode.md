---
title: Opencode
order: 8
---

# Opencode Desktop Integration Guide

## Overview

Integrate CreateOS with Opencode Desktop using the MCP (Model Context Protocol) to enable AI-powered deployment workflows directly from your development environment.

---

## Prerequisites

- Active Opencode Desktop installation
- CreateOS account with API key access
- Project directory set up in Opencode

---

## Setup Instructions

### Step 1: Open Your Project

1. Launch **Opencode Desktop**
2. Select or open your project directory

### Step 2: Create Configuration File

1. In your project root directory, create a new file named `opencode.json`
2. The file path should be: `<project-directory>/opencode.json`

### Step 3: Configure Connection Settings

Paste the following configuration into your `opencode.json` file:
```json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {
    "createos": {
      "type": "remote",
      "url": "https://api-createos.nodeops.network/mcp",
      "headers": {
        "X-Api-Key": "CREATEOS_API_KEY"
      }
    }
  }
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

**Replace** `"CREATEOS_API_KEY"` with your actual CreateOS API key in the configuration.

### Step 4: Apply Configuration

1. Save the `opencode.json` file
2. **Restart Opencode Desktop** to apply the configuration changes
![Opencode1](https://ik.imagekit.io/nodeops/Docs%20Images/OpenCode/1.png?updatedAt=1769668492745)

### Step 5: Verify Connection

1. After restarting, you should see the **MCP connection indicator** in Opencode
2. Click on the **MCP connection indicator**
3. Verify that **"createos"** is shown as **enabled**
![Opencode2](https://ik.imagekit.io/nodeops/Docs%20Images/OpenCode/2.png?updatedAt=1769668492736)

---

## Troubleshooting

If the server doesn't appear as connected, confirm:

- `opencode.json` file is in the correct project root directory
- JSON syntax is valid (no trailing commas or formatting errors)
- API key is active and has proper permissions
- Opencode Desktop was fully restarted

---

## Next Steps

Once connected, you can:

- Use AI prompts to deploy services directly from Opencode
- Access CreateOS infrastructure through the MCP interface
- Manage deployments without leaving your development environment
