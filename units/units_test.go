package units

import "testing"

func TestEmptyUnitReturnsEmptyString(t *testing.T) {
	u, err := ToOtelUnit("metric", "")
	if u != "" || err != nil {
		t.Errorf("Empty unit should return an empty string and no error")
	}
}

func TestDimensionlessUnitReturnsOne(t *testing.T) {
	u, err := ToOtelUnit("metric", Dimensionless)
	if u != "1" || err != nil {
		t.Errorf("Dimensionless unit should return '1' and no error")
	}
}

func TestBytesUnitReturnsBy(t *testing.T) {
	u, err := ToOtelUnit("metric", Bytes)
	if u != "by" || err != nil {
		t.Errorf("Bytes unit should return 'by' and no error")
	}
}

func TestMillisUnitReturnsMs(t *testing.T) {
	u, err := ToOtelUnit("metric", Millis)
	if u != "ms" || err != nil {
		t.Errorf("Millis unit should return 'ms' and no error")
	}
}

func TestInvalidUnitReturnsError(t *testing.T) {
	u, err := ToOtelUnit("metric", "invalid")
	if u != "" || err == nil {
		t.Errorf("Invalid unit should return an empty string and an error")
	}
}
