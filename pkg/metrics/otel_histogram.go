package metrics

import (
	"context"

	t "github.com/vnworkday/go-metrics/pkg/tags"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type OtelHistogram struct {
	instrument[metric.Int64Histogram]
}

func newOtelHistogram(name string, histogram metric.Int64Histogram, tagCleaner t.TagCleaner, attrs ...attribute.KeyValue) OtelHistogram {
	return OtelHistogram{
		instrument: newInstrument(histogram, name, tagCleaner, attrs...),
	}
}

func (h OtelHistogram) Record(ctx context.Context, value int, tags ...t.Tag) {
	unionTags := append(t.ToTags(h.attrs), tags...)
	cleanedTags := h.tagCleaner.Clean(h.name, unionTags)
	h.metric.Record(ctx, int64(value), metric.WithAttributes(t.ToAttributes(cleanedTags)...))
}
