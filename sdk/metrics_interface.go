package sdk

type Metrics interface {
	IncEvent(event, level string)
	ObserveDuration(event string, seconds float64)
}
