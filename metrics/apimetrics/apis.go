package apimetrics

import (
	"context"

	"github.com/vnworkday/go-metrics/metrics"
)

type empty struct{}

func collectParams[T any](opName string, options ...ExecOption[T]) ExecParameters[T] {
	params := NewExecParameters[T]()
	for _, option := range options {
		option(&params)
	}
	params.tags = append(params.tags, metrics.OpTag(opName))
	return params
}

// DoRequest executes a request and records metrics for the request.
// The metrics recorded are:
// - A counter for the number of requests made.
// - A histogram for the latency of the request.
// The following tags are always present in the metrics:
// - The operation name.
// - The status of the request.
// - The error type of the request (if an error occurred).
// The status and error type are determined by the statusConverter and errTypeConverter functions passed in the options.
//
// @see DoRequestWithResponse for more details.
func DoRequest(
	ctx context.Context,
	metric Metrics,
	opName string,
	makeRequestWoResponse func() error,
	options ...ExecOption[any],
) error {
	makeRequest := func() (interface{}, error) {
		return empty{}, makeRequestWoResponse()
	}

	_, err := DoRequestWithResponse(ctx, metric, opName, makeRequest, options...)

	return err
}

// DoRequestWithResponse executes a request and records metrics for the request.
// The metrics recorded are:
// - A counter for the number of requests made.
// - A histogram for the latency of the request.
// The following tags are always present in the metrics:
// - The operation name.
// - The status of the request.
// - The error type of the request (if an error occurred).
// The status and error type are determined by the statusConverter and errTypeConverter functions passed in the options.
func DoRequestWithResponse[T any](
	ctx context.Context,
	metric Metrics,
	opName string,
	makeRequest func() (T, error),
	options ...ExecOption[T],
) (T, error) {
	params := collectParams(opName, options...)

	metric.GetRequestCounter().Add(ctx, 1, params.tags...)

	startTime := metric.UtcNow()
	resp, err := makeRequest()
	latency := metric.UtcNow().Sub(startTime)

	var tags []metrics.Tag

	if err != nil {
		tags = append(params.tags,
			metrics.StatusTag(params.statusConverter(resp, err)),
			metrics.ErrorTypeTag(params.errTypeConverter(err)),
		)
	} else {
		tags = append(params.tags,
			metrics.StatusTag(params.statusConverter(resp, nil)),
		)
	}

	metric.GetLatencyHistogram().Record(ctx, int(latency.Milliseconds()), tags...)

	return resp, err
}
