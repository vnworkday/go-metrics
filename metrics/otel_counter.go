package metrics

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type otelCounter struct {
	instrument[metric.Int64Counter]
}

func newOtelCounter(name string, counter metric.Int64Counter, tagTracker Tracker, attrs ...attribute.KeyValue) otelCounter {
	return otelCounter{
		instrument: newInstrument(counter, name, tagTracker, attrs...),
	}
}

func (c otelCounter) Add(ctx context.Context, value uint, tags ...Tag) {
	unionTags := append(attributesToTags(c.attrs), tags...)
	cleanedTags := c.tagTracker.CleanTags(c.name, unionTags)
	c.metric.Add(ctx, int64(value), metric.WithAttributes(tagsToAttributes(cleanedTags)...))
}
