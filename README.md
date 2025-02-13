# Dupignore

**Dupignore** is a simple CLI tool to remove duplicate entries from `.gitignore`.

---
## Installation

### Using Go

```bash
go install github.com/immohika/dupignore@latest
```

### Pre-built Binaries

Check the [Releases page](https://github.com/immohika/dupignore/releases).

#### Linux/macOS

```bash
# Download the binary
curl -L https://github.com/immohika/dupignore/releases/latest/download/dupignore-<OS>-<ARCH>.tar.gz -o dupignore.tar.gz

# Extract
tar -xzf dupignore.tar.gz

# Install system wide
sudo install -Dm755 dupignore* /usr/bin/dupignore

# Or install local
install -Dm755 dupignore*  ~/.local/bin/dupignore
```

#### Windows

1. Download the `dupignore-windows-*.zip` file from
   the [Releases page](https://github.com/dupignore/dupignore/releases).
2. Extract the `.zip` file.
3. Add the extracted `dupignore.exe` to your system's `PATH`.

---

## Usage

### Basic Command

Process a `.gitignore` file in the current directory:

```bash
dupignore dedup

# You can also specify the file
dupignore dedup path/to/.gitignore
```

### Override Configuration

```bash
dupignore dedup --keepcomments=false --keepnewlines=true
```

### Manage Configuration

```bash
# View all config values
dupignore config

# View a specific config value
dupignore config keepcomments

# Update a config value
dupignore config keepcomments false
```

---

## Configuration

The global configuration file is stored in the XDG-compliant location:

- **Linux/macOS**: `~/.config/dupignore/config.toml`
- **Windows**: `%AppData%\dupignore\config.toml`

### Default Configuration

```toml
keep_comments = true
keep_new_lines = false
```
