package warnings

import (
	"testing"

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
			warning:  newWarning("metric1", "message1"),
			expected: "metric1, message1, empty tags",
		},
		{
			name:     "WithTag",
			warning:  newWarning("metric1", "message1").withTag("tag1", "value1"),
			expected: "value1",
		},
		{
			name:     "Metric",
			warning:  newWarning("metric1", "message1"),
			expected: "metric1",
		},
		{
			name:     "Message",
			warning:  newWarning("metric1", "message1"),
			expected: "message1",
		},
		{
			name:     "Tags",
			warning:  newWarning("metric1", "message1").withTag("tag1", "value1"),
			expected: "value1",
		},
		{
			name: "String",
			warning: newWarning("metric1", "message1").
				withTag("tag1", "value1").
				withTag("tag2", "value2").
				withTag("tag3", "value3"),
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
				assertThat.Equal(tt.expected, tt.warning.String())
			}
		})
	}
}
