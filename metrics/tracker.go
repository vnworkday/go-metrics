package metrics

import (
	"github.com/vnworkday/go-metrics/common"
	"github.com/vnworkday/go-metrics/warnings"
	"sync"
)

const (
	CardinalityOverflowValue = "cardinality_overflow"
	CardinalityOverflowLimit = 120
	NameMaxLen               = 200
	ValueMaxLen              = 200
)

type Tracker interface {
	CleanTags(metric string, tags []Tag) []Tag
}

type tracker struct {
	// trackable is a map of metric name to a map of tag name to a set of tag values.
	trackable      map[string]map[string]common.Set[string]
	lock           sync.Mutex
	warningHandler warnings.Handler
}

func NewTracker(warningHandler warnings.Handler) Tracker {
	return &tracker{
		trackable:      make(map[string]map[string]common.Set[string]),
		warningHandler: warningHandler,
	}
}

var _ Tracker = (*tracker)(nil)

func (t *tracker) CleanTags(metric string, tags []Tag) []Tag {
	t.lock.Lock()
	defer t.lock.Unlock()

	cleanedTags := make([]Tag, 0, len(tags))
	for _, tag := range tags {
		if t.trackable[metric] == nil {
			t.trackable[metric] = make(map[string]common.Set[string])
		}

		trimmedTag := t.trimTag(metric, tag)
		tagValues := t.trackable[metric][trimmedTag.Name()]

		if tagValues.IsEmpty() {
			tagValues = common.NewSet[string]()
		}

		if tagValues.Len() >= CardinalityOverflowLimit && !tagValues.Contains(trimmedTag.Value()) {
			trimmedTag = NewTag(trimmedTag.Name(), CardinalityOverflowValue)
			t.warningHandler(warnings.TagOverflow(metric, tag.Name(), tag.Value()))
		}

		t.trackable[metric][trimmedTag.Name()] = tagValues.Add(trimmedTag.Value())
		cleanedTags = append(cleanedTags, trimmedTag)
	}

	return cleanedTags
}

func (t *tracker) trimTag(metric string, tag Tag) Tag {
	tagName := tag.Name()
	tagValue := tag.Value()

	if len(tagName) > NameMaxLen {
		tagName = tagName[:NameMaxLen]
		t.warningHandler(warnings.TagNameTooLong(metric, tag.Name(), tag.Value()))
	}

	if len(tagValue) > ValueMaxLen {
		tagValue = tagValue[:ValueMaxLen]
		t.warningHandler(warnings.TagValueTooLong(metric, tag.Name(), tag.Value()))
	}

	return NewTag(tagName, tagValue)
}

var _ Tracker = NoopTracker{}

type NoopTracker struct{}

func (NoopTracker) CleanTags(metric string, tags []Tag) []Tag {
	return tags
}
