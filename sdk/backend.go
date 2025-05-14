package sdk

type BillingEvent struct {
	Level     string
	Timestamp string
	Event     string
	Metadata  map[string]string
}

type Backend interface {
	SendEvent(event BillingEvent)
}
