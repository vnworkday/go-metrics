package metrics

import "github.com/vnworkday/go-metrics/warnings"

// asOtelUnit tries to convert a unit to an OpenTelemetry unit string.
// If the unit is invalid, a warning is generated.
// If the unit is empty, an empty string is returned.
func asOtelUnit(metric string, unit Unit, handler warnings.Handler) string {
	if unit == "" {
		return ""
	}

	if unit.Valid() {
		switch unit {
		case Dimensionless:
			return "1"
		case Bytes:
			return "by"
		case Millis:
			return "ms"
		default:
			handler(warnings.InvalidUnit(metric, string(unit)))
			return ""
		}
	} else {
		handler(warnings.InvalidUnit(metric, string(unit)))
		return ""
	}
}
