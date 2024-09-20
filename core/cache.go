package core

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type cacheClient struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewCacheClient(addr string) (c *cacheClient, err error) {
	c = &cacheClient{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "",
			DB:       0,
		}),
		ctx: context.Background(),
	}
	_, err = c.redisClient.Ping(c.ctx).Result()
	return
}

func (c *cacheClient) Set(key string, value any, expSecond int64) error {
	return c.redisClient.Set(c.ctx, key, value, time.Duration(expSecond)*time.Second).Err()
}

func (c *cacheClient) Get(key string) (string, error) {
	return c.redisClient.Get(c.ctx, key).Result()
}

func (c *cacheClient) TTL(key string) (time.Duration, error) {
	return c.redisClient.TTL(c.ctx, key).Result()
}
