package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/vnworkday/go-metrics/pkg/metrics"
)

type MockAPIMetricClient struct {
	mock.Mock
}

func NewMockAPIMetricClient() *MockAPIMetricClient {
	m := &MockAPIMetricClient{}
	m.On("GetRequestCounter").Return(NewMockCounter())
	m.On("GetLatencyHistogram").Return(NewMockHistogram())
	return m
}

func (m *MockAPIMetricClient) GetLatencyHistogram() metrics.Histogram {
	args := m.Called()
	return args.Get(0).(metrics.Histogram)
}

func (m *MockAPIMetricClient) GetRequestCounter() metrics.Counter {
	args := m.Called()
	return args.Get(0).(metrics.Counter)
}

type MockQueueMetricClient struct {
	mock.Mock
}

func NewMockQueueMetricClient() *MockQueueMetricClient {
	m := &MockQueueMetricClient{}
	m.On("GetLatencyHistogram").Return(NewMockHistogram())
	m.On("GetMessageCounter").Return(NewMockCounter())
	return m
}

func (m *MockQueueMetricClient) GetLatencyHistogram() metrics.Histogram {
	args := m.Called()
	return args.Get(0).(metrics.Histogram)
}

func (m *MockQueueMetricClient) GetMessageCounter() metrics.Counter {
	args := m.Called()
	return args.Get(0).(metrics.Counter)
}
