---
title: Google Workspace CLI
order: 6
---

# Google Workspace CLI Integration

## Overview

The **Google Workspace CLI** provides a command-line interface for working with Gmail, Docs, Sheets, and other Google Workspace services. Use it to script workflows, automate tasks, and build tools that interact with your emails, documents, and spreadsheets from your terminal.

---

## Prerequisites

- **Node.js** and **npm**
- A **Google account** with access to Google Workspace services
- A **Google OAuth application** ([Google Cloud Console](https://console.cloud.google.com/))
- A terminal (Cursor, VS Code terminal, or standard terminal)

---

## Setup Instructions

### Step 1: Install the CLI

Open your terminal and run:

```bash
npm install -g @googleworkspace/cli
```

Once installed, the `gws` command is available in your terminal.

**Video: Installing Google Workspace CLI**

<MDXVideo
  src="https://ik.imagekit.io/4v04d3yhr/CreateOS/gws%20cli%20install.mov"
  caption="Installing Google Workspace CLI"
/>

### Step 2: Setup Authentication

Run:

```bash
gws auth setup
```

If the CLI reports a missing dependency, install it when prompted and run the command again.

Credentials are typically stored at:

```
~/.config/gws/client_secret.json
```

You may need to create or configure a **Google OAuth client** depending on your environment.

### Step 3: Login

Authenticate with your Google account:

```bash
gws auth login --account your_email_here@gmail.com
```

**Video: Authentication Setup**

<MDXVideo
  src="https://ik.imagekit.io/4v04d3yhr/CreateOS/gws%20cli%20setup%20and%20login.mov"
  caption="Google Workspace CLI setup and login"
/>

### Step 4: Verify Setup

List files from Google Drive:

```bash
gws drive files list --params '{"pageSize": 5}'
```

Send a test email from the terminal:

```bash
gws gmail +send --to your_email_here@gmail.com --subject 'from terminal' --body 'hello! from the terminal!'
```

**Video: Testing the CLI**

<MDXVideo
  src="https://ik.imagekit.io/4v04d3yhr/CreateOS/gws%20cli%20test%20and%20send%20email.mov"
  caption="Testing Google Workspace CLI and sending email"
/>

### Step 5: Troubleshooting Authentication

If commands do not return results, check authentication status:

```bash
gws auth list
```

If no credentials appear, export them manually:

```bash
gws auth export --unmasked > ~/gws-credentials.json
export GOOGLE_WORKSPACE_CLI_CREDENTIALS_FILE=~/gws-credentials.json
```

Then run the verification command again:

```bash
gws drive files list --params '{"pageSize": 5}'
```

---

## Example Commands

View all available commands:

```bash
gws --help
```

The CLI includes **100+ prebuilt skills** for different Google Workspace APIs. Full list: <a href="https://github.com/googleworkspace/cli/blob/main/docs/skills.md" target="_blank">skills.md</a>

**Send an email:**

```bash
gws gmail +send --to example@email.com --subject "From terminal" --body "Hello from the terminal!"
```

**Get Gmail profile:**

```bash
gws gmail users getProfile --params '{"userId": "me"}'
```

---

## What You Can Build

- Automated email workflows
- AI agents that use Workspace tools
- Scripts that read or update Google Sheets
- Internal tools powered by Workspace APIs

---

## Demo Video

End to end install: <a href="https://ik.imagekit.io/4v04d3yhr/CreateOS/gws%20cli.mov" target="_blank">Demo</a>

<MDXVideo
  src="https://ik.imagekit.io/4v04d3yhr/CreateOS/gws%20cli.mov"
  caption="Google Workspace CLI demo"
/>
