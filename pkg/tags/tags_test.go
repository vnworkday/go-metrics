package tags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTags(t *testing.T) {
	assertThat := assert.New(t)

	tests := []struct {
		name     string
		tag      Tag
		expected string
	}{
		{
			name:     "NewTag",
			tag:      NewTag("tag1", "value1"),
			expected: "tag1, value1",
		},
		{
			name:     "Name",
			tag:      NewTag("tag1", "value1"),
			expected: "tag1",
		},
		{
			name:     "Value",
			tag:      NewTag("tag1", "value1"),
			expected: "value1",
		},
		{
			name:     "String",
			tag:      NewTag("tag1", "value1"),
			expected: "tag1=value1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "NewTag":
				assertThat.Equal(tt.expected, tt.tag.Name()+", "+tt.tag.Value())
			case "Name":
				assertThat.Equal(tt.expected, tt.tag.Name())
			case "Value":
				assertThat.Equal(tt.expected, tt.tag.Value())
			case "String":
				assertThat.Equal(tt.expected, tt.tag.String())
			}
		})
	}

}
