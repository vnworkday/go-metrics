package tags

import (
	"fmt"
	"sync"

	"github.com/vnworkday/go-metrics/common"
	"github.com/vnworkday/go-metrics/warnings"
)

const (
	TagMaxSize     = 120 // Maximum number of tags that can be added.
	TagNameMaxLen  = 200 // Maximum length of a single tag name.
	TagValueMaxLen = 200 // Maximum length of a single tag value.
)

const (
	TagRedactedDueToOverflow = "tag_redacted_due_to_overflow"
)

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

type ErrTagNameTooLong struct{}

func (e ErrTagNameTooLong) Error() string {
	return "tag name is too long"
}

type ErrTagValueTooLong struct{}

func (e ErrTagValueTooLong) Error() string {
	return "tag value is too long"
}

// TagCleaner is an interface that defines a method for cleaning tags.
type TagCleaner interface {
	// Clean trim the tag name and value to the maximum length, and redact all subsequent tags if the number of allowed tags reached the limit.
	Clean(metric string, tags []Tag) []Tag
}

// DefaultTagCleaner is a tag cleaner that trims tag names and values to the maximum length, and redacts tags if the number of allowed tags is exceeded.
type DefaultTagCleaner struct {
	// metricTags is a map of metric name to a map of tag name to a set of tag values.
	metricTags map[string]map[string]common.Set[string]
	// warningHandler is a function that handles warnings generated by the tag store.
	warningHandler warnings.WarningHandler
	// lock is a mutex that protects the tag store from concurrent access.
	lock sync.Mutex
}

func (t *DefaultTagCleaner) Clean(metric string, tags []Tag) []Tag {
	t.lock.Lock()
	defer t.lock.Unlock()

	cleanedTags := make([]Tag, 0, len(tags))

	for _, tag := range tags {
		if t.metricTags[metric] == nil {
			t.metricTags[metric] = make(map[string]common.Set[string])
		}

		trimmedTag, nameErr, valueErr := TrimTag(tag)

		if nameErr != nil {
			t.warningHandler(warnings.TagNameTooLong(metric, tag.Name(), tag.Value()))
		}

		if valueErr != nil {
			t.warningHandler(warnings.TagValueTooLong(metric, tag.Name(), tag.Value()))
		}

		tagValues := t.metricTags[metric][trimmedTag.Name()]

		if tagValues.IsEmpty() {
			tagValues = common.NewSet[string]()
		}

		// We only redact the tag value if the tag value is not in the set of tag values given by the tag name.
		if tagValues.Len() >= TagMaxSize && !tagValues.Contains(trimmedTag.Value()) {
			trimmedTag = NewTag(trimmedTag.Name(), TagRedactedDueToOverflow)
			t.warningHandler(warnings.TagOverflow(metric, tag.Name(), tag.Value()))
		}

		t.metricTags[metric][trimmedTag.Name()] = tagValues.Add(trimmedTag.Value())
		cleanedTags = append(cleanedTags, trimmedTag)
	}

	return cleanedTags
}

func NewTagCleaner(warningHandler warnings.WarningHandler) TagCleaner {
	return &DefaultTagCleaner{
		metricTags:     make(map[string]map[string]common.Set[string]),
		warningHandler: warningHandler,
	}
}

// NoopTagCleaner is a tag cleaner that does not clean tags.
type NoopTagCleaner struct{}

func (t *NoopTagCleaner) Clean(metric string, tags []Tag) []Tag {
	return tags
}
