package models

import "gorm.io/plugin/soft_delete"

type BaseModel struct {
	ID      uint `json:"id"`
	Created uint `gorm:"notnull;autoCreateTime" json:"created"`
	Updated uint `gorm:"notnull;autoUpdateTime" json:"updated"`

	IsDel soft_delete.DeletedAt `gorm:"type:tinyint(1);notnull;softDelete:flag" json:"-"`
}

type Pagination struct {
	Count int `json:"count"`
}

type SearchURLParams[T int | string] struct {
	Page         int            `json:"page"`
	PageSize     int            `json:"page_size"`
	FilterText   string         `json:"filter_text"`
	FilterValues map[string][]T `json:"filter_values"`
	SortKey      struct {
		ID    string `json:"id"`
		Order string `json:"order"`
	} `json:"sort_key"`
}
