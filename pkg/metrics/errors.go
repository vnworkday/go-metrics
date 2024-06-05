package metrics

import "github.com/pkg/errors"

var ErrMultipleUnitsSpecified = errors.New("multiple units specified")
var ErrNoInstrumentOptionsSpecified = errors.New("no instrument options specified")

type ErrTypeConverter func(error) string

func DefaultErrTypeConverter(err error) string {
	if err == nil {
		return ""
	}

	return "unknown error"
}
