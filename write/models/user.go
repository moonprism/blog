package models

import (
	"git.kicoe.com/blog/write/modules/db"
	"git.kicoe.com/blog/write/modules/err/errors"
	"github.com/go-xorm/xorm"
)

type User struct {
	ID   int64  `xorm:"pk autoincr 'id'" json:"id"`
	Name string `json:"name"`
	// Avatar	string	`json:"avatar"`
	Password string `json:"password"`
}

type UserEngine struct {
	db xorm.Interface
	User
}

func NewUser(xi xorm.Interface) *UserEngine {
	if xi == nil {
		xi = db.MysqlXEngine
	}
	return &UserEngine{
		db: xi,
	}
}

func (u *UserEngine) Insert() (affected int64, err error) {
	affected, err = u.db.Insert(&u.User)
	return
}

func (u *UserEngine) Update(id int64) (affected int64, err error) {
	affected, err = u.db.ID(id).Update(&u.User)
	return
}

func (u *UserEngine) GetByName(name string) (user *User, err error) {
	u.Name = name
	has, err := u.db.Get(&u.User)
	if err != nil {
		return
	}
	if !has {
		err = errors.ServiceResourceNotFoundError
	}
	user = &u.User
	return
}

func (u *UserEngine) ExistUser() (has bool, err error) {
	count, err := u.CountUser()
	has = count > 0
	return
}

func (u *UserEngine) ExistUserById(id int64) (has bool, err error) {
	user := &User{
		ID: id,
	}
	has, err = u.db.Exist(user)
	return
}

func (u *UserEngine) CountUser() (count int64, err error) {
	return u.db.Count(&u.User)
}
