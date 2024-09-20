package models

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
