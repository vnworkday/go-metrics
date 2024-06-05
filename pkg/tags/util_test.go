package tags

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/attribute"
)

func TestToAttribute(t *testing.T) {
	assertThat := assert.New(t)

	tests := []struct {
		name     string
		tag      Tag
		expected attribute.KeyValue
	}{
		{
			name:     "ToAttributeConversion",
			tag:      NewTag("name1", "value1"),
			expected: attribute.Key("name1").String("value1"),
		},
		{
			name:     "ToAttributeConversionWithEmptyValue",
			tag:      NewTag("name2", ""),
			expected: attribute.Key("name2").String(""),
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr := ToAttribute(tt.tag)
			assertThat.Equal(tt.expected, attr)
		})
	}
}

func TestFromAttribute(t *testing.T) {
	tests := []struct {
		name     string
		attr     attribute.KeyValue
		expected Tag
	}{
		{
			name:     "FromAttributeConversion",
			attr:     attribute.Key("name1").String("value1"),
			expected: NewTag("name1", "value1"),
		},
		{
			name:     "FromAttributeConversionWithEmptyValue",
			attr:     attribute.Key("name2").String(""),
			expected: NewTag("name2", ""),
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertThat := assert.New(t)
			tag := FromAttribute(tt.attr)
			assertThat.Equal(tt.expected, tag)
		})
	}
}

func TestToAttributes(t *testing.T) {
	tests := []struct {
		name     string
		tags     []Tag
		expected []attribute.KeyValue
	}{
		{
			name:     "ToAttributesConversion",
			tags:     []Tag{NewTag("name1", "value1"), NewTag("name2", "value2")},
			expected: []attribute.KeyValue{attribute.Key("name1").String("value1"), attribute.Key("name2").String("value2")},
		},
		{
			name:     "ToAttributesConversionWithEmptyTags",
			tags:     []Tag{},
			expected: []attribute.KeyValue{},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertThat := assert.New(t)
			attrs := ToAttributes(tt.tags)
			assertThat.Equal(tt.expected, attrs)
		})
	}
}

func TestToTagsConversion(t *testing.T) {
	tests := []struct {
		name     string
		attrs    []attribute.KeyValue
		expected []Tag
	}{
		{
			name:     "ToTagsConversion",
			attrs:    []attribute.KeyValue{attribute.Key("name1").String("value1"), attribute.Key("name2").String("value2")},
			expected: []Tag{NewTag("name1", "value1"), NewTag("name2", "value2")},
		},
		{
			name:     "ToTagsConversionWithEmptyAttributes",
			attrs:    []attribute.KeyValue{},
			expected: []Tag{},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertThat := assert.New(t)
			tags := ToTags(tt.attrs)
			assertThat.Equal(tt.expected, tags)
		})
	}
}

func TestAddTags(t *testing.T) {
	tests := []struct {
		name     string
		attrs    []attribute.KeyValue
		tags     []Tag
		expected []attribute.KeyValue
	}{
		{
			name:     "AddTagsAppendsCorrectly",
			attrs:    []attribute.KeyValue{attribute.Key("name1").String("value1")},
			tags:     []Tag{NewTag("name2", "value2"), NewTag("name3", "value3")},
			expected: []attribute.KeyValue{attribute.Key("name1").String("value1"), attribute.Key("name2").String("value2"), attribute.Key("name3").String("value3")},
		},
		{
			name:     "AddTagsAppendsCorrectlyWithEmptyAttributes",
			attrs:    []attribute.KeyValue{},
			tags:     []Tag{NewTag("name2", "value2"), NewTag("name3", "value3")},
			expected: []attribute.KeyValue{attribute.Key("name2").String("value2"), attribute.Key("name3").String("value3")},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertThat := assert.New(t)
			attrs := AddTags(tt.attrs, tt.tags...)
			assertThat.Equal(tt.expected, attrs)
		})
	}
}

func TestTrimTag(t *testing.T) {
	tests := []struct {
		name     string
		tag      Tag
		expected Tag
	}{
		{
			name:     "TrimTagTrimsCorrectly",
			tag:      NewTag("name1", "value1"),
			expected: NewTag("name1", "value1"),
		},
		{
			name:     "TrimTagTrimsCorrectlyWithLongName",
			tag:      NewTag("longName"+strings.Repeat(".", 200), "value1"),
			expected: NewTag("longName"+strings.Repeat(".", 200-len("longName")), "value1"),
		},
		{
			name:     "TrimTagTrimsCorrectlyWithLongValue",
			tag:      NewTag("name1", "longValue"+strings.Repeat(".", 200)),
			expected: NewTag("name1", "longValue"+strings.Repeat(".", 200-len("longValue"))),
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertThat := assert.New(t)
			trimmedTag, _, _ := TrimTag(tt.tag)
			assertThat.Equal(tt.expected, trimmedTag)
		})
	}
}
