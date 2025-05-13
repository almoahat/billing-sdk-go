package sdk

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestTrackEvent(t *testing.T) {
	sdk := NewBillingSDK()

	var buf bytes.Buffer
	ConfigureLogger(&buf, true) // Disable timestamps and set output

	metadata := map[string]string{
		"user_id": "42",
		"plan":    "premium",
	}

	sdk.TrackEvent("INFO", "subscription_billed", metadata)

	var output map[string]interface{}
	err := json.Unmarshal(bytes.TrimSpace(buf.Bytes()), &output)
	if err != nil {
		t.Fatalf("Failed to parse log output as JSON: %v", err)
	}

	if output["event"] != "subscription_billed" {
		t.Errorf("Expected event 'subscription_billed', got: %v", output["event"])
	}

	meta, ok := output["metadata"].(map[string]interface{})
	if !ok {
		t.Fatalf("Expected metadata to be a map, got: %T", output["metadata"])
	}

	if meta["user_id"] != "42" || meta["plan"] != "premium" {
		t.Errorf("Expected metadata user_id=42 and plan=premium, got: %v", meta)
	}
}
