package units

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

type ErrUnitInvalid struct{}

func (ErrUnitInvalid) Error() string {
	return "invalid unit"
}
