package observability

import (
	"sync"
)

type ThirdPartyLog struct {
	URL          string            `json:"url"`
	Method       string            `json:"method"`
	StatusCode   int               `json:"status_code"`
	Headers      map[string]string `json:"headers"`
	RequestBody  string            `json:"request_body"`
	ResponseBody string            `json:"response_body"`
}

type LogRecord struct {
	Path    string `json:"path"`
	Mode    string `json:"mode"`
	Message string `json:"message"`
}

type LogContainer struct {
	TraceId        string            `json:"trace_id"`
	SpanId         string            `json:"span_id"`
	ServiceName    string            `json:"service_name"`
	StatusCode     uint              `json:"status_code"`
	RequestID      string            `json:"request_id"`
	UserIdentifier string            `json:"user_identifier"`
	HttpMethod     string            `json:"http_method"`
	Headers        map[string]string `json:"headers"`
	ThirdParties   []ThirdPartyLog   `json:"third_parties"`
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
