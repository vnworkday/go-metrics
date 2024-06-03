package metrics

import (
	"context"

	t "github.com/vnworkday/go-metrics/tags"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type otelHistogram struct {
	instrument[metric.Int64Histogram]
}

func newOtelHistogram(name string, histogram metric.Int64Histogram, tagCleaner t.TagCleaner, attrs ...attribute.KeyValue) otelHistogram {
	return otelHistogram{
		instrument: newInstrument(histogram, name, tagCleaner, attrs...),
	}
}

func (h otelHistogram) Record(ctx context.Context, value int, tags ...t.Tag) {
	unionTags := append(t.ToTags(h.attrs), tags...)
	cleanedTags := h.tagCleaner.Clean(h.name, unionTags)
	h.metric.Record(ctx, int64(value), metric.WithAttributes(t.ToAttributes(cleanedTags)...))
}
