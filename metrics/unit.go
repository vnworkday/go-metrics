package metrics

type Unit string

const (
	Dimensionless Unit = "1"
	Bytes         Unit = "b"
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

var _ error = InvalidUnit{}

type InvalidUnit struct{}

func (InvalidUnit) Error() string {
	return "invalid unit"
}
