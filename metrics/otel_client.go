package metrics

import (
	"context"
	"fmt"
	"time"

	"github.com/vnworkday/go-metrics/tags"
	"github.com/vnworkday/go-metrics/units"

	"github.com/pkg/errors"
	"github.com/vnworkday/go-metrics/warnings"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	otelmetric "go.opentelemetry.io/otel/sdk/metric"
)

type OtelClient struct {
	config         Config
	meter          metric.Meter
	attrs          []attribute.KeyValue
	warningHandler warnings.WarningHandler
	tagCleaner     tags.TagCleaner
}

func (c OtelClient) RegisterMeter(metricName string, meter Meter, options ...InstrumentOptions) (Unregister, error) {
	mergeOption, err := MergeInstrumentOptions(options...)

	if err != nil {
		return nil, err
	}

	otelUnit, err := units.ToOtelUnit(metricName, mergeOption.Unit())

	if err != nil {
		c.warningHandler(warnings.UnitInvalid(metricName, string(mergeOption.Unit())))
	}

	gauge, err := c.meter.Int64ObservableGauge(
		metricName,
		metric.WithDescription(mergeOption.Desc()),
		metric.WithUnit(otelUnit),
	)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to create observable gauge for metric %s", metricName)
	}

	cleanedTags := c.tagCleaner.Clean(metricName, tags.ToTags(tags.AddTags(c.attrs, mergeOption.Tags()...)))

	registration, err := c.meter.RegisterCallback(func(ctx context.Context, observer metric.Observer) error {
		meteredValue := meter()
		observer.ObserveInt64(gauge, int64(meteredValue), metric.WithAttributes(tags.ToAttributes(cleanedTags)...))
		return nil
	}, gauge)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to register callback for metric %s", metricName)
	}

	return func() error {
		return errors.Wrapf(registration.Unregister(), "failed to unregister callback for metric %s", metricName)
	}, nil
}

func (c OtelClient) GetCounter(metricName string, options ...InstrumentOptions) (Counter, error) {
	mergedOptions, err := MergeInstrumentOptions(options...)

	if err != nil {
		return nil, err
	}

	otelUnit, err := units.ToOtelUnit(metricName, mergedOptions.Unit())

	if err != nil {
		c.warningHandler(warnings.UnitInvalid(metricName, string(mergedOptions.Unit())))
	}

	counter, err := c.meter.Int64Counter(
		metricName,
		metric.WithUnit(otelUnit),
		metric.WithDescription(mergedOptions.Desc()),
	)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to create counter for metric %s", metricName)
	}

	return newOtelCounter(metricName, counter, c.tagCleaner, tags.AddTags(c.attrs, mergedOptions.Tags()...)...), nil
}

func (c OtelClient) GetHistogram(metricName string, options ...InstrumentOptions) (Histogram, error) {
	mergedOptions, err := MergeInstrumentOptions(options...)

	if err != nil {
		return nil, err
	}

	otelUnit, err := units.ToOtelUnit(metricName, mergedOptions.Unit())

	if err != nil {
		c.warningHandler(warnings.UnitInvalid(metricName, string(mergedOptions.Unit())))
	}

	histogram, err := c.meter.Int64Histogram(
		metricName,
		metric.WithUnit(otelUnit),
		metric.WithDescription(mergedOptions.Desc()),
	)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to create histogram for metric %s", metricName)
	}

	return newOtelHistogram(metricName, histogram, c.tagCleaner, tags.AddTags(c.attrs, mergedOptions.Tags()...)...), nil
}

func (c OtelClient) GetUpDownCounter(metricName string, options ...InstrumentOptions) (UpDownCounter, error) {
	mergedOptions, err := MergeInstrumentOptions(options...)

	if err != nil {
		return nil, err
	}

	otelUnit, err := units.ToOtelUnit(metricName, mergedOptions.Unit())

	if err != nil {
		c.warningHandler(warnings.UnitInvalid(metricName, string(mergedOptions.Unit())))
	}

	upDownCounter, err := c.meter.Int64UpDownCounter(
		metricName,
		metric.WithUnit(otelUnit),
		metric.WithDescription(mergedOptions.Desc()),
	)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to create up down counter for metric %s", metricName)
	}

	return newOtelUpDownCounter(metricName, upDownCounter, c.tagCleaner, tags.AddTags(c.attrs, mergedOptions.Tags()...)...), nil
}

func (c OtelClient) GetGauge(metricName string, options ...InstrumentOptions) (Gauge, error) {
	mergedOptions, err := MergeInstrumentOptions(options...)

	if err != nil {
		return nil, err
	}

	otelUnit, err := units.ToOtelUnit(metricName, mergedOptions.Unit())

	if err != nil {
		c.warningHandler(warnings.UnitInvalid(metricName, string(mergedOptions.Unit())))
	}

	gauge, err := c.meter.Int64Gauge(
		metricName,
		metric.WithUnit(otelUnit),
		metric.WithDescription(mergedOptions.Desc()),
	)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to create gauge for metric %s", metricName)
	}

	return newOtelGauge(metricName, gauge, c.tagCleaner, tags.AddTags(c.attrs, mergedOptions.Tags()...)...), nil
}

type OtelClientOption = func(*OtelClient)

// Noop causes all functions to be no-ops and return nil errors.
func Noop() OtelClientOption {
	return func(c *OtelClient) {
		c.meter = noop.Meter{}
	}
}

func NewOtelClient(ctx context.Context, opts ...OtelClientOption) (OtelClient, error) {
	c := OtelClient{
		config:         DefaultConfig,
		warningHandler: warnings.DefaultWarningHandler(),
	}
	for _, opt := range opts {
		opt(&c)
	}

	if c.meter == nil {
		meter, err := defaultMeter(ctx, c.config)

		if err != nil {
			return c, err
		}

		c.meter = meter
	}

	if c.tagCleaner == nil {
		c.tagCleaner = tags.NewTagCleaner(c.warningHandler)
	}

	return c, nil
}

func defaultMeter(ctx context.Context, config Config) (metric.Meter, error) {
	if !config.isValid() {
		config = DefaultConfig
	}

	return newAdhocMeter(ctx, config.Host, config.Port)
}

func newAdhocMeter(ctx context.Context, host string, port int) (metric.Meter, error) {
	endpoint := fmt.Sprintf("%s:%d", host, port)

	grpcMetricExporter, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithEndpointURL(endpoint),
		otlpmetricgrpc.WithInsecure(),
	)

	if err != nil {
		return nil, errors.Wrap(err, "failed to create metric exporter")
	}

	meterProvider := otelmetric.NewMeterProvider(
		otelmetric.WithReader(
			otelmetric.NewPeriodicReader(
				grpcMetricExporter,
				otelmetric.WithInterval(10*time.Second),
			),
		),
	)

	return meterProvider.Meter("adhoc_meter"), nil
}

func WithConfig(config Config) OtelClientOption {
	return func(c *OtelClient) {
		c.config = config
		c.attrs = append(c.attrs, tags.ToAttributes(config.tags())...)
	}
}
