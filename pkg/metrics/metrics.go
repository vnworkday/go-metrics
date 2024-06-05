package metrics

import (
	"context"

	"github.com/vnworkday/go-metrics/pkg/tags"
)

type Meter = func() int
type Unregister = func() error

type MetricProvider interface {
	RegisterMeter(name string, meter Meter, options ...InstrumentOptions) (Unregister, error)
	GetCounter(name string, options ...InstrumentOptions) (Counter, error)
	GetHistogram(name string, options ...InstrumentOptions) (Histogram, error)
	GetUpDownCounter(name string, options ...InstrumentOptions) (UpDownCounter, error)
	GetGauge(name string, options ...InstrumentOptions) (Gauge, error)
}

// Counter is a metric that represents a single numerical value that only ever goes up.
// It is typically used to measure things like the number of requests, the total number of errors, or the rate of incoming events.
type Counter interface {
	Add(ctx context.Context, value uint, tags ...tags.Tag)
}

// Histogram is a metric that represents the distribution of a set of values.
// It is typically used to measure things like request latency, response sizes, allowing you to analyze percentiles and outliers.
type Histogram interface {
	Record(ctx context.Context, value int, tags ...tags.Tag)
}

// UpDownCounter is a metric that represents a cumulative value that can both increase and decrease.
// It is typically used to measure things like the number of items in a queue, the number of active threads, or the number of open connections.
type UpDownCounter interface {
	Add(ctx context.Context, value uint, tags ...tags.Tag)
}

// Gauge is a metric that represents an instantaneous value at a point in time.
// It is typically used to measure things like the current CPU usage, or the amount of free memory.
type Gauge interface {
	Record(ctx context.Context, value int, tags ...tags.Tag)
}
