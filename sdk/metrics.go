package sdk

import (
	"github.com/prometheus/client_golang/prometheus"
)

type PromMetrics struct {
	eventCounter   *prometheus.CounterVec
	eventHistogram *prometheus.HistogramVec
}

func NewPromMetrics(reg prometheus.Registerer) *PromMetrics {
	m := &PromMetrics{
		eventCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "billing_events_total",
				Help: "Total number of billing events tracked.",
			},
			[]string{"event", "level"},
		),
		eventHistogram: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "billing_event_duration_seconds",
				Help:    "Duration of TrackEvent calls in seconds.",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"event"},
		),
	}
	reg.MustRegister(m.eventCounter)
	reg.MustRegister(m.eventHistogram)
	return m
}

func (m *PromMetrics) IncEvent(event, level string) {
	m.eventCounter.WithLabelValues(event, level).Inc()
}

func (m *PromMetrics) ObserveDuration(event string, seconds float64) {
	m.eventHistogram.WithLabelValues(event).Observe(seconds)
}
