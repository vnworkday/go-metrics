package metrics

import (
	"context"

	t "github.com/vnworkday/go-metrics/pkg/tags"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type otelUpDownCounter struct {
	instrument[metric.Int64UpDownCounter]
}

func newOtelUpDownCounter(name string, upDownCounter metric.Int64UpDownCounter, tagCleaner t.TagCleaner, attrs ...attribute.KeyValue) otelUpDownCounter {
	return otelUpDownCounter{
		instrument: newInstrument(upDownCounter, name, tagCleaner, attrs...),
	}
}

func (c otelUpDownCounter) Add(ctx context.Context, value uint, tags ...t.Tag) {
	unionTags := append(t.ToTags(c.attrs), tags...)
	cleanedTags := c.tagCleaner.Clean(c.name, unionTags)
	c.metric.Add(ctx, int64(value), metric.WithAttributes(t.ToAttributes(cleanedTags)...))
}
