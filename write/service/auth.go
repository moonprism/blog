package service

import (
	"git.kicoe.com/blog/write/config"
)

type AuthBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(auth *AuthBody) bool {
	if auth.Username == config.Auth.User && auth.Password == config.Auth.Password {
		return true
	}

	return false
}
