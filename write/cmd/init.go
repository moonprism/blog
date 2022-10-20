package cmd

import (
	"sync"

	"git.kicoe.com/blog/write/modules/db"
	"git.kicoe.com/blog/write/modules/logs"
	"git.kicoe.com/blog/write/modules/redis"
	"git.kicoe.com/blog/write/modules/setting"
)

var initOnce sync.Once

func Init() {
	initOnce.Do(func() {
		setting.Init()
		logs.Init()
		db.Init()
		redis.Init()
	})
}
