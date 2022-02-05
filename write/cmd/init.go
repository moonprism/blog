package cmd

import (
	"git.kicoe.com/blog/write/modules/db"
	"git.kicoe.com/blog/write/modules/logs"
	"git.kicoe.com/blog/write/modules/redis"
	"git.kicoe.com/blog/write/modules/setting"
)

func Init(env string) {
	setting.Init(env)
	logs.Init()
	db.Init()
	redis.Init()
}
