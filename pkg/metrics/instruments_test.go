package metrics

import (
	"github.com/vnworkday/go-metrics/pkg/tags"
	"github.com/vnworkday/go-metrics/pkg/units"
	"testing"
)

func TestInstrumentOptions_WithUnit(t *testing.T) {
	tests := []struct {
		name string
		unit units.Unit
		want units.Unit
	}{
		{"WithUnit", "unit1", "unit1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			io := NewInstrumentOptions()
			if got := io.WithUnit(tt.unit).Unit(); got != tt.want {
				t.Errorf("InstrumentOptions.WithUnit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstrumentOptions_WithTags(t *testing.T) {
	tests := []struct {
		name string
		tags []tags.Tag
		want []tags.Tag
	}{
		{"WithTags", []tags.Tag{tags.NewTag("key", "value")}, []tags.Tag{tags.NewTag("key", "value")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			io := NewInstrumentOptions()
			if got := io.WithTags(tt.tags...).Tags(); len(got) != len(tt.want) {
				t.Errorf("InstrumentOptions.WithTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstrumentOptions_WithDesc(t *testing.T) {
	tests := []struct {
		name string
		desc string
		want string
	}{
		{"WithDesc", "description", "description"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			io := NewInstrumentOptions()
			if got := io.WithDesc(tt.desc).Desc(); got != tt.want {
				t.Errorf("InstrumentOptions.WithDesc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeInstrumentOptions(t *testing.T) {
	tests := []struct {
		name    string
		options []InstrumentOptions
		wantErr bool
	}{
		{
			"MergeSingleOption",
			[]InstrumentOptions{
				NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value")).WithDesc("description"),
			},
			false,
		},
		{
			"MergeMultipleOptionsWithDifferentUnits",
			[]InstrumentOptions{
				NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key1", "value1")).WithDesc("description1"),
				NewInstrumentOptions().WithUnit("unit2").WithTags(tags.NewTag("key2", "value2")).WithDesc("description2"),
			},
			true,
		},
		{
			"MergeMultipleOptionsWithSameUnits",
			[]InstrumentOptions{
				NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key1", "value1")).WithDesc("description1"),
				NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key2", "value2")).WithDesc("description2"),
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := MergeInstrumentOptions(tt.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MergeInstrumentOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
