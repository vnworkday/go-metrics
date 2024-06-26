package metrics

import (
	"context"

	t "github.com/vnworkday/go-metrics/pkg/tags"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type OtelCounter struct {
	instrument[metric.Int64Counter]
}

func newOtelCounter(name string, counter metric.Int64Counter, tagCleaner t.TagCleaner, attrs ...attribute.KeyValue) OtelCounter {
	return OtelCounter{
		instrument: newInstrument(counter, name, tagCleaner, attrs...),
	}
}

func (c OtelCounter) Add(ctx context.Context, value uint, tags ...t.Tag) {
	unionTags := append(t.ToTags(c.attrs), tags...)
	cleanedTags := c.tagCleaner.Clean(c.name, unionTags)
	c.metric.Add(ctx, int64(value), metric.WithAttributes(t.ToAttributes(cleanedTags)...))
}
