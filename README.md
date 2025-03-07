# VKOG - Lightweight Key-Value Store

VKOG is a simple, lightweight, and extendable key-value store written in Go. It provides an HTTP interface for storing and retrieving data with support for plugins and persistent storage.

## Features

- Simple HTTP API for key-value operations
- Persistent storage using MessagePack serialization
- Memory management with configurable size limits
- Plugin system for extending functionality
- Thread-safe operations
- Lightweight and easy to deploy

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/vkog.git
cd vkog

# Build the application
go build -o vkog
```

## Usage

### Starting the Server

```bash
# Start with default settings (128MB max size, local.vkog file)
./vkog

# Custom max size (256MB) and storage file
./vkog -s 268435456 -f data.vkog
```

Command-line options:
- `-s`: Maximum size in bytes (default: 128MB)
- `-f`: Storage file name (default: local.vkog)

### HTTP API

The server listens on port 3110 by default and exposes the following endpoints:

#### GET /k/{key}

Retrieves the value associated with the specified key.

```bash
curl http://localhost:3110/k/mykey
```

Response:
- 200 OK with the value in the response body if found
- 404 Not Found if the key doesn't exist

#### POST /k/{key}

Creates a new key-value pair. The value is provided in the request body.

```bash
curl -X POST -d "myvalue" http://localhost:3110/k/mykey
```

Response:
- 200 OK with "Ok" in the response body if successful
- 400 Bad Request if no value is provided
- 413 Request Entity Too Large if adding the key-value pair would exceed the maximum size

#### PUT /k/{key}

Updates an existing key-value pair. The value is provided in the request body.

```bash
curl -X PUT -d "newvalue" http://localhost:3110/k/mykey
```

Response:
- 200 OK with "Ok" in the response body if successful
- 400 Bad Request if no value is provided
- 413 Request Entity Too Large if updating the key-value pair would exceed the maximum size

#### DELETE /k/{key}

Deletes a key-value pair.

```bash
curl -X DELETE http://localhost:3110/k/mykey
```

Response:
- 200 OK with "Ok" in the response body if successful
- 404 Not Found if the key doesn't exist

## Plugin System

VKOG includes a plugin system that enables extending the functionality. Currently implemented plugins:

### Conflict Plugin

Prevents conflicts when writing to the same key simultaneously.

### Bytes Plugin

Manages memory usage and tracks the size of stored key-value pairs.

## Development

### Build Commands

- Build: `go build -o vkog`
- Run: `go run vkog.go` or `./vkog`

### Project Structure

- `vkog.go`: Main entry point and server initialization
- `handler.go`: HTTP request handling
- `memorymap.go`: Core key-value store implementation
- `plugin.go`: Plugin interface definitions
- `registry.go`: Plugin registry implementation
- `bytesplugin.go`: Memory management plugin
- `conflictplugin.go`: Conflict resolution plugin
- `helpers.go`: Utility functions for persistence

### Extending with Plugins

To create a new plugin:

1. Implement the Plugin interface and required event handlers
2. Register your plugin in vkog.go

Example plugin implementation:

```go
type MyPlugin struct{}

func (p *MyPlugin) getName() string {
    return "my.plugin"
}

// Implement desired event handlers
func (p *MyPlugin) beforeGet(key string) {
    // Custom logic before retrieving a key
}
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.