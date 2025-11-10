# Version Info for Go Applications

A lightweight Go package to embed build-time version information into your binaries — including version, Git commit, build timestamp, and Go version.

> 🏷 Perfect for CLIs, microservices, health checks, and observability.

---

## 🚀 Features

- Embed **version**, **Git revision**, and **build date** at compile time.
- No runtime dependencies — uses Go's `-ldflags`.
- Compatible with **CI/CD pipelines** (GitHub Actions, GitLab CI, etc.).
- Human-readable output via `version.Info()`.
- Falls back to `dev` mode if built without flags.

---

## 📦 Usage

### 1. Import the package

```bash
go get https://github.com/ayden1st/go-version
```

### 2. Use in your main function
```go
package main

import (
    "fmt"
    "github.com/ayden1st/go-version"
)

func main() {
    fmt.Println(version.Info("myapp"))
}
```
Example output:

`myapp, version: 1.2.3, revision: a1b2c3d, build date: 2025-04-05T10:20:30, go1.24.0`

---

## 🛠 Build with Version Info
Use `-ldflags` to inject build-time variables:

```bash
# Example values (usually set in CI)
VERSION=$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2>/dev/null | sed 's/^v//')
COMMIT=$(git rev-parse --short HEAD)
DATE=$(date '+%Y-%m-%dT%H:%M:%S')

# Build the binary
go build -ldflags "
    -X 'github.com/ayden1st/go-version.Version=${VERSION}' \
    -X 'github.com/ayden1st/go-version.Revision=${COMMIT}' \
    -X 'github.com/ayden1st/go-version.BuildDate=${DATE}'
" -o myapp main.go
```

---

## 🧩 Example Output
`myapp, version: 1.5.0, revision: abc1234, build date: 2025-04-05T10:20:30, go1.24.0`

Or when built with `go run` (no flags):
`myapp, version: dev, revision: none, build date: Mon Apr  5 10:20:30 UTC 2006, go1.24.0`

---

### 🔍 How It Works
The package uses Go’s linker flags (`-ldflags -X`) to inject values into variables at build time:

```go
var (
    Version   = "dev"        // Overridden at build
    Revision  = "none"       // Overridden at build
    BuildDate = ""           // Overridden at build
    GoVersion = runtime.Version()
)
```
If no `-ldflags` are provided, defaults are used (`dev`, `none`, current time).
If built in CI or release pipeline — actual values are injected.

---

## 🧪 Testing
You can test the output without a full build:

```bash
go run -ldflags "
    -X 'github.com/ayden1st/go-version.Version=1.0.0' \
    -X 'github.com/ayden1st/go-version.Revision=abc123' \
    -X 'github.com/ayden1st/go-version.BuildDate=2025-01-01T00:00:00'
" main.go
```

---

## 📄 License
MIT — feel free to use in any project.

---
