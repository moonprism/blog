package models

import (
	"git.kicoe.com/blog/write/modules/db"
	"time"
)

type CodeMeta struct {
	ID          int64     `xorm:"pk autoincr 'id'" json:"id"`
	Description string    `json:"description"`
	Lang        string    `json:"lang"`
	Tags        string    `json:"tags"`
	CreatedTime time.Time `xorm:"created 'created_time'" json:"created_time"`
	UpdatedTime time.Time `xorm:"updated 'updated_time'" json:"updated_time"`
}

type Code struct {
	*CodeMeta `xorm:"extends -"`
	DeletedAt time.Time `xorm:"deleted" json:"-"`
	Content   string    `json:"content"`
}

type CodeDetail struct {
	*Code
}

func FetchCodesByIds(ids []int64) (codes []*Code, err error) {
	codes = make([]*Code, 0)
	err = db.MysqlXEngine.Table("code").In("id", ids).Find(&codes)
	return
}

func FetchCodeMetas(offset, limit int, whereField interface{}) (codeMetas []*CodeMeta, err error) {
	codeMetas = make([]*CodeMeta, 0)
	engine := db.MysqlXEngine.Table("code").Limit(limit, offset).Desc("updated_time")
	if whereField != nil {
		err = engine.Where(whereField).Find(&codeMetas)
	} else {
		err = engine.Find(&codeMetas)
	}
	return
}

func CountCode(whereField interface{}) (count int64, err error) {
	if whereField != nil {
		count, err = db.MysqlXEngine.Where(whereField).Count(new(Code))
	} else {
		count, err = db.MysqlXEngine.Count(new(Code))
	}
	return
}

func FetchCode(id int64) (code *Code, has bool, err error) {
	code = new(Code)
	has, err = db.MysqlXEngine.Id(id).Get(code)
	return
}

func InsertCode(code *Code) (affected int64, err error) {
	affected, err = db.MysqlXEngine.Insert(code)
	return
}

func UpdateCode(id int64, code *Code, mustCols []string) (affected int64, err error) {
	affected, err = db.MysqlXEngine.Id(id).MustCols(mustCols...).Update(code)
	return
}

func DeleteCode(id int64) (affected int64, err error) {
	var code Code
	affected, err = db.MysqlXEngine.Id(id).Delete(&code)
	return
}
