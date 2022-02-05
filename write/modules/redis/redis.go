package redis

import (
	"git.kicoe.com/blog/write/modules/setting"
	"github.com/go-redis/redis/v7"
)

var Client *redis.Client

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     setting.Redis.Host + ":" + setting.Redis.Port,
		Password: "",
		DB:       0,
	})
	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
