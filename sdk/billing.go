// sdk/billing.go
package sdk

import (
	"context"
	"log"
	"strings"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type logPayload struct {
	Level     string            `json:"level"`
	Timestamp string            `json:"timestamp"`
	Event     string            `json:"event"`
	Metadata  map[string]string `json:"metadata"`
}

// BillingSDK is the main struct for tracking billing events.
type BillingSDK struct {
	metrics  Metrics
	backends []Backend
}

func NewBillingSDKWithMetrics(m Metrics) *BillingSDK {
	return &BillingSDK{metrics: m}
}

func NewBillingSDKWithBackends(metrics Metrics, backends ...Backend) *BillingSDK {
	return &BillingSDK{metrics: metrics, backends: backends}
}

func NewBillingSDK() *BillingSDK {
	log.Fatal("NewBillingSDK requires explicit Metrics. Use NewBillingSDKWithMetrics instead.")
	return nil
}

func (b *BillingSDK) TrackEvent(level, eventName string, metadata map[string]string) {
	normalizedLevel := strings.ToUpper(level)
	if normalizedLevel == "" {
		normalizedLevel = "INFO"
	}

	_, span := otel.Tracer("billing-sdk").Start(context.Background(), eventName)
	defer span.End()

	span.SetAttributes(
		attribute.String("event.level", normalizedLevel),
		attribute.String("event.name", eventName),
	)
	for k, v := range metadata {
		span.SetAttributes(attribute.String(k, v))
	}

	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		b.metrics.ObserveDuration(eventName, duration)
	}()

	b.metrics.IncEvent(eventName, normalizedLevel)

	event := BillingEvent{
		Level:     normalizedLevel,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Event:     eventName,
		Metadata:  metadata,
	}

	for _, backend := range b.backends {
		backend.SendEvent(event)
	}
}
