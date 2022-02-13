package models

import (
	"git.kicoe.com/blog/write/modules/db"
	"time"
)

// Comment DO
type Comment struct {
	ID          int64     `xorm:"pk autoincr 'id'" json:"id"`
	ArtID       int64     `xorm:"'art_id'" json:"art_id"`
	ToID        int64     `xorm:"'to_id'" json:"to_id"`
	TopID       int64     `xorm:"'top_id'" json:"top_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Text        string    `json:"text"`
	CreatedTime time.Time `xorm:"created 'created_time'" json:"created_time"`
	DeletedAt    time.Time `xorm:"deleted" json:"-"`
}

func FetchComments(offset, limit int, whereField interface{}) (comments []*Comment, err error) {
	comments = make([]*Comment, 0)
	engine := db.MysqlXEngine.Limit(limit, offset).Desc("created_time")
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

func FetchComment(id int64) (comment *Comment, has bool, err error) {
	comment = new(Comment)
	has, err = db.MysqlXEngine.Id(id).Get(comment)
	return
}

func DeleteComment(id int64) (affected int64, err error) {
	session := db.MysqlXEngine.NewSession()
	defer session.Close()

	comments := make([]*Comment, 0)
	comment := new(Comment)
	session.Id(id).Get(comment)
	if comment.TopID == 0 {
		session.Where("top_id = ?", comment.ID).Find(&comments)
	} else {
		session.Where("to_id = ?", comment.ID).Find(&comments)
	}

	ids := []int64{id}
	for _, c := range(comments) {
		ids = append(ids, c.ID)
	}

	affected, err = session.In("id", ids).Delete(&Comment{})
	if err != nil {
		session.Rollback()
	}
	session.Commit()
	return
}
