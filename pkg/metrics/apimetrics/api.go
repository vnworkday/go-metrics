package apimetrics

import (
	"github.com/vnworkday/go-metrics/pkg/metrics"
	"github.com/vnworkday/go-metrics/pkg/tags"
)

const (
	APIRequestHistogramName = "api_request_histogram"
	APIRequestHistogramDesc = "API Request Histogram"
	APIRequestCounterName   = "api_request_counter"
	APIRequestCounterDesc   = "API Request Counter"
)

type Client interface {
	GetLatencyHistogram() metrics.Histogram
	GetRequestCounter() metrics.Counter
}

type client struct {
	client     metrics.MetricProvider
	latency    metrics.Histogram
	reqCounter metrics.Counter
	tags       []tags.Tag
}

func (m client) GetRequestCounter() metrics.Counter {
	return m.reqCounter
}

func (m client) GetLatencyHistogram() metrics.Histogram {
	return m.latency
}

func New(apiName string, provider metrics.MetricProvider, options ...MetricOption) (Client, error) {
	if apiName == "" {
		return client{}, metrics.ErrAPINameNotSpecified
	}

	if provider == nil {
		return client{}, metrics.ErrMetricClientNotSpecified
	}

	m := client{
		client: provider,
	}

	for _, option := range options {
		option(&m)
	}

	m.tags = append(m.tags, tags.APIName(apiName))

	latency, err := provider.GetHistogram(
		APIRequestHistogramName,
		metrics.NewInstrumentOptions().
			WithTags(m.tags...).
			WithDesc(APIRequestHistogramDesc),
	)

	if err != nil {
		return client{}, err
	}

	reqCounter, err := provider.GetCounter(
		APIRequestCounterName,
		metrics.NewInstrumentOptions().
			WithTags(m.tags...).
			WithDesc(APIRequestCounterDesc),
	)

	if err != nil {
		return client{}, err
	}

	m.latency = latency
	m.reqCounter = reqCounter

	return m, nil
}

type MetricOption func(*client)

func WithMetricTags(tags ...tags.Tag) MetricOption {
	return func(m *client) {
		m.tags = append(m.tags, tags...)
	}
}
