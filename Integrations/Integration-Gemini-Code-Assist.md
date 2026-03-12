---
title: Gemini Code Assist
order: 6
---

# Gemini Code Assist MCP Integration

## Overview

Integrate CreateOS with Gemini Code Assist using the MCP (Model Context Protocol) to enable AI-powered deployment workflows directly from your IDE.

Gemini Code Assist and Gemini CLI share the same configuration. Once you set up the following, Gemini Code Assist will automatically connect to the CreateOS MCP.

---

## Prerequisites

- Gemini CLI installed
- CreateOS account with API key access
- Command line access (PowerShell, Terminal, or Command Prompt)

**Note:** Gemini CLI and Gemini Code Assist share the same configuration.

---

## Setup Instructions

### Step 1: Verify Gemini CLI Installation

1. Open your terminal or command prompt
2. Run the following command to check if Gemini CLI is installed:
```powershell
gemini --version
```

If not installed, install Gemini CLI according to its official installation instructions.

### Step 2: Locate Configuration Directory

Gemini CLI reads MCP configuration from:
```
<project-root>\.gemini\settings.json
```

**Example:**
```
G:\NodeOps\GeminiCLI\.gemini\settings.json
```

If the `.gemini` directory does not exist, create it.

### Step 3: Configure API Settings

1. Create or open `settings.json` in the `.gemini` directory
2. Paste the following configuration and replace `"CREATEOS_API_KEY"` with your CreateOS API key:
```json
{
  "mcpServers": {
    "createos": {
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
2. Navigate to Profile Settings
3. Copy your API key from the API Key section

<br/>

<a href="https://ik.imagekit.io/nodeops/No%20Sound/API%20Key%20Retrieval.mp4" target="_blank">Watch Tutorial →</a>

<MDXVideo
  src="https://ik.imagekit.io/nodeops/No%20Sound/API%20Key%20Retrieval.mp4"
  caption="API key retrieval tutorial"
/>

**Security Note:** Keep your API key confidential and never share it publicly.

3. Save the `settings.json` file

### Step 4: Restart Gemini CLI

1. Close Gemini completely by running:
```powershell
exit
```

2. Reopen Gemini:
```powershell
gemini
```

### Step 5: Verify Installation

Run the following command:
```powershell
gemini mcp list
```

**Expected output:**
```
✓ user-autogen  connected
```

### Step 6: Start Using CreateOS

Once configured, Gemini Code Assist will automatically have access to CreateOS MCP. \
**Sample Task:** Tell Gemini Code Assist that you want to deploy `<repo_name>` through NodeOps MCP, and it will deploy automatically.

---

## Demo Video

Watch our step-by-step walkthrough: **<a href="https://drive.google.com/file/d/1kxZMaN6SKEIV0pAKlY3aZT5rsZm-wpEx/view?usp=sharing" target="_blank">Gemini Code Assist MCP Integration Tutorial</a>**
