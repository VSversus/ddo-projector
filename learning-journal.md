# Learning Journal

Simple checklist to track what you learned while building Go and "Dungeon and Dragons Online" projects.

Priority:
- P1 = Core now
- P2 = Important next
- P3 = Later/advanced

## Golang

- [x] [P1] Variables and constants
- [ ] [P1] Data types (int, string, slice, map, struct)
- [ ] [P1] Functions and return values
- [ ] [P1] Error handling (error interface, wrapping)
- [ ] [P1] Structs and methods
- [ ] [P1] Interfaces
- [ ] [P1] Packages and imports
- [ ] [P1] Goroutines and concurrency
- [ ] [P1] Channels
- [ ] [P1] JSON marshaling/unmarshaling
- [ ] [P1] HTTP client basics
- [ ] [P1] File I/O
- [ ] [P1] Testing (table-driven tests)
- [ ] [P1] Defer and panic recovery
- [ ] [P1] Context package (cancellation, timeouts)
- [ ] [P2] Reflection basics
- [ ] [P2] Generics (Go 1.18+)
- [ ] [P1] Race detector
- [ ] [P3] Profiling (CPU, memory)
- [ ] [P2] Standard library exploration

## Databases

- [ ] [P1] Relational databases (SQL, database/sql package)
- [ ] [P1] SQL basics (SELECT, INSERT, UPDATE, DELETE)
- [ ] [P1] Prepared statements
- [ ] [P1] Transactions
- [ ] [P2] Connection pooling
- [ ] [P3] Document databases (MongoDB, etc.)
- [ ] [P2] JSON storage and retrieval
- [ ] [P2] Query optimization basics

## Database Tooling

- [ ] [P1] Database migrations (golang-migrate, goose)
- [ ] [P2] sqlx basics
- [ ] [P2] sqlc basics
- [ ] [P2] GORM basics
- [ ] [P3] Ent basics

## HTTP Server (Backend)

- [ ] [P1] net/http server basics
- [ ] [P1] Routing (stdlib 1.22+, chi, or gin)
- [ ] [P1] Handlers (request parsing, response writing)
- [ ] [P1] Middleware basics
- [ ] [P1] Timeouts on server and handlers
- [ ] [P2] Graceful shutdown (SIGTERM, draining)

## REST/API Design

- [ ] [P1] HTTP status codes
- [ ] [P1] Request validation
- [ ] [P2] Pagination patterns
- [ ] [P2] API versioning
- [ ] [P1] Error response format consistency
- [ ] [P2] OpenAPI/Swagger basics

## gRPC and Protocol Buffers

- [ ] [P3] Protocol Buffers basics
- [ ] [P3] Defining .proto contracts
- [ ] [P3] Generating Go code from proto files
- [ ] [P3] Unary RPC basics
- [ ] [P3] Streaming RPC basics

## Caching and Messaging

- [ ] [P2] Redis basics
- [ ] [P2] Caching patterns (read-through, write-through)
- [ ] [P2] Cache invalidation basics
- [ ] [P2] Message queues overview (Kafka, RabbitMQ, NATS)
- [ ] [P2] Producer/consumer basics
- [ ] [P2] Idempotency and retry handling

## Go CLI Tools

- [ ] [P1] go mod (init, tidy, add, download)
- [ ] [P1] go run
- [ ] [P1] go build
- [ ] [P1] go test (running tests, coverage)
- [ ] [P1] go fmt (code formatting)
- [ ] [P1] go vet (static analysis)
- [ ] [P2] golangci-lint basics
- [ ] [P1] go install
- [ ] [P1] go get (fetching dependencies)
- [ ] [P1] Go package patterns for commands (., ./..., ./cmd/...)

## Project Structure

- [ ] [P1] Module organization (go.mod, go.sum)
- [ ] [P1] Package naming conventions
- [ ] [P1] Organizing code into packages
- [ ] [P1] Internal packages and visibility
- [ ] [P2] Project layout patterns
- [ ] [P2] Executable vs library projects
- [ ] [P1] Dependency management best practices

## Debugging

