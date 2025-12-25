# 📦 Stash

A lightweight, fully CLI-based version control system built from scratch in Go.

## Overview

Stash is a simple yet powerful tool that lets you track file changes, create snapshots, and push to a remote server — all from the command line. No GUI, no external dependencies, just a clean terminal experience.

## ✨ Features

- **100% CLI** — Everything happens in the terminal, no GUI required
- **Track Files** — Watch specific files or entire directories for changes
- **Create Snapshots** — Store the current state of tracked files
- **Global Config** — Set your username and email once, use everywhere
- **Project-based** — Each project gets its own `.stash` folder
- **Remote Sync** — Push and pull changes to a remote server directly from the CLI

## 🚀 Installation

```bash
# Download the repository and navigate to it
cd stash/client

# Build the CLI
go build -o stash .

# Move to PATH (optional)
sudo mv stash /usr/local/bin/
```

## 📖 Usage

### Initialize a Project

```bash
# Create a new stash project in current directory
stash create

# Or with a custom project name
stash create my-project
```

### Configure User (Global)

```bash
# Set your username
stash config user.name "Your Name"

# Set your email
stash config user.email "you@example.com"
```

### Track Files

```bash
# Track a specific file
stash watch main.go

# Track all files in the project
stash watch all
```

### Store Changes

```bash
# Create a snapshot with a message
stash store "Initial commit"
```

### Authentication

```bash
# Login to remote server
stash login
```

### Get Help

```bash
# List all available commands
stash help
```

## 📁 Project Structure

```
.stash/
├── projectConfig.json    # Project configuration
└── store/                # Stored snapshots
```

### projectConfig.json

```json
{
  "projectId": "uuid",
  "projectName": "my-project",
  "trackedFile": ["main.go", "src/app.go"],
  "role": "owner"
}
```

## 🛠️ Commands

| Command                  | Description                     |
| ------------------------ | ------------------------------- |
| `create [name]`          | Create a new stash project      |
| `config <key> [value]`   | Get or set global configuration |
| `watch <file\|dir\|all>` | Add files to track              |
| `store <message>`        | Store current changes           |
| `login`                  | Authenticate with remote server |
| `push`                   | Push changes to remote server   |
| `pull`                   | Pull changes from remote server |
| `help`                   | Show available commands         |

## 🗺️ Roadmap

| Status | Feature                     |
| :----: | --------------------------- |
|   ✅   | Project initialization      |
|   ✅   | File tracking               |
|   ✅   | Store snapshots             |
|   ✅   | Global user config          |
|   🔲   | View diff between snapshots |
|   🔲   | Push to remote server       |
|   🔲   | Pull from remote server     |
|   🔲   | Branching support           |
|   🔲   | Merge conflicts             |

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## 📄 License

MIT License — feel free to use this project however you like.

---

## Author

**Saugat Adhikari**  
GitHub: [@adk-saugat](https://github.com/adk-saugat)
