package model

import (
	db "git.kicoe.com/blog/write/database"
	"time"
)

type ArticleMeta struct {
	ID	int64	`xorm:"pk autoincr 'id'" json:"id"`
	Title	string	`json:"title"`
	Status	int	`json:"status"`
	Image	string	`json:"image"`
	Summary	string	`json:"summary"`
	CreatedTime	time.Time	`xorm:"created 'created_time'" json:"created_time"`
	UpdatedTime	time.Time	`xorm:"updated 'updated_time'" json:"updated_time"`
}

type ArticleInfo struct {
	*ArticleMeta
	Tags	[]*TagMeta `json:"tags"`
}

// article DO
type Article struct {
	*ArticleMeta	`xorm:"extends -"`
	DeletedAt	time.Time	`xorm:"deleted" json:"-"`
	Content	string	`json:"content"`
	HTML	string	`xorm:"'html'" json:"html"`
}

type ArticleDetail struct {
	*Article
	Tags	[]*TagMeta `json:"tags"`
}

func FetchArticleMetas(offset, limit int, whereField interface{}) (articleMetas []*ArticleMeta, err error) {
	articleMetas = make([]*ArticleMeta, 0)
	engine := db.MysqlXEngine.Table("article").Limit(limit, offset).Desc("created_time")
	if whereField != nil {
		err = engine.Where(whereField).Find(&articleMetas)
	} else {
		err = engine.Find(&articleMetas)
	}
	return
}

func CountArticle(whereField interface{}) (count int64, err error) {
	if whereField != nil {
		count, err = db.MysqlXEngine.Where(whereField).Count(new(Article))
	} else {
		count, err = db.MysqlXEngine.Count(new(Article))
	}
	return
}

func FetchArticle(id int64) (article *Article, has  bool, err error) {
	article = new(Article)
	has, err = db.MysqlXEngine.Id(id).Get(article)
	return
}

func InsertArticle(article *Article) (affected int64, err error) {
	affected, err = db.MysqlXEngine.Insert(article)
	return
}

func UpdateArticle(id int64, article *Article, mustCols []string) (affected int64, err error) {
	affected, err = db.MysqlXEngine.Id(id).MustCols(mustCols...).Update(article)
	return
}

func DeleteArticle(id int64) (affected int64, err error) {
	var article Article
	affected, err = db.MysqlXEngine.Id(id).Delete(&article)
	return
}