package units

import "github.com/pkg/errors"

type Unit string

const (
	Dimensionless Unit = "1"
	Bytes         Unit = "by"
	Millis        Unit = "ms"
)

func (u Unit) Valid() bool {
	switch u {
	case Dimensionless, Bytes, Millis:
		return true
	default:
		return false
	}
}

var ErrUnitInvalid = errors.New("invalid unit")
