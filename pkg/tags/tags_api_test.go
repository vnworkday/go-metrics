package tags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagsApi(t *testing.T) {
	assertThat := assert.New(t)

	tests := []struct {
		name     string
		apiName  string
		expected Tag
	}{
		{
			name:     "APINameTagCreation",
			apiName:  "api1",
			expected: NewTag("api", "api1"),
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag := APIName(tt.apiName)
			assertThat.Equal(tt.expected, tag)
		})
	}

}
