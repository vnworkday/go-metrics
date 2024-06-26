package apimetrics

import (
	"context"
	"time"

	"github.com/vnworkday/go-metrics/internal/common"
	"github.com/vnworkday/go-metrics/pkg/metrics"
	"github.com/vnworkday/go-metrics/pkg/tags"
)

func collectParams[T any](opName string, options ...metrics.ExecOption[T]) metrics.ExecParameters[T] {
	params := metrics.NewExecParameters[T]()
	for _, option := range options {
		option(&params)
	}
	params.Tags = append(params.Tags, tags.APIOp(opName))
	return params
}

// DoRequest executes a request and records metrics for the request.
// The metrics recorded are:
// - A counter for the number of requests made.
// - A histogram for the latency of the request.
// The following tags are always present in the metrics:
// - The operation name.
// - The statuses of the request.
// - The error type of the request (if an error occurred).
// The statuses and error type are determined by the statusConverter and errTypeConverter functions passed in the options.
//
// @see DoRequestWithResponse for more details.
func DoRequest(
	ctx context.Context,
	metric Client,
	opName string,
	makeRequestWoResponse func() error,
	options ...metrics.ExecOption[any],
) error {
	makeRequest := func() (interface{}, error) {
		return common.Empty{}, makeRequestWoResponse()
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
// - The statuses of the request.
// - The error type of the request (if an error occurred).
// The statuses and error type are determined by the statusConverter and errTypeConverter functions passed in the options.
func DoRequestWithResponse[T any](
	ctx context.Context,
	metric Client,
	opName string,
	makeRequest func() (T, error),
	options ...metrics.ExecOption[T],
) (T, error) {
	params := collectParams(opName, options...)

	metric.GetRequestCounter().Add(ctx, 1, params.Tags...)

	startTime := time.Now().UTC()
	resp, err := makeRequest()
	latency := time.Now().UTC().Sub(startTime)

	var metricTags []tags.Tag

	if err != nil {
		metricTags = append(params.Tags,
			tags.Status(params.StatusConverter(resp, err)),
			tags.ErrorType(params.ErrTypeConverter(err)),
		)
	} else {
		metricTags = append(params.Tags,
			tags.Status(params.StatusConverter(resp, nil)),
		)
	}

	metric.GetLatencyHistogram().Record(ctx, int(latency.Milliseconds()), metricTags...)

	return resp, err
}
