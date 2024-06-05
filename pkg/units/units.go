package units

// ToOtelUnit tries to convert a unit to an OpenTelemetry unit string.
// If the unit is invalid, an error is returned.
// If the unit is empty, an empty string is returned.
func ToOtelUnit(metric string, unit Unit) (string, error) {
	if unit == "" {
		return "", nil
	}

	switch unit {
	case Dimensionless:
		return "1", nil
	case Bytes:
		return "by", nil
	case Millis:
		return "ms", nil
	default:
		return "", ErrUnitInvalid
	}
}
