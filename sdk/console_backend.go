package sdk

import (
	"encoding/json"
	"log"
)

type ConsoleBackend struct{}

func (c *ConsoleBackend) SendEvent(event BillingEvent) {
	data, err := json.Marshal(event)
	if err != nil {
		log.Printf(`{"level":"ERROR","message":"failed to encode event","error":"%v"}`, err)
		return
	}
	log.Println(string(data))
}
