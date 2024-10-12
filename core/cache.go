package core

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	localCache "github.com/moonprism/blog/core/pkg/cache"
)

type Cache interface {
	// 操作成功返回true
	Set(key string, value any, ttl time.Duration) bool

	// Get 根据key获取值，key不存在返回nil
	Get(key string) any

	// TTL 无过期时间返回-1，key不存在返回 -2
	TTL(key string) time.Duration

	Del(key string) bool
}

func NewCache(addr string) (cache Cache, err error) {
	if addr == "local" {
		cache = localCache.New()
	} else {
		c := &redisCache{
			rdb: redis.NewClient(&redis.Options{
				Addr:     addr,
				Password: "",
				DB:       0,
			}),
			ctx: context.Background(),
		}
		_, err = c.rdb.Ping(c.ctx).Result()
		cache = c
	}
	return cache, err
}

type redisCache struct {
	rdb *redis.Client
	ctx context.Context
}

func (c *redisCache) Set(key string, value any, ttl time.Duration) bool {
	return c.rdb.Set(c.ctx, key, value, ttl).Err() == nil
}

func (c *redisCache) Get(key string) any {
	val, err := c.rdb.Get(c.ctx, key).Result()
	if err != nil {
		return nil
	}
	return val
}

func (c *redisCache) TTL(key string) time.Duration {
	val, _ := c.rdb.TTL(c.ctx, key).Result()
	return val
}

func (c *redisCache) Del(key string) bool {
	return c.rdb.Del(c.ctx, key).Err() != nil
}
