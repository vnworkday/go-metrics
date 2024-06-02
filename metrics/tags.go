package metrics

import (
	"github.com/vnworkday/go-metrics/status"
	"go.opentelemetry.io/otel/attribute"
)

/**
 * Tags provides methods for creating and working with tags.
 */

func ServiceTag(name string) Tag {
	return NewTag("service", name)
}

func OpTag(name string) Tag {
	return NewTag("op", name)
}

func StatusTag(status status.Status) Tag {
	return NewTag("status", status.String())
}

func APITag(name string) Tag {
	return NewTag("api", name)
}

func QueueTag(name string) Tag {
	return NewTag("queue", name)
}

func QueueTypeTag(name string) Tag {
	return NewTag("queue_type", name)
}

func QueueRoleTag(name string) Tag {
	return NewTag("queue_role", name)
}

func ErrorTypeTag(name string) Tag {
	return NewTag("error_type", name)
}

func tagsToAttributes(tags []Tag) []attribute.KeyValue {
	attrs := make([]attribute.KeyValue, 0, len(tags))
	for _, tag := range tags {
		attrs = append(attrs, attribute.Key(tag.Name()).String(tag.Value()))
	}
	return attrs
}

func attributesToTags(attrs []attribute.KeyValue) []Tag {
	tags := make([]Tag, 0, len(attrs))
	for _, attr := range attrs {
		tags = append(tags, NewTag(string(attr.Key), attr.Value.AsString()))
	}
	return tags
}

func appendTagsToAttributes(attrs []attribute.KeyValue, tags ...Tag) []attribute.KeyValue {
	return append(attrs, tagsToAttributes(tags)...)
}
