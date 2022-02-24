package models

import "time"

type Link struct {
	ID          int64     `xorm:"pk autoincr 'id'" json:"id"`
	Link        string    `json:"link"`
	Email       string    `json:"email"`
	Status      int       `json:"status"`
	Avatar      string    `json:"avatar"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Message     string    `json:"message"`
	Data        string    `json:"data"`
	CreatedTime time.Time `xorm:"created 'created_time'" json:"created_time"`
}
