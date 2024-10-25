package models

type Settings struct {
	BaseModel

	Title      string `gorm:"type:varchar(255);notnull" json:"title"`
	Background string `gorm:"type:varchar(255);notnull" json:"background"`
}
