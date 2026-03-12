---
title: Zapier
order: 11
---

# Zapier MCP Integration Guide

## Overview

Integrate CreateOS with Zapier using the MCP (Model Context Protocol) Client to automate workflows and connect your deployments with thousands of applications.

---

## Prerequisites

- Active Zapier account
- CreateOS account with API key access
- Administrative permissions in Zapier

---

## Setup Instructions

### Step 1: Access App Connections

1. Log in to the <a href="https://zapier.com" target="_blank">Zapier web portal</a>
2. Click **"App Connections"** in the bottom left navigation panel
![Zapier1](https://ik.imagekit.io/nodeops/Docs%20Images/Zapier/1.png?updatedAt=1769668492796)

### Step 2: Add New Connection

1. Click **"Add Connection"** in the top right corner
![Zapier2](https://ik.imagekit.io/nodeops/Docs%20Images/Zapier/2.png?updatedAt=1769668492748)
2. Search for **"MCP Client by Zapier"** and select it
![Zapier3](https://ik.imagekit.io/nodeops/Docs%20Images/Zapier/3.png?updatedAt=1769668492754)

### Step 3: Configure Connection Settings

Enter the following information in the configuration form:

| Field | Value |
|-------|-------|
| **Server URL** | `https://api-createos.nodeops.network/mcp` |
| **Transport** | Streamable HTTP |
| **OAuth** | No |
| **Bearer Token** | Your CreateOS API Key |

![Zapier4](https://ik.imagekit.io/nodeops/Docs%20Images/Zapier/4.png?updatedAt=1769668492793)
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

### Step 4: Verify Connection

1. Click **"Test"** or **"Save"** to verify the connection
2. Upon successful verification, the connection will appear in your **App Connections** list

![Zapier5](https://ik.imagekit.io/nodeops/Docs%20Images/Zapier/5.png?updatedAt=1769668492787)

