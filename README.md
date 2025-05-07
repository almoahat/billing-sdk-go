# Billing SDK with Observability Hooks

🚧 **Work in Progress**

This SDK provides a simple interface for emitting billing events while integrating observability primitives such as logging, metrics, and tracing.

## ✨ Features

- `TrackEvent(eventName string, metadata map[string]string)`
- Logs billing events
- Emits metrics via Prometheus
- Tracing hooks with OpenTelemetry (coming soon)

## 📦 Tech Stack

- Go
- Log: Standard Go log
- Metrics: Prometheus client
- Tracing: OpenTelemetry (planned)

## 📁 Structure

- `sdk/`: SDK implementation
- `example_service/`: Minimal demo using the SDK

## 🚀 Getting Started

```bash
go run example_service/main.go
```

## 🛠️ TODO

- Add configurable output sinks
- Support for async processing
- Full tracing support

