package tags

// ****************************************************************************
// This file contains the functions to create tags related to queues.
// ****************************************************************************

const (
	TagQueueName = "queue"
	TagQueueType = "queue_type"
	TagQueueRole = "queue_role"
)

// QueueName creates a new tag with the key "queue" and the given queue name as the value.
func QueueName(name string) Tag {
	return NewTag(TagQueueName, name)
}

// QueueType creates a new tag with the key "queue_type" and the given queue type as the value.
func QueueType(name string) Tag {
	return NewTag(TagQueueType, name)
}

// QueueRole creates a new tag with the key "queue_role" and the given queue role as the value.
func QueueRole(name string) Tag {
	return NewTag(TagQueueRole, name)
}
