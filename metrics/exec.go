package metrics

import (
	"github.com/vnworkday/go-metrics/statuses"
	"github.com/vnworkday/go-metrics/tags"
)

type ExecOption[T any] func(parameters *ExecParameters[T])

type ExecParameters[T any] struct {
	StatusConverter      statuses.StatusConverter[T]
	ErrTypeConverter     ErrTypeConverter
	Tags                 []tags.Tag
	BatchSizeCounterName string // used for API metrics only
	LatencyHistogramName string
}

func NewExecParameters[T any]() ExecParameters[T] {
	return ExecParameters[T]{
		StatusConverter:      statuses.DefaultConverter[T],
		ErrTypeConverter:     DefaultErrTypeConverter,
		BatchSizeCounterName: "batch_size_counter",
		LatencyHistogramName: "e2e_latency",
	}
}

func WithErrTypeConverterWithResponse[T any](errTypeConverter ErrTypeConverter) ExecOption[T] {
	return func(parameters *ExecParameters[T]) {
		parameters.ErrTypeConverter = errTypeConverter
	}
}

func WithErrTypeConverter(errTypeConverter ErrTypeConverter) ExecOption[any] {
	return WithErrTypeConverterWithResponse[any](func(err error) string {
		return errTypeConverter(err)
	})
}

func WithStatusConverterWithResponse[T any](statusConverter statuses.StatusConverter[T]) ExecOption[T] {
	return func(parameters *ExecParameters[T]) {
		parameters.StatusConverter = statusConverter
	}
}

func WithStatusConverter(converter func(error) statuses.Status) ExecOption[any] {
	return WithStatusConverterWithResponse(func(_ any, err error) statuses.Status {
		return converter(err)
	})
}

func WithExecTags[T any](tags ...tags.Tag) ExecOption[T] {
	return func(parameters *ExecParameters[T]) {
		parameters.Tags = append(parameters.Tags, tags...)
	}
}