package queuemetrics

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/vnworkday/go-metrics/internal/mocks"
	"github.com/vnworkday/go-metrics/pkg/metrics"
	"github.com/vnworkday/go-metrics/pkg/tags"
)

func TestQueue_New(t *testing.T) {
	mockMetricProvider := new(mocks.MockMetricProvider)

	tests := []struct {
		name           string
		queue          string
		metricProvider metrics.MetricProvider
		options        []MetricOption
		wantErr        bool
	}{
		{"ValidQueue", "valid", mockMetricProvider, []MetricOption{WithMetricTags(tags.NewTag("key", "value"))}, false},
		{"WithoutQueueShouldFail", "", mockMetricProvider, nil, true},
		{"WithoutClientShouldFail", "valid", nil, nil, true},
		{"WithoutTagsShouldOK", "valid", mockMetricProvider, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.metricProvider != nil {
				c := tt.metricProvider.(*mocks.MockMetricProvider)
				c.On("GetHistogram", QueueMessageLatencyHistogramName, mock.Anything).Return(metrics.OtelHistogram{}, nil)
				c.On("GetCounter", QueueMessageCounterName, mock.Anything).Return(metrics.OtelCounter{}, nil)
			}

			_, err := New(tt.queue, tt.metricProvider, tt.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
