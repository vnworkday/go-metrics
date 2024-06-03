package metrics

import (
	"context"

	t "github.com/vnworkday/go-metrics/tags"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type otelGauge struct {
	instrument[metric.Int64Gauge]
}

func newOtelGauge(name string, gauge metric.Int64Gauge, tagCleaner t.TagCleaner, attrs ...attribute.KeyValue) otelGauge {
	return otelGauge{
		instrument: newInstrument(gauge, name, tagCleaner, attrs...),
	}
}

func (g otelGauge) Record(ctx context.Context, value int, tags ...t.Tag) {
	unionTags := append(t.ToTags(g.attrs), tags...)
	cleanedTags := g.tagCleaner.Clean(g.name, unionTags)
	g.metric.Record(ctx, int64(value), metric.WithAttributes(t.ToAttributes(cleanedTags)...))
}
