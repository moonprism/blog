package db

import (
	"fmt"
	"git.kicoe.com/blog/write/modules/setting"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	MysqlXEngine *xorm.Engine
)

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local",
		setting.Mysql.User,
		setting.Mysql.Password,
		setting.Mysql.Host,
		setting.Mysql.Port,
		setting.Mysql.DataBase,
	)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// engine.ShowSQL(true)
	MysqlXEngine = engine
}
