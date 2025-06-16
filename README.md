# Desktop Automation MCP Server

A Model Context Protocol (MCP) server that exposes desktop automation capabilities, allowing LLM applications to interact with your desktop environment through standardized tools.

## Overview

This MCP server provides a bridge between LLM applications and desktop automation functionality. It exposes mouse and keyboard automation capabilities through the MCP protocol, enabling AI assistants to:

- Click, right-click, and double-click at specific coordinates
- Move the mouse cursor with smooth or instant movement
- Type text with optional character delays
- Press individual keys or key combinations
- Get current mouse cursor position

## Features

### Mouse Automation
- **Click**: Left click at specified coordinates
- **Right Click**: Right click at specified coordinates
- **Double Click**: Double click at specified coordinates
- **Move Mouse**: Move cursor with optional smooth animation
- **Get Position**: Retrieve current mouse coordinates

### Keyboard Automation
- **Type Text**: Type text at current cursor position
- **Press Key**: Press individual keys or key combinations
- **Delayed Typing**: Type with configurable delays between characters

## Installation

### Prerequisites

- Go 1.23.0 or later
- Task runner (optional, for using Taskfile commands)

```bash
# Install Task (optional)
go install github.com/go-task/task/v3/cmd/task@latest
```

### Build from Source

```bash
# Clone the repository (if not already available)
git clone <repository-url>
cd desktop-automation-mcp

# Download dependencies
go mod download
go mod tidy

# Build the server
go build -o mcp-server ./cmd/mcp-server

# Or using Task
task build
```

## Usage

### Running the Server

The server uses stdio transport for communication:

```bash
# Run directly
./mcp-server

# Or using Task
task run
```

### Integration with LLM Applications

Configure your LLM application (Claude Desktop, etc.) to connect to this MCP server:

```json
{
  "mcpServers": {
    "desktop-automation": {
      "command": "/path/to/mcp-server"
    }
  }
}
```

## Available Tools

### `click`
Click at specified screen coordinates.

**Parameters:**
- `x` (number, required): X coordinate
- `y` (number, required): Y coordinate

### `right_click`
Right click at specified screen coordinates.

**Parameters:**
- `x` (number, required): X coordinate
- `y` (number, required): Y coordinate

### `double_click`
Double click at specified screen coordinates.

**Parameters:**
- `x` (number, required): X coordinate
- `y` (number, required): Y coordinate

### `move_mouse`
Move mouse cursor to specified coordinates.

**Parameters:**
- `x` (number, required): X coordinate
- `y` (number, required): Y coordinate
- `smooth` (boolean, optional): Use smooth movement animation
- `duration` (number, optional): Duration for smooth movement in seconds (default: 1.0)

### `get_mouse_position`
Get current mouse cursor position.

**Parameters:** None

### `type_text`
Type text at current cursor position.

**Parameters:**
- `text` (string, required): Text to type
- `delay` (number, optional): Delay between characters in milliseconds

### `press_key`
Press a key or key combination.

**Parameters:**
- `key` (string, required): Key to press (e.g., 'enter', 'space', 'ctrl')
- `modifiers` (array, optional): Modifier keys (e.g., ['ctrl', 'shift'])

## Architecture

```
desktop-automation-mcp/
├── cmd/
│   └── mcp-server/          # MCP server entry point
│       └── main.go
├── internal/
│   └── automation/          # Desktop automation logic (copied from desktop-automation)
│       ├── keyboard.go
│       └── mouse.go
├── go.mod                   # Go module definition
├── Taskfile.yml            # Task runner configuration
├── .gitignore              # Git ignore rules
└── README.md               # This file
```

## Development

### Task Commands

```bash
# Build the server
task build

# Run the server
task run

# Clean build artifacts
task clean

# Download and tidy dependencies
task deps

# Run tests
task test

# Install to GOPATH/bin
task install
```

### Manual Commands

```bash
# Build
go build -o mcp-server ./cmd/mcp-server

# Run
./mcp-server

# Test
go test ./...

# Clean
go clean
rm -f mcp-server
```

## Dependencies

- **[github.com/mark3labs/mcp-go](https://github.com/mark3labs/mcp-go)**: MCP protocol implementation for Go
- **[github.com/go-vgo/robotgo](https://github.com/go-vgo/robotgo)**: Cross-platform desktop automation library

## Safety Considerations

- **Screen Bounds**: All coordinate inputs are validated against screen dimensions
- **Input Validation**: Negative coordinates and invalid parameters are rejected
- **Error Handling**: Comprehensive error handling with descriptive messages
- **Recovery**: Built-in panic recovery for robust operation

## Platform Support

This server supports the same platforms as robotgo:
- Windows
- macOS
- Linux

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project follows the same license as the parent desktop-automation project.

## Related Projects

- [desktop-automation](../desktop-automation): The original CLI-based desktop automation tool
- [mcp-go](https://github.com/mark3labs/mcp-go): Go implementation of the Model Context Protocol
