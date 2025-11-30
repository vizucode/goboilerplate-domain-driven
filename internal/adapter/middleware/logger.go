package middleware

import (
	"bytes"
	"goboilerplate-domain-driven/internal/infra/observability"
	"goboilerplate-domain-driven/pkg/utils"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type responseRecorder struct {
	http.ResponseWriter
	Body       *bytes.Buffer
	StatusCode uint
}

func (r responseRecorder) Write(b []byte) (int, error) {
	r.Body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r *responseRecorder) WriteHeader(code int) {
	r.StatusCode = uint(code)
	r.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var span trace.Span

		tracer := otel.Tracer("Logger")
		ctx, span = tracer.Start(ctx, "Logger:Init")
		defer span.End()

		// 1. Capture response
		rec := &responseRecorder{
			ResponseWriter: w,
			Body:           bytes.NewBuffer([]byte{}),
			StatusCode:     200,
		}

		// 2. Prepare log container
		lc := &observability.LogContainer{
			RequestID:    uuid.New().String(),
			HttpMethod:   r.Method,
			Headers:      map[string]string{},
			ThirdParties: []observability.ThirdPartyLog{},
		}

		// Extract headers
		for k, v := range r.Header {
			if len(v) > 0 {
				lc.Headers[k] = v[0]
			}
		}

		// Extract request body
		bodyBytes, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		lc.BodyRequest = string(bodyBytes)

		// Put into context
		ctx = utils.SetLogContainer(ctx, lc)

		// 3. Run request
		next.ServeHTTP(rec, r.WithContext(ctx))

		// 4. After handler completed
		lc.Response = rec.Body.String()
		lc.StatusCode = rec.StatusCode

		// 5. Log as JSON via zerolog
		log.Info().
			Str("request_id", lc.RequestID).
			Uint("status_code", lc.StatusCode).
			Str("trace_id", span.SpanContext().TraceID().String()).
			Str("span_id", span.SpanContext().SpanID().String()).
			Str("user_identifier", lc.UserIdentifier).
			Str("http_method", lc.HttpMethod).
			Interface("headers", lc.Headers).
			Interface("third_parties", lc.ThirdParties).
			Str("body_request", lc.BodyRequest).
			Interface("logs", lc.Logs).
			Str("response", lc.Response).
			Msg("http_request_log")
	})
}
