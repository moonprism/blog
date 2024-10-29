package models

import "strings"

type Gist struct {
	BaseModel

	Title      string      `gorm:"type:varchar(255);notnull" json:"title"`
	Lang       string      `gorm:"type:varchar(20);notnull" json:"lang"`
	Content    string      `gorm:"type:text;notnull" json:"content"`
	GistOutput *GistOutput `json:"output"`
}

type GistOutput struct {
	GistID uint   `gorm:"primaryKey" json:"-"`
	HTML   string `gorm:"type:text;notnull;" json:"html"`
}

var GistFtsInitSQL = `
DROP TABLE IF EXISTS gists_fts;
CREATE VIRTUAL TABLE gists_fts USING fts5(
	fulltext,
	tokenize = 'simple'
)`

type GistFts struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Lang    string `json:"lang"`
	Content string `json:"content"`
}

func Gist2Text(g *Gist) string {
	var builder strings.Builder
	builder.WriteString(g.Title)
	builder.WriteString(".")
	builder.WriteString(g.Lang)
	builder.WriteString("\n")
	builder.WriteString(g.Content)
	return builder.String()
}

func Text2Gist(id uint, s *string) *GistFts {
	index := strings.IndexByte(*s, '\n')
	firstLine := (*s)[:index]
	content := (*s)[index+1:]
	dotIndex := strings.LastIndex(firstLine, ".")
	title := firstLine[:dotIndex]
	lang := firstLine[dotIndex+1:] // '.' 字符后的部分
	return &GistFts{
		id,
		title,
		lang,
		content,
	}
}
