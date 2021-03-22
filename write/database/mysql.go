package database

import (
	"fmt"
	"git.kicoe.com/blog/write/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"sync"
)

var (
	MysqlXEngine	*xorm.Engine
	once	sync.Once
)

func InitMysqlEngine() {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local",
			config.Mysql.User,
			config.Mysql.Password,
			config.Mysql.Host,
			config.Mysql.Port,
			config.Mysql.DataBase,
		)
		engine, err := xorm.NewEngine("mysql", dsn)
		if err != nil {
			panic(err)
		}
		// engine.ShowSQL(true)
		MysqlXEngine = engine
	})
}
