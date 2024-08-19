package model

type Tag struct {
	BaseModel

	Name  string `gorm:"type:varchar(20);notnull;index" json:"name"`
	Color string `gorm:"type:varchar(15);notnull" json:"color"`
}
