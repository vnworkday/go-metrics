package queuemetrics

import (
	"context"

	"github.com/vnworkday/go-metrics/common"
	"github.com/vnworkday/go-metrics/metrics"
	"github.com/vnworkday/go-metrics/tags"
)

func collectParams[T any](queueType, queueRole string, options ...metrics.ExecOption[T]) metrics.ExecParameters[T] {
	params := metrics.NewExecParameters[T]()
	for _, option := range options {
		option(&params)
	}
	params.Tags = append(params.Tags, tags.QueueType(queueType))
	params.Tags = append(params.Tags, tags.QueueRole(queueRole))
	return params
}

// DoMessage processes a message and records metrics for the message.
// The metrics recorded are:
// - A counter for the number of messages processed.
// - A histogram for the latency of the message processing.
// The following tags are always present in the metrics:
// - The queue type.
// - The queue role.
// - The statuses of the message processing.
// - The error type of the message processing (if an error occurred).
// The statuses and error type are determined by the statusConverter and errTypeConverter functions passed in the options.
//
// @see DoMessageWithResponse for more details.
func DoMessage(
	ctx context.Context,
	metric QueueMetrics,
	queueType, queueRole string,
	processMessageWoResponse func() error,
	options ...metrics.ExecOption[any],
) error {
	processMessage := func() (interface{}, error) {
		err := processMessageWoResponse()
		return common.Empty{}, err
	}

	_, err := DoMessageWithResponse(ctx, metric, queueType, queueRole, processMessage, options...)
	return err
}

// DoMessageWithResponse processes a message and records metrics for the message.
// The metrics recorded are:
// - A counter for the number of messages processed.
// - A histogram for the latency of the message processing.
// The following tags are always present in the metrics:
// - The queue type.
// - The queue role.
// - The statuses of the message processing.
// - The error type of the message processing (if an error occurred).
// The statuses and error type are determined by the statusConverter and errTypeConverter functions passed in the options.
func DoMessageWithResponse[T any](
	ctx context.Context,
	metric QueueMetrics,
	queueType, queueRole string,
	processMessage func() (T, error),
	options ...metrics.ExecOption[T],
) (interface{}, error) {
	params := collectParams(queueType, queueRole, options...)

	metric.GetMessageCounter().Add(ctx, 1, params.Tags...)

	start := metric.UtcNow()
	resp, err := processMessage()
	latency := metric.UtcNow().Sub(start)

	if err != nil {
		params.Tags = append(params.Tags,
			tags.Status(params.StatusConverter(resp, err)),
			tags.ErrorType(params.ErrTypeConverter(err)),
		)
	} else {
		params.Tags = append(params.Tags, tags.Status(params.StatusConverter(resp, nil)))
	}

	metric.GetLatencyHistogram().Record(ctx, int(latency.Milliseconds()), params.Tags...)

	return resp, err
}
