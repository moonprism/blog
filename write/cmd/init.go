package cmd

import (
	"git.kicoe.com/blog/write/modules/db"
	"git.kicoe.com/blog/write/modules/logs"
	"git.kicoe.com/blog/write/modules/redis"
	"git.kicoe.com/blog/write/modules/se"
	"git.kicoe.com/blog/write/modules/setting"
)

func Init() {
	setting.Init()
	logs.Init()
	db.Init()
	redis.Init()
	se.InitRiot()
}
