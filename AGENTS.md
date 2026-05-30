# AGENTS.md

This document outlines the conventions and commands for working with this Go codebase.

## Build, Lint, and Test

### Build

To build the project, run:

```sh
go build ./...
```

### Lint

This project does not currently have a configured linter.

### Test

Run all tests with:

```sh
go test ./...
```

To run a single test, use the `-run` flag with a regex matching the test name. For example, to run `TestDecodeArrayString`:

```sh
go test ./core -run TestDecodeArrayString
```

## Code Style

### Formatting

All Go code should be formatted with `gofmt`.

### Imports

Imports should be organized into two groups: standard library and third-party/project imports.

Example:

```go
import (
    "fmt"
    "net"

    "rediss/core"
)
```

### Types

- Use built-in Go types where possible.
- Define custom `struct` types for complex data structures.

### Naming Conventions

- **Packages**: `lowercase`
- **Files**: `lowercase` with words separated by hyphens (e.g., `file-name.go`).
- **Functions**:
  - Public: `PascalCase`
  - Private: `camelCase`
- **Variables**: `camelCase`

### Error Handling

- Functions that can fail should return an `error` as the last value.
- Check for errors immediately after the function call.
- Use `log.Fatalln` for errors that should stop the application.
- Use `fmt.Sprintf("-%s\r\n", err)` to format errors for a Redis client.

### Comments

No comments. Ever.

### Caveman Mode (Always follow these rules)

- Drop articles (a/an/the).
- Drop filler words (just, really, basically, actually, simply).
- Drop pleasantries (sure, certainly, of course, happy to).
- Drop hedging.
- Use fragments.
- Use short synonyms (e.g., "big" not "extensive", "fix" not "implement a solution for").
- Abbreviate common terms (DB, auth, config, req, res, fn, impl).
- Strip conjunctions.
- Use arrows for causality (X -> Y).
- Use one word when one word is enough.

Technical terms stay exact. Code blocks remain unchanged. Errors are quoted exactly.

**Pattern**: `[thing] [action] [reason]. [next step].`

**Example**:
Not: "Sure! I'd be happy to help you with that. The issue you're experiencing is likely caused by..."
Yes: "Bug in auth middleware. Token expiry check use < not <=. Fix:"

## Project Structure

- `main.go`: Entry point of the application.
- `server/`: TCP server logic.
- `core/`: Core Redis command evaluation and RESP (REdis Serialization Protocol) handling.

## Dependencies

Dependencies are managed using Go modules (`go.mod`).
