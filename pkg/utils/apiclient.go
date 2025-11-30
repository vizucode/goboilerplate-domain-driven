package utils

import (
	"bytes"
	"context"
	"goboilerplate-domain-driven/internal/infra/observability"
	"io"
	"net/http"
)

type LoggedHTTPClient struct {
	Client *http.Client
	Logs   *[]observability.ThirdPartyLog
}

func newClient(c *http.Client, logs *[]observability.ThirdPartyLog) *LoggedHTTPClient {
	return &LoggedHTTPClient{
		Client: c,
		Logs:   logs,
	}
}

func (l *LoggedHTTPClient) Do(req *http.Request) (*http.Response, error) {
	// --- Capture request body ---
	var reqBody string
	if req.Body != nil {
		bodyBytes, _ := io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		reqBody = string(bodyBytes)
	}

	// --- Execute real request ---
	resp, err := l.Client.Do(req)
	if err != nil {
		*l.Logs = append(*l.Logs, observability.ThirdPartyLog{
			URL:          req.URL.String(),
			Method:       req.Method,
			StatusCode:   0,
			Headers:      map[string]string{},
			RequestBody:  reqBody,
			ResponseBody: err.Error(),
		})
		return resp, err
	}

	// --- Capture response body ---
	var respBody string
	if resp.Body != nil {
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		respBody = string(bodyBytes)
	}

	// --- Capture response headers ---
	respHeaders := map[string]string{}
	for k, v := range resp.Header {
		if len(v) > 0 {
			respHeaders[k] = v[0]
		}
	}

	// --- Save log ---
	*l.Logs = append(*l.Logs, observability.ThirdPartyLog{
		URL:          req.URL.String(),
		Method:       req.Method,
		StatusCode:   resp.StatusCode,
		Headers:      respHeaders,
		RequestBody:  reqBody,
		ResponseBody: respBody,
	})

	return resp, nil
}

func NewClient(ctx context.Context, client *http.Client) *LoggedHTTPClient {
	lc := GetLogContainer(ctx)
	return newClient(client, &lc.ThirdParties)
}
