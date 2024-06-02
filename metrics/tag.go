package metrics

import "fmt"

type Tag struct {
	name  string
	value string
}

func NewTag(name, value string) Tag {
	return Tag{name: name, value: value}
}

func (t Tag) Name() string {
	return t.name
}

func (t Tag) Value() string {
	return t.value
}

// String returns a string representation of the tag in the format: name=value
func (t Tag) String() string {
	return fmt.Sprintf("%s=%s", t.name, t.value)
}
