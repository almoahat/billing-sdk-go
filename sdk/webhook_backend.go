package sdk

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type WebhookBackend struct {
	URL        string
	HTTPClient *http.Client
	Headers    map[string]string
}

// NewWebhookBackend creates a new instance with optional headers
func NewWebhookBackend(url string, headers map[string]string) *WebhookBackend {
	return &WebhookBackend{
		URL: url,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		Headers: headers,
	}
}

// SendEvent serializes and POSTs the billing event
func (w *WebhookBackend) SendEvent(event BillingEvent) {
	body, err := json.Marshal(event)
	if err != nil {
		log.Printf("WebhookBackend: failed to marshal event: %v", err)
		return
	}

	req, err := http.NewRequest("POST", w.URL, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("WebhookBackend: failed to create request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range w.Headers {
		req.Header.Set(k, v)
	}

	resp, err := w.HTTPClient.Do(req)
	if err != nil {
		log.Printf("WebhookBackend: HTTP POST failed: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Printf("WebhookBackend: unexpected response code: %d", resp.StatusCode)
	}
}
