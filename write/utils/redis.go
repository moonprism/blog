package utils

import (
	"git.kicoe.com/blog/write/config"
	"github.com/go-redis/redis/v7"
	"sync"
)

var (
	redisClient *redis.Client
	redisOnce sync.Once
)

func NewRedisClient() *redis.Client {
	redisOnce.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     config.Redis.Host + ":" + config.Redis.Port,
			Password: "",
			DB:       0,
		})
		_, err := client.Ping().Result()
		if err != nil {
			panic(err)
		}
		redisClient = client
	})
	return redisClient
}

func InitRedis() {
	NewRedisClient()
}