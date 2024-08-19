package models

type Gist struct {
	BaseModel

	Title   string `gorm:"type:varchar(255);notnull" json:"title"`
	Lang    string `gorm:"type:varchar(20);notnull" json:"lang"`
	Content string `gorm:"type:text;notnull" json:"content"`
}
