package metrics

import (
	"github.com/vnworkday/go-metrics/tags"
	"go.opentelemetry.io/otel/attribute"
)

type instrument[T any] struct {
	metric     T
	attrs      []attribute.KeyValue
	name       string
	tagCleaner tags.TagCleaner
}

func newInstrument[T any](metric T, name string, tagCleaner tags.TagCleaner, attrs ...attribute.KeyValue) instrument[T] {
	return instrument[T]{
		metric:     metric,
		attrs:      attrs,
		name:       name,
		tagCleaner: tagCleaner,
	}
}
