package warnings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagOverflowCreatesCorrectWarning(t *testing.T) {
	assertThat := assert.New(t)
	w := TagOverflow("metric1", "tag1", "value1")
	assertThat.Equal("metric1", w.Metric())
	assertThat.Equal(WarningTagCollectionLimitReachedMessage, w.Message())
	assertThat.Equal("value1", w.Tags()["tag1"])
}

func TestTagNameTooLongCreatesCorrectWarning(t *testing.T) {
	assertThat := assert.New(t)
	w := TagNameTooLong("metric1", "tag1", "value1")
	assertThat.Equal("metric1", w.Metric())
	assertThat.Equal(WarningTagNameSizeLimitReachedMessage, w.Message())
	assertThat.Equal("value1", w.Tags()["tag1"])
}

func TestTagValueTooLongCreatesCorrectWarning(t *testing.T) {
	assertThat := assert.New(t)
	w := TagValueTooLong("metric1", "tag1", "value1")
	assertThat.Equal("metric1", w.Metric())
	assertThat.Equal(WarningTagValueSizeLimitReachedMessage, w.Message())
	assertThat.Equal("value1", w.Tags()["tag1"])
}
