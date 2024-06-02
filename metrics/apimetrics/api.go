package apimetrics

import (
	"github.com/vnworkday/go-metrics/metrics"
	"time"
)

const (
	ApiRequestHistogramName = "api_request_histogram"
	ApiRequestHistogramDesc = "API Request Histogram"
	ApiRequestCounterName   = "api_request_counter"
	ApiRequestCounterDesc   = "API Request Counter"
)

type Metrics interface {
	metrics.Client
	GetRequestCounter() metrics.Counter
	GetLatencyHistogram() metrics.Histogram
	UtcNow() time.Time
}

type Metric struct {
	metrics.Client
	name       string
	latency    metrics.Histogram
	reqCounter metrics.Counter
	tags       []metrics.Tag
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

	m.tags = append(m.tags, metrics.ApiTag(name))

	latency, err := client.GetHistogram(
		ApiRequestHistogramName,
		metrics.NewInstrumentOptions().
			WithTags(m.tags...).
			WithDesc(ApiRequestHistogramDesc),
	)

	if err != nil {
		return Metric{}, err
	}

	reqCounter, err := client.GetCounter(
		ApiRequestCounterName,
		metrics.NewInstrumentOptions().
			WithTags(m.tags...).
			WithDesc(ApiRequestCounterDesc),
	)

	if err != nil {
		return Metric{}, err
	}

	m.latency = latency
	m.reqCounter = reqCounter

	return m, nil
}

type MetricOption func(*Metric)
