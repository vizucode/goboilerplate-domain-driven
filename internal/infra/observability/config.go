package observability

type Config struct {
	ServiceName string
	OtelMode    string // stdout or otlp
	Endpoint    string // otel-collector:4317
}
