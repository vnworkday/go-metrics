package metrics

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type otelHistogram struct {
	instrument[metric.Int64Histogram]
}

func newOtelHistogram(name string, histogram metric.Int64Histogram, tagTracker Tracker, attrs ...attribute.KeyValue) otelHistogram {
	return otelHistogram{
		instrument: newInstrument(histogram, name, tagTracker, attrs...),
	}
}

func (h otelHistogram) Record(ctx context.Context, value int, tags ...Tag) {
	unionTags := append(attributesToTags(h.attrs), tags...)
	cleanedTags := h.tagTracker.CleanTags(h.name, unionTags)
	h.metric.Record(ctx, int64(value), metric.WithAttributes(tagsToAttributes(cleanedTags)...))
}
