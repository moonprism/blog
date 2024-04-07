package model

import "time"

type BaseModel struct {
	ID      int64     `xorm:"pk autoincr 'id'" json:"id"`
	Created time.Time `xorm:"created 'created'" json:"created"`
	Updated time.Time `xorm:"updated 'updated'" json:"updated"`
}
