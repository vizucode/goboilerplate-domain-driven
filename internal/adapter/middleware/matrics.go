package middleware

import (
	"goboilerplate-domain-driven/internal/infra/observability"
	"net/http"
)

type StatusResponseWriter struct {
	http.ResponseWriter
	Status int
}

func (w *StatusResponseWriter) WriteHeader(statusCode int) {
	w.Status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Wrap writer
		srw := &StatusResponseWriter{ResponseWriter: w, Status: 200}

		// Execute handler
		next.ServeHTTP(srw, r)

		// Hitung berdasarkan status code
		switch {
		case srw.Status >= 200 && srw.Status < 300:
			observability.HttpStatus2xx.Add(r.Context(), 1)

		case srw.Status >= 400 && srw.Status < 500:
			observability.HttpStatus4xx.Add(r.Context(), 1)

		case srw.Status >= 500:
			observability.HttpStatus5xx.Add(r.Context(), 1)
		}
	})
}
