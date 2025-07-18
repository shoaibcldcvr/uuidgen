# UUID Generator (Go)

A fast, customizable UUID generator with CLI and library support.

## Features
- Generate UUID v1 (time-based) and v4 (random)
- Simple CLI interface
- Easy-to-use Go library
- Extensible for other UUID versions

## Installation

### Prerequisites
- Go 1.20+ installed

### Install the CLI
```bash
go install github.com/shoaibcldcvr/uuidgen/cmd/cli@latest
```

## How to Use
### As a CLI Tool
```bash
# Install
go install github.com/yourusername/uuidgen/cmd/cli@latest

# Generate UUIDs
uuidgen -v=4       # Single UUIDv4
uuidgen -v=1 -n=5  # Five UUIDv1s
```

### As a Library
```go
import "github.com/yourusername/uuidgen/pkg/uuid"

func main() {
    u, _ := uuid.NewV4()
    fmt.Println("Generated UUID:", u.String())
}
```
