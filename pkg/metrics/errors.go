package metrics

import "github.com/pkg/errors"

var ErrMultipleUnitsSpecified = errors.New("multiple units specified")
var ErrMetricNameEmpty = errors.New("metric name is empty")
var ErrMetricClientNotSpecified = errors.New("metric client is not specified")

var ErrAPIOpNameEmpty = errors.New("api operation name is empty")

type ErrTypeConverter func(error) string

func DefaultErrTypeConverter(err error) string {
	if err == nil {
		return ""
	}

	return "unknown error"
}
