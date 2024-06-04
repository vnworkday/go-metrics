package queuemetrics

import (
	"time"

	"github.com/vnworkday/go-metrics/metrics"
	"github.com/vnworkday/go-metrics/tags"
)

const (
	QueueMessageLatencyHistogramName = "queue_message_latency_histogram"
	QueueMessageLatencyHistogramDesc = "Queue Message Latency History"
	QueueMessageCounterName          = "queue_message_counter"
	QueueMessageCounterDesc          = "Queue Message Counter"
)

type QueueMetrics interface {
	metrics.Metrics
	GetMessageCounter() metrics.Counter
}

type Metric struct {
	metrics.Client
	name       string
	latency    metrics.Histogram
	msgCounter metrics.Counter
	tags       []tags.Tag
}

func (m Metric) GetLatencyHistogram() metrics.Histogram {
	return m.latency
}

func (m Metric) GetMessageCounter() metrics.Counter {
	return m.msgCounter
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

	m.tags = append(m.tags, tags.QueueName(name))

	latency, err := client.GetHistogram(
		QueueMessageLatencyHistogramName,
		metrics.NewInstrumentOptions().
			WithTags(m.tags...).
			WithDesc(QueueMessageLatencyHistogramDesc),
	)

	if err != nil {
		return Metric{}, err
	}

	msgCounter, err := client.GetCounter(
		QueueMessageCounterName,
		metrics.NewInstrumentOptions().
			WithTags(m.tags...).
			WithDesc(QueueMessageCounterDesc),
	)

	if err != nil {
		return Metric{}, err
	}

	m.latency = latency
	m.msgCounter = msgCounter

	return m, nil
}

type MetricOption func(*Metric)

func WithMetricTags(tags ...tags.Tag) MetricOption {
	return func(m *Metric) {
		m.tags = append(m.tags, tags...)
	}
}
