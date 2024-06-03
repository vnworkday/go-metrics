package apimetrics

import (
	"github.com/vnworkday/go-metrics/metrics"
	"github.com/vnworkday/go-metrics/statuses"
	"github.com/vnworkday/go-metrics/tags"
)

type ExecOption[T any] func(parameters *ExecParameters[T])

type ExecParameters[T any] struct {
	statusConverter      statuses.StatusConverter[T]
	errTypeConverter     metrics.ErrTypeConverter
	tags                 []tags.Tag
	batchSizeCounterName string
	latencyHistogramName string
}

func NewExecParameters[T any]() ExecParameters[T] {
	return ExecParameters[T]{
		statusConverter:      statuses.DefaultConverter[T],
		errTypeConverter:     metrics.DefaultErrTypeConverter,
		batchSizeCounterName: "batch_size_counter",
		latencyHistogramName: "e2e_latency",
	}
}

func WithErrTypeConverterWithResponse[T any](errTypeConverter metrics.ErrTypeConverter) ExecOption[T] {
	return func(parameters *ExecParameters[T]) {
		parameters.errTypeConverter = errTypeConverter
	}
}

func WithErrTypeConverter(errTypeConverter metrics.ErrTypeConverter) ExecOption[any] {
	return WithErrTypeConverterWithResponse[any](func(err error) string {
		return errTypeConverter(err)
	})
}

func WithStatusConverterWithResponse[T any](statusConverter statuses.StatusConverter[T]) ExecOption[T] {
	return func(parameters *ExecParameters[T]) {
		parameters.statusConverter = statusConverter
	}
}

func WithStatusConverter(converter func(error) statuses.Status) ExecOption[any] {
	return WithStatusConverterWithResponse(func(_ any, err error) statuses.Status {
		return converter(err)
	})
}

func WithTags[T any](tags ...tags.Tag) ExecOption[T] {
	return func(parameters *ExecParameters[T]) {
		parameters.tags = append(parameters.tags, tags...)
	}
}
