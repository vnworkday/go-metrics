package metrics

import "time"

type Metrics interface {
	Client
	GetLatencyHistogram() Histogram
	UtcNow() time.Time
}
