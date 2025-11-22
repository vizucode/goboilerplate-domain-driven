package infra

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	rdc *redis.Client
}

func NewRedis(cfg CacheCfg) Cache {
	rdc := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &redisClient{
		rdc,
	}
}

func (rd *redisClient) SetCtx(ctx context.Context, key string, value any, exp time.Duration) (err error) {
	if err = rd.rdc.Set(ctx, key, value, exp).Err(); err != nil {
		return err
	}
	return nil
}

func (rd *redisClient) GetCtx(ctx context.Context, key string) (resp string, err error) {
	resultCmd := rd.rdc.Get(ctx, key)
	if resultCmd.Err() != nil {
		return resp, resultCmd.Err()
	}

	return resultCmd.String(), nil
}

func (rd *redisClient) DelCtx(ctx context.Context, key string) (err error) {
	if err = rd.rdc.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
