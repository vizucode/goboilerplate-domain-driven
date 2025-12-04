package middleware

import (
	"goboilerplate-domain-driven/internal/infra/observability"
	"net/http"
	"strconv"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type StatusResponseWriter struct {
	http.ResponseWriter
	Status int
}

func (w *StatusResponseWriter) WriteHeader(statusCode int) {
	w.Status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

type MetricsResponseWriter struct {
	*StatusResponseWriter
	BytesWritten int64
}

func (w *MetricsResponseWriter) Write(p []byte) (int, error) {
	n, err := w.StatusResponseWriter.ResponseWriter.Write(p)
	w.BytesWritten += int64(n)
	return n, err
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		start := time.Now()

		// Wrap writer
		statusWriter := &StatusResponseWriter{ResponseWriter: w, Status: 200}
		mrw := &MetricsResponseWriter{StatusResponseWriter: statusWriter}

		// Execute handler
		next.ServeHTTP(mrw, r)

		attrs := []attribute.KeyValue{
			attribute.String("http.method", r.Method),
			attribute.String("http.route", r.URL.Path),
			attribute.String("http.status_code", strconv.Itoa(mrw.Status)),
		}

		observability.HttpRequestCounter.Add(ctx, 1, metric.WithAttributes(attrs...))

		latency := float64(time.Since(start).Milliseconds())
		observability.HttpLatency.Record(ctx, latency, metric.WithAttributes(attrs...))

		// Record request size (Content-Length header if present)
		var reqSize int64 = 0
		if r.ContentLength > 0 {
			reqSize = r.ContentLength
		}
		observability.HttpRequestSize.Record(r.Context(), reqSize, metric.WithAttributes(attrs...))

		// Record response size
		observability.HttpResponseSize.Record(r.Context(), mrw.BytesWritten, metric.WithAttributes(attrs...))

		// Hitung berdasarkan status code
		switch {
		case mrw.Status >= 200 && mrw.Status < 300:
			observability.HttpStatus2xx.Add(r.Context(), 1, metric.WithAttributes(attrs...))

		case mrw.Status >= 400 && mrw.Status < 500:
			observability.HttpStatus4xx.Add(r.Context(), 1, metric.WithAttributes(attrs...))

		case mrw.Status >= 500:
			observability.HttpStatus5xx.Add(r.Context(), 1, metric.WithAttributes(attrs...))
		}
	})
}
