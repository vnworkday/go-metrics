package apimetrics

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/vnworkday/go-metrics/internal/mocks"
)

var mockClient = mocks.NewMockAPIMetricClients()

func TestDoRequest(t *testing.T) {
	okRequest := func() error {
		return nil
	}
	nokRequest := func() error {
		return errors.New("unexpected error")
	}

	tests := []struct {
		name        string
		metric      Client
		opName      string
		makeRequest func() error
		wantErr     bool
	}{
		{"Valid", mockClient, "valid", okRequest, false},
		{"WithNokRequestShouldFail", mockClient, "valid", nokRequest, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DoRequest(context.Background(), tt.metric, tt.opName, tt.makeRequest)

			if (err != nil) != tt.wantErr {
				t.Errorf("DoRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoRequestWithResponse(t *testing.T) {
	okRequest := func() (interface{}, error) {
		return "response", nil
	}
	nokRequest := func() (interface{}, error) {
		return nil, errors.New("unexpected error")
	}

	tests := []struct {
		name        string
		metric      Client
		opName      string
		makeRequest func() (interface{}, error)
		wantErr     bool
	}{
		{"Valid", mockClient, "valid", okRequest, false},
		{"WithNokRequestShouldFail", mockClient, "valid", nokRequest, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DoRequestWithResponse(context.Background(), tt.metric, tt.opName, tt.makeRequest)

			if (err != nil) != tt.wantErr {
				t.Errorf("DoRequestWithResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
