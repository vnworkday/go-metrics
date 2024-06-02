package metrics

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type otelGauge struct {
	instrument[metric.Int64Gauge]
}

func newOtelGauge(name string, gauge metric.Int64Gauge, tagTracker Tracker, attrs ...attribute.KeyValue) otelGauge {
	return otelGauge{
		instrument: newInstrument(gauge, name, tagTracker, attrs...),
	}
}

func (g otelGauge) Record(ctx context.Context, value int, tags ...Tag) {
	unionTags := append(attributesToTags(g.attrs), tags...)
	cleanedTags := g.tagTracker.CleanTags(g.name, unionTags)
	g.metric.Record(ctx, int64(value), metric.WithAttributes(tagsToAttributes(cleanedTags)...))
}
