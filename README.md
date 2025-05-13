# Billing SDK with Observability Hooks

**Work in Progress**  
A lightweight SDK in Go to emit billing events with built-in observability support — including structured logs, log levels, and extensible metadata.

---

## Features

- `TrackEvent(level, eventName, metadata)` logs structured billing events
- JSON-formatted logs with metadata
- Support for log levels: `INFO`, `ERROR`, etc.
- Simple, self-contained SDK design
- Plug-and-play with existing services
- Unit-testable with injected log output
- Example consumer app included (`example_service/`)

---

## Tech Stack

- Go
- Logging: Standard `log` package + JSON formatting
- File output supported
- Designed for future Prometheus + OpenTelemetry support

---

## Getting Started

### Clone & Run

```bash
git clone https://github.com/your-username/billing-sdk-go.git
cd billing-sdk-go

go mod tidy
go run example_service/main.go
