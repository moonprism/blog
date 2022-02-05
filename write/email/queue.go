package email

import (
	"time"

	"github.com/go-redis/redis/v7"
)

type RedisQueue struct {
	c      redis.UniversalClient
	prefix string
}

func (q *RedisQueue) Pop(key string) (string, error) {
	res, err := q.c.BLPop(100*time.Second, q.prefix+key).Result()
	if err != nil {
		return "", err
	}
	return res[1], err
}

func NewRedisQueue(c redis.UniversalClient, prefix string) *RedisQueue {
	return &RedisQueue{
		c,
		prefix,
	}
}
