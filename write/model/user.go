package model

import (
	db "git.kicoe.com/blog/write/database"
	"git.kicoe.com/blog/write/errors"
)

type User struct {
	ID	int64	`xorm:"pk autoincr 'id'" json:"id"`
	Name	string	`json:"name"`
	// Avatar	string	`json:"avatar"`
	Password	string	`json:"password"`
}

func InsertUser(user *User) (affected int64, err error) {
	affected, err = db.MysqlXEngine.Insert(user)
	return
}

// func UpdateUserAvatar(id int64, avatarUrl string) (affected int64, err error) {
// 	user := &User{
// 		ID:	id,
// 		Avatar:	avatarUrl,
// 	}
// 	affected, err = db.MysqlXEngine.Update(user)
// 	return
// }

func GetUserByName(name string) (user *User, err error) {
	user = new(User)
	user.Name = name
	has, err := db.MysqlXEngine.Get(user)
	if err != nil {
		return
	}
	if !has {
		err = errors.ServiceResourceNotFoundError
	}
	return
}

func ExistUser(name string, password string) (has bool, err error) {
	user := &User{
		Name: name,
		Password: password,
	}
	has, err = db.MysqlXEngine.Exist(user)
	return
}

func CountUser() (num int64, err error) {
	user := new(User)
	num, err = db.MysqlXEngine.Count(user)
	return
}
