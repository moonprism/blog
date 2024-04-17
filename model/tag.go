package model

type Tag struct {
	BaseModel

	Name string `gorm:"type:varchar(20);notnull;index" json:"name"`
}
