package warnings

import (
	"testing"

	"github.com/vnworkday/go-metrics/internal/common"

	"github.com/stretchr/testify/assert"
)

func TestWarning(t *testing.T) {
	assertThat := assert.New(t)

	tests := []struct {
		name     string
		warning  Warning
		expected string
	}{
		{
			name:     "NewWarning",
			warning:  NewWarning("metric1", "message1"),
			expected: "metric1, message1, empty tags",
		},
		{
			name:     "WithTag",
			warning:  NewWarning("metric1", "message1").addTag("tag1", "value1"),
			expected: "value1",
		},
		{
			name:     "Metric",
			warning:  NewWarning("metric1", "message1"),
			expected: "metric1",
		},
		{
			name:     "Message",
			warning:  NewWarning("metric1", "message1"),
			expected: "message1",
		},
		{
			name:     "Tags",
			warning:  NewWarning("metric1", "message1").addTag("tag1", "value1"),
			expected: "value1",
		},
		{
			name: "String",
			warning: NewWarning("metric1", "message1").
				addTag("tag1", "value1").
				addTag("tag2", "value2").
				addTag("tag3", "value3"),
			expected: `metric="metric1" message="message1" label_name="tag1" label_value="value1" label_name="tag2" label_value="value2" label_name="tag3" label_value="value3"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "NewWarning":
				assertThat.Equal(tt.expected, tt.warning.metric+", "+tt.warning.message+", empty tags")
			case "WithTag":
				assertThat.Equal(tt.expected, tt.warning.tags["tag1"])
			case "Metric":
				assertThat.Equal(tt.expected, tt.warning.Metric())
			case "Message":
				assertThat.Equal(tt.expected, tt.warning.Message())
			case "Tags":
				assertThat.Equal(tt.expected, tt.warning.Tags()["tag1"])
			case "String":
				assertThat.True(common.StringToSet(tt.expected, " ").Equals(common.StringToSet(tt.warning.String(), " ")))
			}
		})
	}
}
