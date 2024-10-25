package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/models"
	"golang.org/x/crypto/bcrypt"
)

func bindAuthApi(app *core.App, r chi.Router) {
	api := authApi{app}
	r.Post("/login", api.login)
}

type authApi struct {
	*core.App
}

type loginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponseBody struct {
	Token string `json:"token"`
}

func (api *authApi) login(w http.ResponseWriter, r *http.Request) {
	var req loginRequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	defer func() {
		record := new(models.LoginRecord)
		record.IP = r.RemoteAddr
		record.IsFailed = err != nil
		core.P(api.O.Create(record).Error)
	}()
	core.P(err)
	var user models.User
	err = api.O.Where(&models.User{Name: req.Username}).First(&user).Error
	if api.O.IsRecordNotFoundErr(err) {
		core.PanicErr("User does not exist", core.ErrCodeLoginFailed)
	}
	core.P(err)
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Pass),
		[]byte(req.Password),
	)
	if err != nil {
		core.PanicErr("Incorrect password", core.ErrCodeLoginFailed)
	}
	_, tokenString, err := api.TokenAuth.Encode(map[string]interface{}{
		"user": req.Username,
		// jwt 过期时间
		"exp": time.Now().Add(api.Settings.System.TokenExpiryHours * time.Hour).Unix(),
	})
	core.P(err)
	err = api.O.Model(&user).Updates(models.User{LastLogin: uint(time.Now().Unix())}).Error
	core.P(err)
	res := new(loginResponseBody)
	res.Token = tokenString
	json.NewEncoder(w).Encode(&res)
}
