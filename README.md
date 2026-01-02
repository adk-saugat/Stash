# Stash

A lightweight version control system built in Go. Track files, store snapshots, and view your project history — all from the command line.

## Features

- **Simple CLI** — Intuitive commands inspired by Git
- **File Tracking** — Watch individual files or entire projects
- **SHA-256 Hashing** — Detect changes efficiently without storing duplicates
- **Store History** — View a tree-style log of all your snapshots
- **Global Config** — Set your identity once, use it everywhere

## Installation

```bash
# Clone the repository
git clone https://github.com/adk-saugat/stash.git
cd stash

# Build the binary
make build

# (Optional) Move to PATH
sudo mv stash /usr/local/bin/
```

## Quick Start

```bash
# 1. Configure your identity (one-time setup)
stash config <username> <email>

# 2. Initialize a new project
cd your-project
stash create

# 3. Add files to track
stash watch main.py          # Single file
stash watch all              # Everything

# 4. Store a snapshot
stash store "initial commit"

# 5. View history
stash log
```

## Commands

| Command | Description |
|---------|-------------|
| `stash config <username> <email>` | Set global username and email |
| `stash create [project-name]` | Initialize a new stash repository |
| `stash watch <file\|all>` | Add files to track |
| `stash store <message>` | Create a new snapshot |
| `stash log` | Show store history |
| `stash help` | List all available commands |

## How It Works

### Project Structure

When you run `stash create`, a `.stash` folder is created:

```
your-project/
├── .stash/
│   ├── projectConfig.json    # Project metadata & tracked files
│   └── stores/               # Snapshot history
│       ├── abc123...json
│       └── def456...json
└── your-files...
```

### Change Detection

Stash uses SHA-256 hashes to detect file changes. When you run `stash store`:

1. Each tracked file is hashed
2. Hashes are compared with the latest store
3. If nothing changed, the store is blocked
4. If changes exist, a new snapshot is created

### Store Format

Each store is a JSON file containing:

```json
{
    "store_id": "680e4cd0-6b68-4244-b2de-9bfb25d80045",
    "project_id": "f779a6dc-f4e4-474c-91a3-329a3d917f05",
    "author": "you@email.com",
    "message": "your commit message",
    "date": "2025-01-02T10:30:00Z",
    "files": [
        {
            "path": "main.py",
            "hash": "a1b2c3...",
            "content": "base64-encoded-content"
        }
    ]
}
```

## Example Workflow

```bash
$ stash create myproject
Repository initialized.
Project configuration created.

$ stash watch all
Project configuration loaded.
Files discovered.
Watch list updated.

$ stash store "initial commit"
Project configuration loaded.
User configuration loaded.
Tracked files processed.
Store created.

$ stash log

Store History (1 stores)
─────────────────────────

● 680e4cd0 - initial commit (just now) <you@email.com>
```

## Roadmap

- [ ] `stash login` — Authenticate using config and password
- [ ] `stash logout` — Log out (clear local session)
- [ ] `stash join <project-id>` — Join an existing project
- [ ] `stash share` — Upload local stores to server
- [ ] `stash fetch` — Download remote stores from server

## Requirements

- Go 1.21+

## Author

**Saugat Adhikari** — [@adk-saugat](https://github.com/adk-saugat)
