package model

import (
	db "git.kicoe.com/blog/write/database"
	"time"
)

// Comment DO
type Comment struct {
	ID	int64	`xorm:"pk autoincr 'id'" json:"id"`
	ArtID	int64	`xorm:"'art_id'" json:"art_id"`
	ToID	int64	`xorm:"'to_id'" json:"to_id"`
	Name	string	`json:"name"`
	Email	string	`json:"email"`
	Text	string	`json:"text"`
	CreatedTime	time.Time	`xorm:"created 'created_time'" json:"created_time"`
}

func FetchComments(offset, limit int, whereField interface{}) (comments []*Comment, err error) {
	comments = make([]*Comment, 0)
	engine := db.MysqlXEngine.Table("comment").Limit(limit, offset).Desc("created_time")
	if whereField != nil {
		err = engine.Where(whereField).Find(&comments)
	} else {
		err = engine.Find(&comments)
	}
	return
}

func CountComment(whereField interface{}) (count int64, err error) {
	if whereField != nil {
		count, err = db.MysqlXEngine.Where(whereField).Count(new(Comment))
	} else {
		count, err = db.MysqlXEngine.Count(new(Comment))
	}
	return
}

func DeleteComment(id int64) (affected int64, err error) {
	var comment Comment
	affected, err = db.MysqlXEngine.Id(id).Delete(&comment)
	return
}