- [ ] [P1] Using fmt.Println for debugging
- [ ] [P1] Using log package
- [ ] [P2] Debugging with Delve debugger
- [ ] [P2] Setting breakpoints
- [ ] [P2] Stepping through code
- [ ] [P2] Inspecting variables
- [ ] [P1] Understanding panic stack traces

## Logging

- [ ] [P1] Standard log package (log, log/slog basics)
- [ ] [P1] Log levels (info, warn, error, debug)
- [ ] [P1] Structured logging with log/slog
- [ ] [P1] Log output formatting
- [ ] [P3] Third-party logging libraries (zerolog/zap)

## Observability

- [ ] [P2] Metrics basics (Prometheus)
- [ ] [P2] Tracing basics (OpenTelemetry)
- [ ] [P2] Correlation IDs across requests
- [ ] [P1] Health checks and readiness probes
- [ ] [P2] Alerting basics

## Common Patterns

- [ ] [P2] Dependency injection
- [ ] [P1] Configuration management
- [ ] [P1] Stdlib config parsing (os.Getenv, flag)
- [ ] [P2] Config loading libraries (viper, envconfig)
- [ ] [P1] Error wrapping and custom errors
- [ ] [P2] Functional options pattern
- [ ] [P3] Builder pattern
- [ ] [P3] Factory pattern

## Security Basics

- [ ] [P1] Storing secrets securely (env vars, .env files)
- [ ] [P1] API key and token handling
- [ ] [P1] Input validation
- [ ] [P1] SQL injection prevention
- [ ] [P2] HTTPS and certificates
- [ ] [P1] Authentication vs authorization
- [ ] [P2] Avoiding common vulnerabilities

## Git

- [ ] [P1] Basic workflow (init, add, commit, push)
- [ ] [P1] Branching and merging
- [ ] [P2] Rebasing
- [ ] [P2] Stashing changes
- [ ] [P1] Viewing history (log, diff)
- [ ] [P2] Reverting and resetting commits
- [ ] [P2] Tagging and releases
- [ ] [P2] Collaboration (pull requests, code review)

## Docker

- [ ] [P1] Images and containers
- [ ] [P1] Dockerfile basics
- [ ] [P1] Building images
- [ ] [P1] Running containers
- [ ] [P2] Multi-stage Docker builds
- [ ] [P2] Volumes and mounts
- [ ] [P2] Networking between containers
- [ ] [P2] Docker Compose
- [ ] [P2] Push to registry
- [ ] [P1] Environment variables in containers

## CI/CD

- [ ] [P2] CI pipeline basics (build, test, lint)
- [ ] [P2] GitHub Actions basics
- [ ] [P2] Running tests in CI
- [ ] [P2] Docker image build in CI
- [ ] [P2] Basic deployment workflow

## AI Skills

- [ ] [P1] Prompting for coding tasks
- [ ] [P1] Writing clear problem statements and acceptance criteria
- [ ] [P1] Verifying AI output with tests and docs
- [ ] [P1] Detecting hallucinations and unsafe assumptions
- [ ] [P1] Fast review workflow (ask AI, run code, validate, iterate)
- [ ] [P1] Privacy basics (never share secrets or tokens)
- [ ] [P2] Using AI for debugging strategy, not just code fixes
- [ ] [P2] Refactoring with AI while preserving behavior
- [ ] [P2] AI-assisted test generation and edge-case discovery
- [ ] [P2] Building small internal tools with AI pair-programming
- [ ] [P2] Cost and latency awareness for model/API usage
- [ ] [P2] Reusable prompt templates for frequent tasks
- [ ] [P3] RAG fundamentals for project knowledge lookup
- [ ] [P3] Embeddings and vector search basics
- [ ] [P3] Function calling and tool-using agents
- [ ] [P3] Evaluation pipelines for AI outputs
- [ ] [P3] Prompt injection and model security hardening
- [ ] [P3] Fine-tuning vs retrieval trade-offs

# Notes - Programming
- The short variable declaration := is a statement, and statements are only allowed inside function bodies in Go.
- A function can call itself recursively and return that call's result.
- In Go, write boolean checks directly. Avoid == true and == false when the value is already a boolean.

# Notes - VsCode
- Ctrl+P quick open