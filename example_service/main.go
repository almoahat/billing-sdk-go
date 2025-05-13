package main

import (
	"log"
	"os"

	"github.com/almoahat/billing-sdk-go/sdk"
)

func main() {
	// Create or append to a file
	f, err := os.OpenFile("billing.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer f.Close()

	sdk.ConfigureLogger(f, false) // log to file, keep timestamps

	// Redirect log output
	log.SetOutput(f)

	billing := sdk.NewBillingSDK()

	billing.TrackEvent("INFO", "api_call", map[string]string{
		"user_id": "1234",
		"plan":    "pro",
	})
	billing.TrackEvent("ERROR", "billing_failure", map[string]string{
		"user_id": "5678",
		"reason":  "payment_declined",
	})
}
