package apimetrics

import (
	"github.com/vnworkday/go-metrics/pkg/metrics"
	"github.com/vnworkday/go-metrics/pkg/tags"
	"time"
)

const (
	APIRequestHistogramName = "api_request_histogram"
	APIRequestHistogramDesc = "API Request Histogram"
	APIRequestCounterName   = "api_request_counter"
	APIRequestCounterDesc   = "API Request Counter"
)

type APIMetrics interface {
	metrics.Metrics
	GetRequestCounter() metrics.Counter
}

type Metric struct {
	metrics.Client
	name       string
	latency    metrics.Histogram
	reqCounter metrics.Counter
	tags       []tags.Tag
}

func (m Metric) GetRequestCounter() metrics.Counter {
	return m.reqCounter
}

func (m Metric) GetLatencyHistogram() metrics.Histogram {
	return m.latency
}

func (m Metric) UtcNow() time.Time {
	return time.Now().UTC()
}

func New(name string, client metrics.Client, options ...MetricOption) (Metric, error) {
	m := Metric{
		Client: client,
		name:   name,
	}

	for _, option := range options {
		option(&m)
	}

	m.tags = append(m.tags, tags.APIName(name))

	latency, err := client.GetHistogram(
		APIRequestHistogramName,
		metrics.NewInstrumentOptions().
			WithTags(m.tags...).
			WithDesc(APIRequestHistogramDesc),
	)

	if err != nil {
		return Metric{}, err
	}

	reqCounter, err := client.GetCounter(
		APIRequestCounterName,
		metrics.NewInstrumentOptions().
			WithTags(m.tags...).
			WithDesc(APIRequestCounterDesc),
	)

	if err != nil {
		return Metric{}, err
	}

	m.latency = latency
	m.reqCounter = reqCounter

	return m, nil
}

type MetricOption func(*Metric)

func WithMetricTags(tags ...tags.Tag) MetricOption {
	return func(m *Metric) {
		m.tags = append(m.tags, tags...)
	}
}
