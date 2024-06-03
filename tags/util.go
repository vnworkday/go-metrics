package tags

import "go.opentelemetry.io/otel/attribute"

func ToAttribute(tag Tag) attribute.KeyValue {
	return attribute.Key(tag.Name()).String(tag.Value())
}

func FromAttribute(attr attribute.KeyValue) Tag {
	return NewTag(string(attr.Key), attr.Value.AsString())
}

func ToAttributes(tags []Tag) []attribute.KeyValue {
	attrs := make([]attribute.KeyValue, 0, len(tags))
	for i, tag := range tags {
		attrs[i] = ToAttribute(tag)
	}
	return attrs
}

func ToTags(attrs []attribute.KeyValue) []Tag {
	tags := make([]Tag, 0, len(attrs))
	for i, attr := range attrs {
		tags[i] = FromAttribute(attr)
	}
	return tags
}

func AddTags(attrs []attribute.KeyValue, tags ...Tag) []attribute.KeyValue {
	return append(attrs, ToAttributes(tags)...)
}

// TrimTag trims the name and value of a tag to the maximum length.
// It returns the trimmed tag, and any errors that occurred while trimming the name and value.
func TrimTag(tag Tag) (Tag, error, error) {
	name := tag.Name()
	value := tag.Value()

	var nameErr, valueErr error

	if len(name) > TagNameMaxLen {
		name = name[:TagNameMaxLen]
		nameErr = ErrTagNameTooLong{}
	}

	if len(value) > TagValueMaxLen {
		value = value[:TagValueMaxLen]
		valueErr = ErrTagValueTooLong{}
	}

	trimmedTag := NewTag(name, value)

	return trimmedTag, nameErr, valueErr
}
