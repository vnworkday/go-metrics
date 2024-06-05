package queuemetrics

import (
	"github.com/vnworkday/go-metrics/pkg/metrics"
	"github.com/vnworkday/go-metrics/pkg/tags"
)

const (
	QueueMessageLatencyHistogramName = "queue_message_latency_histogram"
	QueueMessageLatencyHistogramDesc = "Queue Message Latency History"
	QueueMessageCounterName          = "queue_message_counter"
	QueueMessageCounterDesc          = "Queue Message Counter"
)

type Client interface {
	GetLatencyHistogram() metrics.Histogram
	GetMessageCounter() metrics.Counter
}

type client struct {
	client     metrics.MetricProvider
	latency    metrics.Histogram
	msgCounter metrics.Counter
	tags       []tags.Tag
}

func (m client) GetLatencyHistogram() metrics.Histogram {
	return m.latency
}

func (m client) GetMessageCounter() metrics.Counter {
	return m.msgCounter
}

func New(queueName string, provider metrics.MetricProvider, options ...MetricOption) (Client, error) {
	m := client{
		client: provider,
	}

	for _, option := range options {
		option(&m)
	}

	m.tags = append(m.tags, tags.QueueName(queueName))

	latency, err := provider.GetHistogram(
		QueueMessageLatencyHistogramName,
		metrics.NewInstrumentOptions().
			WithTags(m.tags...).
			WithDesc(QueueMessageLatencyHistogramDesc),
	)

	if err != nil {
		return client{}, err
	}

	msgCounter, err := provider.GetCounter(
		QueueMessageCounterName,
		metrics.NewInstrumentOptions().
			WithTags(m.tags...).
			WithDesc(QueueMessageCounterDesc),
	)

	if err != nil {
		return client{}, err
	}

	m.latency = latency
	m.msgCounter = msgCounter

	return m, nil
}

type MetricOption func(*client)

func WithMetricTags(tags ...tags.Tag) MetricOption {
	return func(m *client) {
		m.tags = append(m.tags, tags...)
	}
}
