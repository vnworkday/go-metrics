package warnings

import (
	"testing"

	"github.com/vnworkday/go-metrics/internal/common"

	"github.com/stretchr/testify/assert"
)

func TestNewWarning(t *testing.T) {
	assertThat := assert.New(t)
	w := NewWarning("metric1", "message1")
	assertThat.Equal("metric1", w.Metric())
	assertThat.Equal("message1", w.Message())
	assertThat.Empty(w.Tags())
}

func TestAddTag(t *testing.T) {
	assertThat := assert.New(t)
	w := NewWarning("metric1", "message1").addTag("tag1", "value1")
	assertThat.Equal("value1", w.Tags()["tag1"])
}

func TestMetric(t *testing.T) {
	assertThat := assert.New(t)
	w := NewWarning("metric1", "message1")
	assertThat.Equal("metric1", w.Metric())
}

func TestMessage(t *testing.T) {
	assertThat := assert.New(t)
	w := NewWarning("metric1", "message1")
	assertThat.Equal("message1", w.Message())
}

func TestTags(t *testing.T) {
	assertThat := assert.New(t)
	w := NewWarning("metric1", "message1").addTag("tag1", "value1")
	assertThat.Equal("value1", w.Tags()["tag1"])
}

func TestString(t *testing.T) {
	assertThat := assert.New(t)
	w := NewWarning("metric1", "message1").
		addTag("tag1", "value1").
		addTag("tag2", "value2").
		addTag("tag3", "value3")
	expected := `metric="metric1" message="message1" label_name="tag1" label_value="value1" label_name="tag2" label_value="value2" label_name="tag3" label_value="value3"`
	assertThat.True(common.StringToSet(expected, " ").Equals(common.StringToSet(w.String(), " ")))
}
