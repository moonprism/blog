package service

import (
	"git.kicoe.com/blog/write/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(auth *AuthBody) bool {
	count, err := model.CountUser()
	if err == nil && count == 0 {
		// regist
		user := new(model.User)
		user.Name = auth.Username
		// passwd hash
		hash, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
		if err != nil {
			return false
		}
		user.Password = string(hash)
		_, err = model.InsertUser(user)
		if err != nil {
			return false
		}
	}

	user, err := model.GetUserByName(auth.Username)

	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(auth.Password))
		if err == nil {
			return true
		}
	}

	return false
}
