package tags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagsQueue(t *testing.T) {
	assertThat := assert.New(t)

	tests := []struct {
		name            string
		queueName       string
		queueType       string
		queueRole       string
		expectedNameTag Tag
		expectedTypeTag Tag
		expectedRoleTag Tag
	}{
		{
			name:            "QueueTagCreation",
			queueName:       "queue1",
			queueType:       "type1",
			queueRole:       "role1",
			expectedNameTag: NewTag("queue", "queue1"),
			expectedTypeTag: NewTag("queue_type", "type1"),
			expectedRoleTag: NewTag("queue_role", "role1"),
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nameTag := QueueName(tt.queueName)
			typeTag := QueueType(tt.queueType)
			roleTag := QueueRole(tt.queueRole)
			assertThat.Equal(tt.expectedNameTag, nameTag)
			assertThat.Equal(tt.expectedTypeTag, typeTag)
			assertThat.Equal(tt.expectedRoleTag, roleTag)
		})
	}
}
