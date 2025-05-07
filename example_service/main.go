package main

import (
    "billing-sdk-demo/sdk"
)

func main() {
    billing := sdk.NewBillingSDK()

    metadata := map[string]string{
        "user_id": "1234",
        "plan":    "pro",
    }

    billing.TrackEvent("api_call", metadata)
}
