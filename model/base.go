package model

type BaseModel struct {
	ID      int64 `xorm:"int pk autoincr 'id'" json:"id"`
	Created int   `xorm:"notnull default(0) created" json:"created"`
	Updated int   `xorm:"notnull default(0) updated" json:"updated"`
}
