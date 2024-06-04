package tags

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vnworkday/go-metrics/pkg/warnings"
)

func TestNewTagCreatesTagWithGivenNameAndValue(t *testing.T) {
	assertThat := assert.New(t)
	tag := NewTag("name1", "value1")
	assertThat.Equal("name1", tag.Name())
	assertThat.Equal("value1", tag.Value())
}

func TestTagStringReturnsCorrectFormat(t *testing.T) {
	assertThat := assert.New(t)
	tag := NewTag("name1", "value1")
	assertThat.Equal("name1=value1", tag.String())
}

func TestDefaultTagCleanerCleansTagsCorrectly(t *testing.T) {
	assertThat := assert.New(t)
	cleaner := NewTagCleaner(nil)
	tags := []Tag{NewTag("name1", "value1"), NewTag("name2", "value2")}
	cleanedTags := cleaner.Clean("metric1", tags)
	assertThat.Equal(tags, cleanedTags)
}

func TestDefaultTagCleanerRedactsOverflowingTags(t *testing.T) {
	assertThat := assert.New(t)
	cleaner := NewTagCleaner(warnings.DefaultWarningHandler())
	tags := make([]Tag, TagMaxSize+1)
	for i := range tags {
		tags[i] = NewTag("name", fmt.Sprintf("value%d", i))
	}
	cleanedTags := cleaner.Clean("metric1", tags)
	assertThat.Equal(TagRedactedDueToOverflow, cleanedTags[TagMaxSize].Value())
}

func TestDefaultTagCleanerTrimsTagNamesAndValues(t *testing.T) {
	assertThat := assert.New(t)
	cleaner := NewTagCleaner(warnings.DefaultWarningHandler())
	longName := strings.Repeat("a", TagNameMaxLen+1)
	longValue := strings.Repeat("b", TagValueMaxLen+1)
	tags := []Tag{NewTag(longName, "value1"), NewTag("name2", longValue)}
	cleanedTags := cleaner.Clean("metric1", tags)
	assertThat.Equal(TagNameMaxLen, len(cleanedTags[0].Name()))
	assertThat.Equal(TagValueMaxLen, len(cleanedTags[1].Value()))
}

func TestNoopTagCleanerDoesNotCleanTags(t *testing.T) {
	assertThat := assert.New(t)
	cleaner := NoopTagCleaner{}
	longName := strings.Repeat("a", TagNameMaxLen+1)
	longValue := strings.Repeat("b", TagValueMaxLen+1)
	tags := []Tag{NewTag(longName, "value1"), NewTag("name2", longValue)}
	cleanedTags := cleaner.Clean("metric1", tags)
	assertThat.Equal(TagNameMaxLen+1, len(cleanedTags[0].Name()))
	assertThat.Equal(TagValueMaxLen+1, len(cleanedTags[1].Value()))
}
