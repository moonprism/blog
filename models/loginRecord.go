package models

type LoginRecord struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	Created uint `gorm:"notnull;autoCreateTime" json:"created"`
	//IP      uint `json:"ip"`
	IP       string `gorm:"type:varchar(39);notnull" json:"ip"`
	IsFailed bool   `gorm:"default:false" json:"is_failed"`
}
