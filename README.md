# Billing SDK for Billing Event Tracking with Observability Hooks

**Work in Progress**  
A lightweight SDK in Go to emit billing events with built-in observability support — including structured logs, Prometheus metrics, OpenTelemetry spans, and pluggable backends like webhooks and console logging.

---

## Features

- `TrackEvent(level, eventName, metadata)` to emit structured billing events
- JSON logs with user-defined metadata and log levels
- Prometheus metrics:
  - Counter for total events
  - Histogram for event duration
- OpenTelemetry tracing with span attributes
- Pluggable backend architecture:
  - Console logging
  - Webhook POST support
- CLI mock webhook server for local testing
- Unit-testable design with mockable metrics and tracing

---

## Tech Stack

- Language: Go
- Logging: Standard `log` package + JSON output
- Metrics: Prometheus + Histogram/CounterVec
- Tracing: OpenTelemetry + stdout exporter
- Extensibility: Backend plugins (e.g., Webhook, Console)

---

## Getting Started

### 1. Clone & Run Example

```bash
git clone https://github.com/your-username/billing-sdk-go.git
cd billing-sdk-go

go mod tidy
go run example_service/main.go
```

### 2. Run Mock Webhook Server (Optional)

```bash
go run cmd/webhook_server/main.go
```

Then check your terminal for incoming JSON events.

---

## Example Webhook Output

```json
{
  "level": "INFO",
  "timestamp": "2025-05-14T14:48:48Z",
  "event": "api_call",
  "metadata": {
    "user_id": "1234",
    "plan": "pro"
  }
}
```

---

## Project Structure

```
billing-sdk-go/
├── sdk/                # SDK core: tracing, metrics, backends
├── example_service/    # Example usage of the SDK
├── cmd/webhook_server/ # Local CLI webhook test server
├── go.mod / go.sum     # Module dependencies
└── README.md
```
