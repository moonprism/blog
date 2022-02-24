package models

import (
	"fmt"
	"git.kicoe.com/blog/write/modules/db"
	"strconv"
	"time"
)

type TagMeta struct {
	ID    int64  `xorm:"pk autoincr 'id'" json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type TagInfo struct {
	*TagMeta
	Count int64 `json:"count"`
}

// tag DO
type Tag struct {
	*TagMeta    `xorm:"extends -"`
	CreatedTime time.Time `xorm:"created 'created_time'" json:"created_time"`
	UpdatedTime time.Time `xorm:"updated 'updated_time'" json:"updated_time"`
}

// tags
type TagList []Tag

// article_tag DO
type ArticleTag struct {
	ID          int64     `xorm:"pk autoincr 'id'" json:"id"`
	ArtID       int64     `xorm:"'art_id'" json:"art_id"`
	TagID       int64     `xorm:"'tag_id'" json:"tag_id"`
	CreatedTime time.Time `xorm:"created 'created_time'" json:"created_time"`
}

var (
	sqlSelectArticleTags = "select t.id, t.name, t.color, article_tag.art_id from article_tag inner join `tag` as t on article_tag.tag_id = t.id where article_tag.art_id in (%s) order by article_tag.id asc"
)

func FetchArticleTagsMap(ids string) (res map[int64][]*TagMeta, err error) {
	rowSlice, err := db.MysqlXEngine.QueryString(fmt.Sprintf(sqlSelectArticleTags, ids))
	if err != nil {
		return
	}

	res = make(map[int64][]*TagMeta)
	for _, row := range rowSlice {
		tmp := new(TagMeta)
		tmp.ID, _ = strconv.ParseInt(row["id"], 10, 64)
		tmp.Color = row["color"]
		tmp.Name = row["name"]
		artId, _ := strconv.ParseInt(row["art_id"], 10, 64)
		res[artId] = append(res[artId], tmp)
	}
	return
}

func UpdateArticleTags(id int64, tags []*TagMeta) (err error) {
	var articleTag = new(ArticleTag)

	// delete all article tag re install
	_, err = db.MysqlXEngine.Where("art_id = ?", id).Delete(articleTag)
	if err != nil {
		return
	}

	err = InsertArticleTags(id, tags)
	if err != nil {
		return
	}
	return
}

func InsertArticleTags(id int64, tags []*TagMeta) (err error) {
	var articleTag *ArticleTag

	for _, tag := range tags {
		articleTag = new(ArticleTag)
		articleTag.ArtID = id
		articleTag.TagID = tag.ID
		_, err = db.MysqlXEngine.Insert(articleTag)
		if err != nil {
			return
		}
	}
	return
}

func DeleteAllArticleTags(id int64) (err error) {
	_, err = db.MysqlXEngine.Where("tag_id = ?", id).Delete(new(ArticleTag))
	return
}

func FetchTagMetas() (tags []*TagMeta, err error) {
	tags = make([]*TagMeta, 0)
	err = db.MysqlXEngine.Table("tag").Desc("updated_time").Find(&tags)
	return
}

func CountTagReference(id int64) (count int64, err error) {
	count, err = db.MysqlXEngine.Where("tag_id = ?", id).Count(new(ArticleTag))
	return
}

func InsertTag(tag *Tag) (affected int64, err error) {
	affected, err = db.MysqlXEngine.Insert(tag)
	return
}

func UpdateTag(id int64, tag *Tag, mustCols []string) (affected int64, err error) {
	affected, err = db.MysqlXEngine.Id(id).MustCols(mustCols...).Update(tag)
	return
}

func DeleteTag(id int64) (affected int64, err error) {
	affected, err = db.MysqlXEngine.Id(id).Delete(new(Tag))
	return
}
