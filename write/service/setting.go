package service

import (
	"git.kicoe.com/blog/write/utils"
	"git.kicoe.com/blog/write/model"
	"encoding/json"
)

func GetSetting() (setting *model.Setting, err error) {
	data, err := utils.NewRedisClient().Get("aqua:blog:setting").Result()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(data), &setting)
	return
}

func SetSetting(setting *model.Setting) (err error) {
	data, err := json.Marshal(setting)
	if err != nil {
		return
	}
	err = utils.NewRedisClient().Set("aqua:blog:setting", data, 0).Err()
	return
}
