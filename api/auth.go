package api

import (
	"encoding/json"
	"net/http"

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
	if req.Username != api.Setting.Account.Name {
		err = core.NewErr("login failed", core.ErrCodeLoginFailed)
		core.P(err)
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(api.Setting.Account.Pass),
		[]byte(req.Password),
	)
	if err != nil {
		core.P(core.NewErr("login failed", core.ErrCodeLoginFailed))
	}
	_, tokenString, err := api.TokenAuth.Encode(map[string]interface{}{"username": req.Username})
	core.P(err)
	res := new(loginResponseBody)
	res.Token = tokenString
	json.NewEncoder(w).Encode(&res)
}
