---
title: ElevenLabs
order: 12
---

# ElevenLabs Integration Guide

## Overview

Integrate CreateOS with ElevenLabs using the MCP (Model Context Protocol) to enable AI-powered deployment workflows directly from your ElevenLabs workspace.

---

## Prerequisites

- Active ElevenLabs account
- CreateOS account with API key access
- Access to ElevenLabs Integration settings

---

## Setup Instructions

### Step 1: Access ElevenLabs Integrations

1. Log in to your <a href="https://elevenlabs.io" target="_blank">ElevenLabs account</a>
2. Click on **"Integration"** in the left sidebar, or <a href="https://elevenlabs.io/app/integrations" target="_blank">click here</a> to go directly to the integrations page
![ElevenLabs1](https://ik.imagekit.io/nodeops/Docs%20Images/ElevenLabs/1.png?updatedAt=1769668492765)

### Step 2: Add New Integration

1. Click **"Add Integration"**
![ElevenLabs2](https://ik.imagekit.io/nodeops/Docs%20Images/ElevenLabs/2.png?updatedAt=1769668492775)
2. Select **"Custom MCP Server"**
![ElevenLabs3](https://ik.imagekit.io/nodeops/Docs%20Images/ElevenLabs/3.png?updatedAt=1769668492760)

### Step 3: Configure Basic Information

Enter the following details:

| Field | Value |
|-------|-------|
| **Name** | CreateOS (or custom name of your choice) |
| **Description** | Custom description of your choice |

![ElevenLabs4](https://ik.imagekit.io/nodeops/Docs%20Images/ElevenLabs/4.png?updatedAt=1769668492723)

### Step 4: Configure Server Settings

Enter the following server configuration:

| Field | Value |
|-------|-------|
| **Server type** | Streamable HTTP |
| **Type** | Value |
| **URL** | `https://api-createos.nodeops.network/mcp` |

![ElevenLabs5](https://ik.imagekit.io/nodeops/Docs%20Images/ElevenLabs/5.png?updatedAt=1769668492742)

### Step 5: Configure HTTP Headers

Add the authentication header:

| Field | Value |
|-------|-------|
| **Type** | Value |
| **Name** | `X-API-Key` |
| **Value** | Your CreateOS API Key |

![ElevenLabs6](https://ik.imagekit.io/nodeops/Docs%20Images/ElevenLabs/6.png?updatedAt=1769668492763)

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

### Step 6: Complete Setup

1. Check the box for **"I trust this server"**
2. Click **"Add Server"**

![ElevenLabs7](https://ik.imagekit.io/nodeops/Docs%20Images/ElevenLabs/7.png?updatedAt=1769668492750)

3. The CreateOS MCP server will now be available in your ElevenLabs integrations

---

## Troubleshooting

If the integration fails, confirm:

- Server URL is entered correctly without trailing spaces
- API key is valid and active
- "I trust this server" checkbox is checked

---

## Next Steps

Once connected, you can:

- Use CreateOS functionality within ElevenLabs workflows
- Automate deployment processes through ElevenLabs
- Access CreateOS infrastructure from your ElevenLabs projects
