package main

import (
	"log"
	"net/http"
	"os"

	"github.com/almoahat/billing-sdk-go/sdk"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	sdk.InitTracer()

	// Create or append to a file for logs
	f, err := os.OpenFile("billing.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer f.Close()

	// Configure logging
	sdk.ConfigureLogger(f, false)
	log.SetOutput(f)

	// Create default Prometheus metrics implementation
	metrics := sdk.NewPromMetrics(prometheus.DefaultRegisterer)

	// Start /metrics HTTP endpoint in background
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("Serving metrics at http://localhost:2112/metrics")
		log.Fatal(http.ListenAndServe(":2112", nil))
	}()

	// Set up backends
	console := &sdk.ConsoleBackend{}
	webhook := sdk.NewWebhookBackend("http://localhost:9000/hook", map[string]string{
		"Authorization": "Bearer example-token",
	})

	billing := sdk.NewBillingSDKWithBackends(metrics, console, webhook)

	billing.TrackEvent("INFO", "api_call", map[string]string{
		"user_id": "1234",
		"plan":    "pro",
	})

	billing.TrackEvent("ERROR", "billing_failure", map[string]string{
		"user_id": "5678",
		"reason":  "payment_declined",
	})

	// Keep the app running so metrics endpoint stays alive
	select {}
}
