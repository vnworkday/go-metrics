package warnings

import "fmt"

// ****************************************************************************
// This file contains the functions to create warnings related to units.
// ****************************************************************************

// UnitInvalid used when the unit is invalid.
func UnitInvalid(metric, unit string) Warning {
	return NewWarning(metric, fmt.Sprintf("invalid unit \"%s\" will be ignored", unit))
}
