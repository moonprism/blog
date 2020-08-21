package service

import (
	"git.kicoe.com/blog/write/errors"
	"git.kicoe.com/blog/write/model"
)

func CreateUser(user *model.User) (err error) {
	affected, err := model.InsertUser(user)

	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
	}
	return
}
