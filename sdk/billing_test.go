package sdk

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func setupLogCapture() *bytes.Buffer {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
	return &buf
}

func TestBillingEventCounter(t *testing.T) {
	reg := prometheus.NewRegistry()
	metrics := NewPromMetrics(reg)

	sdk := NewBillingSDKWithMetrics(metrics)
	sdk.TrackEvent("INFO", "api_call", map[string]string{
		"user_id": "42",
		"plan":    "test",
	})

	value := testutil.ToFloat64(metrics.eventCounter.WithLabelValues("api_call", "INFO"))
	if value != 1 {
		t.Errorf("Expected billing_events_total to be 1, got: %v", value)
	}
}

func TestBillingEventDurationHistogram(t *testing.T) {
	reg := prometheus.NewRegistry()
	metrics := NewPromMetrics(reg)

	sdk := NewBillingSDKWithMetrics(metrics)
	sdk.TrackEvent("INFO", "billing_latency_test", map[string]string{})

	count := testutil.CollectAndCount(metrics.eventHistogram)
	if count == 0 {
		t.Error("Expected billing_event_duration_seconds to record at least one observation")
	}
}

func TestTrackEventLogsJSON(t *testing.T) {
	buf := setupLogCapture()

	sdk := NewBillingSDKWithMetrics(&mockMetrics{})
	sdk.TrackEvent("INFO", "log_test", map[string]string{"key": "value"})

	// Debug: print raw logs
	logOutput := buf.String()
	t.Logf("Captured log output:\n%s", logOutput)

	lines := bytes.Split(buf.Bytes(), []byte("\n"))
	if len(lines) < 2 {
		t.Fatalf("Expected at least one log line, got: %d", len(lines))
	}

	// Grab last non-empty line
	last := bytes.TrimSpace(lines[len(lines)-2])
	t.Logf("Parsing log line: %s", last)

	var output map[string]interface{}
	err := json.Unmarshal(last, &output)
	if err != nil {
		t.Fatalf("Failed to parse JSON log: %v\nRaw: %s", err, last)
	}

	if output["event"] != "log_test" {
		t.Errorf("Expected event to be 'log_test', got: %v", output["event"])
	}

	meta, ok := output["metadata"].(map[string]interface{})
	if !ok || meta["key"] != "value" {
		t.Errorf("Expected metadata key=value, got: %v", meta)
	}
}

type mockMetrics struct{}

func (m *mockMetrics) IncEvent(event, level string)                  {}
func (m *mockMetrics) ObserveDuration(event string, seconds float64) {}

func TestTrackEvent_EmitsSpanWithAttributes(t *testing.T) {
	// Setup in-memory span recorder
	sr := tracetest.NewSpanRecorder()
	tp := sdktrace.NewTracerProvider()
	tp.RegisterSpanProcessor(sr)
	otel.SetTracerProvider(tp)

	// Use dummy metrics
	metrics := &mockMetrics{}
	sdk := NewBillingSDKWithMetrics(metrics)

	sdk.TrackEvent("INFO", "span_test", map[string]string{
		"user_id": "999",
		"plan":    "enterprise",
	})

	// Get spans from the recorder
	spans := sr.Ended()
	if len(spans) != 1 {
		t.Fatalf("Expected 1 span, got: %d", len(spans))
	}

	span := spans[0]
	if span.Name() != "span_test" {
		t.Errorf("Expected span name 'span_test', got: %s", span.Name())
	}

	attrs := span.Attributes()

	// Helper to verify presence of expected attributes
	expectAttr := func(key, val string) {
		found := false
		for _, a := range attrs {
			if a.Key == attribute.Key(key) && a.Value.AsString() == val {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected span attribute %s=%s", key, val)
		}
	}

	expectAttr("event.level", "INFO")
	expectAttr("event.name", "span_test")
	expectAttr("user_id", "999")
	expectAttr("plan", "enterprise")
}
