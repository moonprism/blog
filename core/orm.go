package core

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type orm struct {
	OrmClient   *gorm.DB
	SqliteFtsDB *sql.DB
	driverName  string
}

func newORM(driver string, source string) (o *orm, err error) {
	o = &orm{}
	switch driver {
	case "mysql":
		o.OrmClient, err = gorm.Open(mysql.Open(source), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	case "sqlite":
		o.OrmClient, err = gorm.Open(sqlite.Open(source), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	default:
		err = errors.New("the Driver is not supported")
	}
	o.driverName = driver
	sql.Register("sqlite3_simple",
		&sqlite3.SQLiteDriver{
			Extensions: []string{
				"libsimple-aarch64-linux-gnu-gcc-9/libsimple",
			},
		},
	)
	if err != nil {
		return
	}
	db, err := sql.Open("sqlite3_simple", source)
	if err != nil {
		return
	}
	o.SqliteFtsDB = db
	return
}

func (o *orm) DateFormatField(field string, format string) string {
	switch o.driverName {
	case "mysql":
		return fmt.Sprintf("FROM_UNIXTIME(%s, '%s')", field, format)
	case "sqlite":
		return fmt.Sprintf("strftime('%s', datetime(%s, 'unixepoch'))", format, field)
	}
	return ""
}

func (o *orm) Model(value interface{}) *gorm.DB {
	return o.OrmClient.Model(value)
}

func (o *orm) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return o.OrmClient.Transaction(fc)
}

func (o *orm) Create(value interface{}) (tx *gorm.DB) {
	return o.OrmClient.Create(value)
}

func (o *orm) Delete(value interface{}, conds ...interface{}) (tx *gorm.DB) {
	return o.OrmClient.Delete(value, conds...)
}

func (o *orm) Order(value interface{}) (tx *gorm.DB) {
	return o.OrmClient.Order(value)
}

func (o *orm) Raw(sql string, values ...interface{}) (tx *gorm.DB) {
	return o.OrmClient.Raw(sql, values)
}

func (o *orm) Where(query interface{}, args ...interface{}) (tx *gorm.DB) {
	return o.OrmClient.Where(query, args)
}

func (o *orm) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return o.OrmClient.First(dest, conds)
}

func (o *orm) Exec(sql string, values ...interface{}) (tx *gorm.DB) {
	return o.OrmClient.Exec(sql, values)
}

func (o *orm) Save(value interface{}) (tx *gorm.DB) {
	return o.OrmClient.Save(value)
}

func (o *orm) IsRecordNotFoundErr(err error) bool {
	return err == gorm.ErrRecordNotFound
}
