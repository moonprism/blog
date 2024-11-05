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

func newORM(driver string, source string, ftsSource string) (o *orm, err error) {
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
	db, err := sql.Open("sqlite3_simple", ftsSource)
	if err != nil {
		return
	}
	o.SqliteFtsDB = db
	return
}

// TODO ORM, 连接池
func (o *orm) FtsExec(query string, args ...any) (sql.Result, error) {
	return o.SqliteFtsDB.Exec(query, args...)
}
func (o *orm) FtsQuery(query string, args ...any) (*sql.Rows, error) {
	return o.SqliteFtsDB.Query(query, args...)
}

func (o *orm) FtsCreateTable(name string) error {
	_, err := o.FtsExec(fmt.Sprintf(`
DROP TABLE IF EXISTS %s_fts;
CREATE VIRTUAL TABLE %s_fts USING fts5(
	fulltext,
	tokenize = 'simple'
)`, name, name))
	return err
}

func (o *orm) FtsInsert(table string, id uint, text *string) error {
	_, err := o.FtsExec(
		fmt.Sprintf("INSERT INTO %s_fts (rowid, fulltext) VALUES (?, ?)", table),
		id,
		*text,
	)
	return err
}

func (o *orm) FtsUpdate(table string, id uint, text *string) error {
	_, err := o.FtsExec(
		fmt.Sprintf("UPDATE %s_fts SET fulltext = ? WHERE rowid = ?", table),
		*text,
		id,
	)
	return err
}

func (o *orm) FtsDelete(table string, id uint) error {
	_, err := o.FtsExec(
		fmt.Sprintf("DELETE FROM %s_fts WHERE rowid = ?", table),
		id,
	)
	return err
}

const FTS_SEARCH_START_IDX = "☾🔮☽"
const FTS_SEARCH_END_IDX = "☾†🔮☽"

var FtsSearchIdxLen = len(FTS_SEARCH_START_IDX) + len(FTS_SEARCH_END_IDX)

func (o *orm) FtsSelect(table string, keyword string, limit int, pos bool) (*sql.Rows, error) {
	var posField string
	if pos {
		posField = fmt.Sprintf("simple_highlight_pos(%s_fts, 0),", table)
	}
	return o.FtsQuery(fmt.Sprintf(`
		SELECT
			rowid,
			%s
			simple_highlight(%s_fts, 0, '%s', '%s')
		FROM %s_fts WHERE
			fulltext match jieba_query(?)
		ORDER BY rank LIMIT ?
			`,
		posField,
		table,
		FTS_SEARCH_START_IDX,
		FTS_SEARCH_END_IDX,
		table,
	),
		keyword,
		limit,
	)
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
	return o.OrmClient.Raw(sql, values...)
}

func (o *orm) Where(query interface{}, args ...interface{}) (tx *gorm.DB) {
	return o.OrmClient.Where(query, args...)
}

func (o *orm) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return o.OrmClient.First(dest, conds...)
}

func (o *orm) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return o.OrmClient.Find(dest, conds...)
}

func (o *orm) Exec(sql string, values ...interface{}) (tx *gorm.DB) {
	return o.OrmClient.Exec(sql, values...)
}

func (o *orm) Save(value interface{}) (tx *gorm.DB) {
	return o.OrmClient.Save(value)
}

func (o *orm) IsRecordNotFoundErr(err error) bool {
	return err == gorm.ErrRecordNotFound
}
