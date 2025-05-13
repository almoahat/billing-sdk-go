// sdk/billing.go
package sdk

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

type logPayload struct {
	Level     string            `json:"level"`
	Timestamp string            `json:"timestamp"`
	Event     string            `json:"event"`
	Metadata  map[string]string `json:"metadata"`
}

// BillingSDK is the main struct for tracking billing events.
type BillingSDK struct{}

// NewBillingSDK initializes and returns a new BillingSDK instance.
func NewBillingSDK() *BillingSDK {
	return &BillingSDK{}
}

// TrackEvent logs a billing event with the given name and metadata.
func (b *BillingSDK) TrackEvent(level, eventName string, metadata map[string]string) {
	normalizedLevel := strings.ToUpper(level)
	if normalizedLevel == "" {
		normalizedLevel = "INFO"
	}

	payload := logPayload{
		Level:     normalizedLevel,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Event:     eventName,
		Metadata:  metadata,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf(`{"level":"ERROR","message":"Failed to marshal log payload","error":"%v"}`, err)
		return
	}

	log.Println(string(jsonData))
}
