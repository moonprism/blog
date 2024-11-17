package models

import "strings"

type Article struct {
	BaseModel

	Title   string `gorm:"type:varchar(255);notnull;" json:"title"`
	Status  uint   `gorm:"type:tinyint;notnull;comment:0=draft,1=published,2=hidden-不在列表显示" json:"status"`
	Rune    uint   `gorm:"type:smallint;notnull;comment:不合范式的冗余字段" json:"rune"`
	Image   string `gorm:"type:varchar(255);notnull" json:"image"`
	Summary string `gorm:"type:varchar(2000);notnull" json:"summary"`

	ArticleContent *ArticleContent `json:"content"`
	Tags           []Tag           `gorm:"many2many:article_tags" json:"tags"`
}

const (
	ArticleStatusDraft     = 0
	ArticleStatusPublished = 1
)

type ArticleContent struct {
	ArticleID uint   `gorm:"primaryKey" json:"-"`
	Text      string `gorm:"type:text;notnull;" json:"text"`
	HTML      string `gorm:"type:mediumtext;notnull;" json:"html"`
}

type ArticleTags struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	ArticleID uint `gorm:"notnull" json:"article_id"`
	TagID     uint `gorm:"notnull" json:"tag_id"`
}

func Art2TextPoint(a *Article) *string {
	var builder strings.Builder
	builder.WriteString(a.Title)
	builder.WriteString("\n")
	builder.WriteString(a.ArticleContent.Text)
	s := builder.String()
	return &s
}

func Text2ArtGistFts(id uint, s *string) *GistFts {
	index := strings.IndexByte(*s, '\n')
	return &GistFts{
		ID:      id,
		Title:   (*s)[:index],
		Lang:    "md",
		Content: (*s)[index+1:],
	}
}
