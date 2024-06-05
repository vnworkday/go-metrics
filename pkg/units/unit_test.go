package units

import "testing"

func TestValidUnitReturnsTrue(t *testing.T) {
	u := Unit("1")
	if !u.Valid() {
		t.Errorf("Valid unit should return true")
	}
}

func TestInvalidUnitReturnsFalse(t *testing.T) {
	u := Unit("invalid")
	if u.Valid() {
		t.Errorf("Invalid unit should return false")
	}
}

func TestErrUnitInvalidReturnsCorrectMessage(t *testing.T) {
	err := ErrUnitInvalid
	if err.Error() != "invalid unit" {
		t.Errorf("ErrUnitInvalid should return 'invalid unit'")
	}
}
