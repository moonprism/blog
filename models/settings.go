package models

type Settings struct {
	BaseModel

	Title        string `gorm:"type:varchar(255);notnull" json:"title"`
	Background   string `gorm:"type:text;notnull;" json:"background"`
	MarginBottom int    `gorm:"notnull;default:0" json:"marginBottom"`
}
