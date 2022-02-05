package system

import (
	"encoding/json"
	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/modules/redis"
)


func GetTheme() (setting models.Theme, err error) {
	data, err := redis.Client.Get("aqua:blog:setting").Result()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(data), &setting)
	return
}

func SetTheme(setting *models.Theme) (err error) {
	data, err := json.Marshal(setting)
	if err != nil {
		return
	}
	err = redis.Client.Set("aqua:blog:setting", data, 0).Err()
	return
}
