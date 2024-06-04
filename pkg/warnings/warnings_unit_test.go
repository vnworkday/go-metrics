package warnings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidUnitCreatesCorrectWarning(t *testing.T) {
	assertThat := assert.New(t)
	w := UnitInvalid("metric1", "unit1")
	assertThat.Equal("metric1", w.Metric())
	assertThat.Equal(`invalid unit "unit1" will be ignored`, w.Message())
	assertThat.Empty(w.Tags())
}
