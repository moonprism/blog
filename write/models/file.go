package models

import (
	"git.kicoe.com/blog/write/modules/db"
	"time"
)

type FileMeta struct {
	Key string `json:"key"`
}

// File DO
type File struct {
	ID          int64 `xorm:"pk autoincr 'id'" json:"id"`
	*FileMeta   `xorm:"extends -"`
	CreatedTime time.Time `xorm:"created 'created_time'" json:"created_time"`
	DeletedAt   time.Time `xorm:"deleted" json:"-"`
}

func FetchFileByCreated(limit int, maxTime string) (files []*File, err error) {
	files = make([]*File, 0)
	engine := db.MysqlXEngine.Limit(limit).Desc("created_time")
	if maxTime != "" {
		engine = engine.Where("created_time < ?", maxTime)
	}
	err = engine.Find(&files)
	return
}

func InsertFile(file *File) (affected int64, err error) {
	affected, err = db.MysqlXEngine.Insert(file)
	return
}

func DeleteFile(id int) (affected int64, err error) {
	var file File
	affected, err = db.MysqlXEngine.Id(id).Delete(&file)
	return
}

func FetchFile(id int) (file *File, has bool, err error) {
	file = new(File)
	has, err = db.MysqlXEngine.Id(id).Get(file)
	return
}
