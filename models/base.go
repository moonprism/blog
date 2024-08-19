package models

import "gorm.io/plugin/soft_delete"

type BaseModel struct {
	ID      uint `json:"id"`
	Created uint `gorm:"notnull;autoCreateTime" json:"created"`
	Updated uint `gorm:"notnull;autoUpdateTime" json:"updated"`

	IsDel soft_delete.DeletedAt `gorm:"type:tinyint(1);notnull;softDelete:flag" json:"-"`
}
