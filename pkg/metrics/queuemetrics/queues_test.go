package queuemetrics

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/vnworkday/go-metrics/internal/mocks"
)

var mockClient = mocks.NewMockQueueMetricClient()

func TestDoMessage(t *testing.T) {
	okRequest := func() error {
		return nil
	}
	nokRequest := func() error {
		return errors.New("unexpected error")
	}

	tests := []struct {
		name         string
		metricClient Client
		queueName    string
		queueType    string
		makeMessage  func() error
		wantErr      bool
	}{
		{"Valid", mockClient, "valid", "valid", okRequest, false},
		{"WithNokRequestShouldFail", mockClient, "valid", "valid", nokRequest, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DoMessage(context.Background(), tt.metricClient, tt.queueName, tt.queueType, tt.makeMessage)

			if (err != nil) != tt.wantErr {
				t.Errorf("DoMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoMessageWithResponse(t *testing.T) {
	okRequest := func() (interface{}, error) {
		return "response", nil
	}
	nokRequest := func() (interface{}, error) {
		return nil, errors.New("unexpected error")
	}

	tests := []struct {
		name         string
		metricClient Client
		queueName    string
		queueType    string
		makeRequest  func() (interface{}, error)
		wantErr      bool
	}{
		{"Valid", mockClient, "valid", "valid", okRequest, false},
		{"WithNokRequestShouldFail", mockClient, "valid", "valid", nokRequest, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DoMessageWithResponse(context.Background(), tt.metricClient, tt.queueName, tt.queueType, tt.makeRequest)

			if (err != nil) != tt.wantErr {
				t.Errorf("DoMessageWithResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
