package observability

import (
	"sync"
)

type LogRecord struct {
	Path    string `json:"path"`
	Mode    string `json:"mode"`
	Message string `json:"message"`
}

type LogContainer struct {
	TraceId        string            `json:"trace_id"`
	SpanId         string            `json:"span_id"`
	ServiceName    string            `json:"service_name"`
	RequestID      string            `json:"request_id"`
	UserIdentifier string            `json:"user_identifier"`
	HttpMethod     string            `json:"http_method"`
	Headers        map[string]string `json:"headers"`
	BodyRequest    string            `json:"body_request"`
	Logs           []LogRecord       `json:"logs"`
	Response       string            `json:"response"`

	mu sync.Mutex
}

func (lc *LogContainer) AddTraceLog(traceId, spanId string) {
	lc.TraceId = traceId
	lc.SpanId = spanId
}

func (lc *LogContainer) AddLog(level, msg, path string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.Logs = append(lc.Logs, LogRecord{
		Path:    path,
		Mode:    level,
		Message: msg,
	})
}
