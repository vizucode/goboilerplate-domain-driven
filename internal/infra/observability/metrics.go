package observability

import (
	"context"
	"log"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

var (
	Meter metric.Meter

	// 4 jenis metrics (plus additional golden metrics)
	HttpRequestCounter metric.Int64Counter
	ActiveUsers        metric.Int64UpDownCounter
	HttpLatency        metric.Float64Histogram
	HttpRequestSize    metric.Int64Histogram
	HttpResponseSize   metric.Int64Histogram

	HttpStatus2xx metric.Int64Counter
	HttpStatus4xx metric.Int64Counter
	HttpStatus5xx metric.Int64Counter
)

// Gauge variable (queue length / dynamic value)
var queueLength int64 = 0

func InitMetrics(ctx context.Context, cfg Config) func() {

	var (
		exp sdkmetric.Exporter
		err error
	)

	if strings.EqualFold(cfg.OtelMode, "otlp") {
		exp, err = otlpmetricgrpc.New(
			ctx,
			otlpmetricgrpc.WithEndpoint(cfg.Endpoint),
			otlpmetricgrpc.WithInsecure(),
		)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		exp, err = stdoutmetric.New(stdoutmetric.WithPrettyPrint())
		if err != nil {
			log.Fatal(err)
		}
	}

	// --- 2. Register provider ---
	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exp)),
		sdkmetric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(cfg.ServiceName),
		)),
	)

	otel.SetMeterProvider(provider)

	Meter = provider.Meter(cfg.ServiceName)

	// --- 3. Init semua instruments ---
	initCounters()
	initUpDownCounters()
	initHistograms()
	initGauges()

	// --- fungsi shutdown ---
	return func() {
		_ = provider.Shutdown(ctx)
	}
}

func initCounters() {
	HttpRequestCounter, _ = Meter.Int64Counter("http_requests_total")
	HttpStatus2xx, _ = Meter.Int64Counter("http_response_status_2xx_total")
	HttpStatus4xx, _ = Meter.Int64Counter("http_response_status_4xx_total")
	HttpStatus5xx, _ = Meter.Int64Counter("http_response_status_5xx_total")
}

func initUpDownCounters() {
	ActiveUsers, _ = Meter.Int64UpDownCounter("active_users")
}

func initHistograms() {
	HttpLatency, _ = Meter.Float64Histogram("http_request_latency_ms")
	HttpRequestSize, _ = Meter.Int64Histogram("http_request_size_bytes")
	HttpResponseSize, _ = Meter.Int64Histogram("http_response_size_bytes")
}

func initGauges() {
	// Observable Gauge
	Meter.Int64ObservableGauge(
		"queue_length",
		metric.WithInt64Callback(func(ctx context.Context, o metric.Int64Observer) error {
			o.Observe(queueLength) // membaca value saat ini
			return nil
		}),
	)
}

// Helper untuk update queue
func IncrementQueue() { queueLength++ }
func DecrementQueue() { queueLength-- }
