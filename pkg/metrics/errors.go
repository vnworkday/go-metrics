package metrics

import "github.com/pkg/errors"

var ErrMultipleUnitsSpecified = errors.New("multiple units specified")
var ErrMetricClientNotSpecified = errors.New("metric client is not specified")

var ErrAPINameNotSpecified = errors.New("api name is not specified")

var ErrQueueNameNotSpecified = errors.New("queue name is not specified")

type ErrTypeConverter func(error) string

func DefaultErrTypeConverter(err error) string {
	if err == nil {
		return ""
	}

	return "unknown: " + err.Error()
}
