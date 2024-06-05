package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vnworkday/go-metrics/pkg/metrics"
	"github.com/vnworkday/go-metrics/pkg/tags"
)

type MockMetricProvider struct {
	mock.Mock
}

func (m *MockMetricProvider) RegisterMeter(name string, meter metrics.Meter, options ...metrics.InstrumentOptions) (metrics.Unregister, error) {
	args := m.Called(name, meter, options)
	return args.Get(0).(metrics.Unregister), args.Error(1)
}

func (m *MockMetricProvider) GetCounter(name string, options ...metrics.InstrumentOptions) (metrics.Counter, error) {
	args := m.Called(name, options)
	return args.Get(0).(metrics.Counter), args.Error(1)
}

func (m *MockMetricProvider) GetHistogram(name string, options ...metrics.InstrumentOptions) (metrics.Histogram, error) {
	args := m.Called(name, options)
	return args.Get(0).(metrics.Histogram), args.Error(1)
}

func (m *MockMetricProvider) GetUpDownCounter(name string, options ...metrics.InstrumentOptions) (metrics.UpDownCounter, error) {
	args := m.Called(name, options)
	return args.Get(0).(metrics.UpDownCounter), args.Error(1)
}

func (m *MockMetricProvider) GetGauge(name string, options ...metrics.InstrumentOptions) (metrics.Gauge, error) {
	args := m.Called(name, options)
	return args.Get(0).(metrics.Gauge), args.Error(1)
}

type MockCounter struct {
	mock.Mock
}

func NewMockCounter() *MockCounter {
	m := &MockCounter{}
	m.On("Add", mock.Anything, mock.Anything, mock.Anything)
	return m
}

func (m *MockCounter) Add(ctx context.Context, value uint, tags ...tags.Tag) {
	m.Called(ctx, value, tags)
}

type MockHistogram struct {
	mock.Mock
}

func NewMockHistogram() *MockHistogram {
	m := &MockHistogram{}
	m.On("Record", mock.Anything, mock.Anything, mock.Anything)
	return m
}

func (m *MockHistogram) Record(ctx context.Context, value int, tags ...tags.Tag) {
	m.Called(ctx, value, tags)
}

type MockUpDownCounter struct {
	mock.Mock
}

func NewMockUpDownCounter() *MockUpDownCounter {
	m := &MockUpDownCounter{}
	m.On("Add", mock.Anything, mock.Anything, mock.Anything)
	return m
}

func (m *MockUpDownCounter) Add(ctx context.Context, value uint, tags ...tags.Tag) {
	m.Called(ctx, value, tags)
}

type MockGauge struct {
	mock.Mock
}

func NewMockGauge() *MockGauge {
	m := &MockGauge{}
	m.On("Record", mock.Anything, mock.Anything, mock.Anything)
	return m
}

func (m *MockGauge) Record(ctx context.Context, value int, tags ...tags.Tag) {
	m.Called(ctx, value, tags)
}
