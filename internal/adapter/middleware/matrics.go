package middleware

import (
	"goboilerplate-domain-driven/internal/infra/observability"
	"net/http"
	"time"
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

		observability.HttpRequestCounter.Add(ctx, 1)

		// Wrap writer
		statusWriter := &StatusResponseWriter{ResponseWriter: w, Status: 200}
		mrw := &MetricsResponseWriter{StatusResponseWriter: statusWriter}

		// Execute handler
		next.ServeHTTP(mrw, r)

		latency := float64(time.Since(start).Milliseconds())
		observability.HttpLatency.Record(ctx, latency)

		// Record request size (Content-Length header if present)
		var reqSize int64 = 0
		if r.ContentLength > 0 {
			reqSize = r.ContentLength
		}
		observability.HttpRequestSize.Record(r.Context(), reqSize)

		// Record response size
		observability.HttpResponseSize.Record(r.Context(), mrw.BytesWritten)

		// Hitung berdasarkan status code
		switch {
		case mrw.Status >= 200 && mrw.Status < 300:
			observability.HttpStatus2xx.Add(r.Context(), 1)

		case mrw.Status >= 400 && mrw.Status < 500:
			observability.HttpStatus4xx.Add(r.Context(), 1)

		case mrw.Status >= 500:
			observability.HttpStatus5xx.Add(r.Context(), 1)
		}
	})
}
