package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/vnworkday/go-metrics/pkg/metrics"
)

type MockAPIMetricClients struct {
	mock.Mock
}

func NewMockAPIMetricClients() *MockAPIMetricClients {
	m := &MockAPIMetricClients{}
	m.On("GetRequestCounter").Return(NewMockCounter())
	m.On("GetLatencyHistogram").Return(NewMockHistogram())
	return m
}

func (m *MockAPIMetricClients) GetLatencyHistogram() metrics.Histogram {
	args := m.Called()
	return args.Get(0).(metrics.Histogram)
}

func (m *MockAPIMetricClients) GetRequestCounter() metrics.Counter {
	args := m.Called()
	return args.Get(0).(metrics.Counter)
}

type MockQueueMetricClients struct {
	mock.Mock
}

func (m *MockQueueMetricClients) GetLatencyHistogram() metrics.Histogram {
	args := m.Called()
	return args.Get(0).(metrics.Histogram)
}

func (m *MockQueueMetricClients) GetMessageCounter() metrics.Counter {
	args := m.Called()
	return args.Get(0).(metrics.Counter)
}
