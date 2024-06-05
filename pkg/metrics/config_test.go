package metrics

import (
	"reflect"
	"testing"

	"github.com/vnworkday/go-metrics/pkg/tags"
)

func TestConfigIsValid(t *testing.T) {
	tests := []struct {
		name string
		c    Config
		want bool
	}{
		{"ValidConfig", Config{Host: "localhost", Port: 4317}, true},
		{"EmptyHost", Config{Host: "", Port: 4317}, false},
		{"ZeroPort", Config{Host: "localhost", Port: 0}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.isValid(); got != tt.want {
				t.Errorf("Config.isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigTags(t *testing.T) {
	tests := []struct {
		name string
		c    Config
		want []tags.Tag
	}{
		{
			"ValidTags",
			Config{Tags: []ConfigTag{{Key: "key1", Value: "value1"}, {Key: "key2", Value: "value2"}}},
			[]tags.Tag{tags.NewTag("key1", "value1"), tags.NewTag("key2", "value2")},
		},
		{
			"EmptyKey",
			Config{Tags: []ConfigTag{{Key: "", Value: "value1"}}},
			[]tags.Tag{},
		},
		{
			"EmptyValue",
			Config{Tags: []ConfigTag{{Key: "key1", Value: ""}}},
			[]tags.Tag{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.tags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.tags() = %v, want %v", got, tt.want)
			}
		})
	}
}
