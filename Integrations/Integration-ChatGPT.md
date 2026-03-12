---
title: ChatGPT
order: 3
---

# ChatGPT Integration Guide

## Overview

Integrate CreateOS with ChatGPT using the MCP (Model Context Protocol) to enable AI-powered deployment workflows directly from the ChatGPT web interface.

## Prerequisites

- ChatGPT account at <a href="https://chatgpt.com" target="_blank">chatgpt.com</a>
- CreateOS account

## Setup Instructions

### Step 1: Access ChatGPT Settings

1. Log in to your ChatGPT account at <a href="https://chatgpt.com" target="_blank">chatgpt.com</a>
2. Click on your profile icon in the bottom left corner
3. Select **<a href="https://chatgpt.com/#settings" target="_blank">Settings</a>** from the list

### Step 2: Enable Developer Mode

1. Navigate to **Apps → <a href="https://chatgpt.com/#settings/Connectors/Advanced" target="_blank">Advanced Settings</a>**
![ChatGPT1](https://ik.imagekit.io/nodeops/Docs%20Images/ChatGPT/1.png?updatedAt=1769668492756)
2. Toggle on **Developer Mode**
![ChatGPT2](https://ik.imagekit.io/nodeops/Docs%20Images/ChatGPT/2.png?updatedAt=1769668492727)


### Step 3: Create New App with MCP Configuration

1. Click the **Create app** button (located above the enabled option)
2. Fill in the app details:
   - **Name:** CreateOS (or your preferred name)
   - **Description:** CreateOS deployment service (as you wish)
3. Add the following MCP URL: `https://api-createos.nodeops.network/mcp`
4. Set **Authentication** to: **OAuth**
5. Click **Create**

![ChatGPT3](https://ik.imagekit.io/nodeops/Docs%20Images/ChatGPT/3.png?updatedAt=1769668492719)

6. Log in using your CreateOS account when prompted

---

**You're all set!** Your ChatGPT instance is now connected to CreateOS via MCP.

