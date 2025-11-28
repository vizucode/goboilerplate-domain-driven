package utils

import (
	"context"
	"fmt"
	"goboilerplate-domain-driven/internal/infra/observability"
	"runtime"

	"github.com/rs/zerolog"
)

type key string

var logContainerKey key = "logContainer"

func SetLogContainer(ctx context.Context, lc *observability.LogContainer) context.Context {
	return context.WithValue(ctx, logContainerKey, lc)
}

func GetLogContainer(ctx context.Context) *observability.LogContainer {
	if v := ctx.Value(logContainerKey); v != nil {
		return v.(*observability.LogContainer)
	}
	return nil
}

// helper
func AddLogDebug(ctx context.Context, message string) {
	_, file, line, _ := runtime.Caller(1) // 1 = caller of this function

	path := file + ":" + fmt.Sprint(line)
	lc := GetLogContainer(ctx)
	if lc != nil {
		lc.AddLog(zerolog.DebugLevel.String(), message, path)
	}
}

func AddLogError(ctx context.Context, message string) {
	_, file, line, _ := runtime.Caller(1) // 1 = caller of this function

	path := file + ":" + fmt.Sprint(line)

	lc := GetLogContainer(ctx)
	if lc != nil {
		lc.AddLog(zerolog.ErrorLevel.String(), message, path)
	}
}
