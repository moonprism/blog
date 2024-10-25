package models

type User struct {
	BaseModel

	Name      string `gorm:"type:varchar(100);notnull;index" json:"name"`
	Pass      string `gorm:"type:char(60);notnull" json:"pass"`
	LastLogin uint   `gorm:"notnull" json:"last_login"`
}
