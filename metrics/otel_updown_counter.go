package metrics

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type otelUpDownCounter struct {
	instrument[metric.Int64UpDownCounter]
}

func newOtelUpDownCounter(name string, upDownCounter metric.Int64UpDownCounter, tagTracker Tracker, attrs ...attribute.KeyValue) otelUpDownCounter {
	return otelUpDownCounter{
		instrument: newInstrument(upDownCounter, name, tagTracker, attrs...),
	}
}

func (c otelUpDownCounter) Add(ctx context.Context, value uint, tags ...Tag) {
	unionTags := append(attributesToTags(c.attrs), tags...)
	cleanedTags := c.tagTracker.CleanTags(c.name, unionTags)
	c.metric.Add(ctx, int64(value), metric.WithAttributes(tagsToAttributes(cleanedTags)...))
}
