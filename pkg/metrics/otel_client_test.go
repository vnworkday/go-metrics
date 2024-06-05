package metrics

import (
	"context"
	"github.com/vnworkday/go-metrics/pkg/tags"
	"github.com/vnworkday/go-metrics/pkg/units"
	"testing"
)

var c, _ = NewOtelClient(context.Background())

func TestOtelClient_RegisterMeter(t *testing.T) {
	tests := []struct {
		name       string
		metricName string
		options    []InstrumentOptions
		wantErr    bool
	}{
		{"ValidMeter", "testMeter", []InstrumentOptions{NewInstrumentOptions().WithUnit(units.Bytes).WithTags(tags.NewTag("key", "value")).WithDesc("description")}, false},
		{"WithoutNameShouldFail", "", []InstrumentOptions{NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value")).WithDesc("description")}, true},
		{"WithoutOptionsShouldOK", "testMeter", []InstrumentOptions{}, false},
		{"WithInvalidUnitInOptionsShouldFail", "testMeter", []InstrumentOptions{NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value")).WithDesc("description"), NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value")).WithDesc("description")}, true},
		{"WithDuplicatedTagsInOptionsShouldFail", "testMeter", []InstrumentOptions{NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value"), tags.NewTag("key", "value")).WithDesc("description")}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.RegisterMeter(tt.metricName, nil, tt.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("OtelClient.RegisterMeter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOtelClient_GetCounter(t *testing.T) {
	tests := []struct {
		name       string
		metricName string
		options    []InstrumentOptions
		wantErr    bool
	}{
		{"ValidCounter", "testCounter", []InstrumentOptions{NewInstrumentOptions().WithUnit(units.Millis).WithTags(tags.NewTag("key", "value")).WithDesc("description")}, false},
		{"WithoutNameShouldFail", "", []InstrumentOptions{NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value")).WithDesc("description")}, true},
		{"WithoutOptionsShouldOK", "testCounter", []InstrumentOptions{}, false},
		{"WithInvalidUnitInOptionsShouldFail", "testCounter", []InstrumentOptions{NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value")).WithDesc("description"), NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value")).WithDesc("description")}, true},
		{"WithDuplicatedTagsInOptionsShouldFail", "testCounter", []InstrumentOptions{NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value"), tags.NewTag("key", "value")).WithDesc("description")}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GetCounter(tt.metricName, tt.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("OtelClient.GetCounter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOtelClient_GetHistogram(t *testing.T) {
	tests := []struct {
		name       string
		metricName string
		options    []InstrumentOptions
		wantErr    bool
	}{
		{"ValidHistogram", "testHistogram", []InstrumentOptions{NewInstrumentOptions().WithUnit(units.Dimensionless).WithTags(tags.NewTag("key", "value")).WithDesc("description")}, false},
		{"WithoutNameShouldFail", "", []InstrumentOptions{NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value")).WithDesc("description")}, true},
		{"WithoutOptionsShouldOK", "testHistogram", []InstrumentOptions{}, false},
		{"WithInvalidUnitInOptionsShouldFail", "testHistogram", []InstrumentOptions{NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value")).WithDesc("description"), NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value")).WithDesc("description")}, true},
		{"WithDuplicatedTagsInOptionsShouldFail", "testHistogram", []InstrumentOptions{NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value"), tags.NewTag("key", "value")).WithDesc("description")}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GetHistogram(tt.metricName, tt.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("OtelClient.GetHistogram() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOtelClient_GetUpDownCounter(t *testing.T) {
	tests := []struct {
		name       string
		metricName string
		options    []InstrumentOptions
		wantErr    bool
	}{
		{"ValidUpDownCounter", "testUpDownCounter", []InstrumentOptions{NewInstrumentOptions().WithUnit(units.Bytes).WithTags(tags.NewTag("key", "value")).WithDesc("description")}, false},
		{"WithoutNameShouldFail", "", []InstrumentOptions{NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value")).WithDesc("description")}, true},
		{"WithoutOptionsShouldOK", "testUpDownCounter", []InstrumentOptions{}, false},
		{"WithInvalidUnitInOptionsShouldFail", "testUpDownCounter", []InstrumentOptions{NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value")).WithDesc("description"), NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value")).WithDesc("description")}, true},
		{"WithDuplicatedTagsInOptionsShouldFail", "testUpDownCounter", []InstrumentOptions{NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value"), tags.NewTag("key", "value")).WithDesc("description")}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GetUpDownCounter(tt.metricName, tt.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("OtelClient.GetUpDownCounter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOtelClient_GetGauge(t *testing.T) {
	tests := []struct {
		name       string
		metricName string
		options    []InstrumentOptions
		wantErr    bool
	}{
		{"ValidGauge", "testGauge", []InstrumentOptions{NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value")).WithDesc("description")}, false},
		{"InvalidGauge", "", []InstrumentOptions{NewInstrumentOptions().WithUnit("unit1").WithTags(tags.NewTag("key", "value")).WithDesc("description")}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GetGauge(tt.metricName, tt.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("OtelClient.GetGauge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
