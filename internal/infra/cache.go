package infra

import (
	"context"
	"time"
)

type CacheCfg struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type Cache interface {
	SetCtx(ctx context.Context, key string, value any, exp time.Duration) (err error)
	GetCtx(ctx context.Context, key string) (resp string, err error)
	DelCtx(ctx context.Context, key string) (err error)
}
