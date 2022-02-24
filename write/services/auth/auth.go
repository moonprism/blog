package auth

import (
	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/modules/db"
	"golang.org/x/crypto/bcrypt"
)

type AuthBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Upsert(auth *AuthBody) error {
	session := db.MysqlXEngine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}

	model := models.NewUser(session)
	has, err := model.ExistUserById(1)
	if err != nil {
		session.Rollback()
		return err
	}

	user := new(models.User)
	user.Name = auth.Username
	// passwd hash
	hash, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		session.Rollback()
		return err
	}
	user.Password = string(hash)
	model.User = *user
	if has {
		_, err = model.Update(1)
	} else {
		_, err = model.Insert()
	}
	if err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}

func Check(auth *AuthBody) bool {
	model := models.NewUser(nil)
	user, err := model.GetByName(auth.Username)

	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(auth.Password))
		if err == nil {
			return true
		}
	}

	return false
}
