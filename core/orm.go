package core

import (
	"database/sql"
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type orm struct {
	OrmClient  *gorm.DB
	driverName string
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
