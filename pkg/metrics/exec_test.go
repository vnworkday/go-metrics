package metrics

import (
	"github.com/pkg/errors"
	"github.com/vnworkday/go-metrics/pkg/statuses"
	"github.com/vnworkday/go-metrics/pkg/tags"
	"reflect"
	"testing"
)

func TestWithErrTypeConverter(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"NilError", nil, ""},
		{"GenericError", errors.New("generic error"), "generic error"},
	}

	var errConverter = func(err error) string {
		if err == nil {
			return ""
		} else {
			return err.Error()
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := NewExecParameters[any]()
			WithErrTypeConverter[any](errConverter)(&params)
			if got := params.ErrTypeConverter(tt.err); got != tt.want {
				t.Errorf("WithErrTypeConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithStatusConverter(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want statuses.Status
	}{
		{"NoError", nil, statuses.Success},
		{"WithError", errors.New("test error"), statuses.ClientError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := NewExecParameters[any]()
			WithStatusConverter[any](statuses.DefaultConverterWithoutResponse)(&params)
			if got := params.StatusConverter(nil, tt.err); got != tt.want {
				t.Errorf("WithStatusConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithExecTags(t *testing.T) {
	tests := []struct {
		name string
		tags []tags.Tag
		want []tags.Tag
	}{
		{"NoTags", []tags.Tag{}, []tags.Tag{}},
		{"WithTags", []tags.Tag{tags.NewTag("key", "value")}, []tags.Tag{tags.NewTag("key", "value")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := NewExecParameters[any]()
			WithExecTags[any](tt.tags...)(&params)
			if !reflect.DeepEqual(params.Tags, tt.want) {
				t.Errorf("WithExecTags() = %v, want %v", params.Tags, tt.want)
			}
		})
	}
}
