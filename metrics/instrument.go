package metrics

import (
	"go.opentelemetry.io/otel/attribute"
)

type instrument[T any] struct {
	metric     T
	attrs      []attribute.KeyValue
	name       string
	tagTracker Tracker
}

func newInstrument[T any](metric T, name string, tagTracker Tracker, attrs ...attribute.KeyValue) instrument[T] {
	return instrument[T]{
		metric:     metric,
		attrs:      attrs,
		name:       name,
		tagTracker: tagTracker,
	}
}
