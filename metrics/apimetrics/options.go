package apimetrics

import (
	"github.com/vnworkday/go-metrics/metrics"
	"github.com/vnworkday/go-metrics/status"
)

type ExecOption[T any] func(parameters *ExecParameters[T])

type ExecParameters[T any] struct {
	statusConverter      status.Converter[T]
	errTypeConverter     metrics.ErrTypeConverter
	tags                 []metrics.Tag
	batchSizeCounterName string
	latencyHistogramName string
}

func NewExecParameters[T any]() ExecParameters[T] {
	return ExecParameters[T]{
		statusConverter:      status.DefaultConverter[T],
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

func WithStatusConverterWithResponse[T any](statusConverter status.Converter[T]) ExecOption[T] {
	return func(parameters *ExecParameters[T]) {
		parameters.statusConverter = statusConverter
	}
}

func WithStatusConverter(converter func(error) status.Status) ExecOption[any] {
	return WithStatusConverterWithResponse(func(_ any, err error) status.Status {
		return converter(err)
	})
}

func WithTags[T any](tags ...metrics.Tag) ExecOption[T] {
	return func(parameters *ExecParameters[T]) {
		parameters.tags = append(parameters.tags, tags...)
	}
}
