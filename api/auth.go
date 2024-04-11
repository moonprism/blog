package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/core/http/middleware"
	"golang.org/x/crypto/bcrypt"
)

func bindAuthApi(app *core.App, r chi.Router) {
	api := authApi{app}
	r.Use(middleware.JsonResponse)
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
	core.Pf(err)
	if req.Username != api.App.Setting.Account.Name {
		core.Pf(errors.New("login failed"), core.CodeLoginFailed)
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(api.App.Setting.Account.Pass),
		[]byte(req.Password),
	)
	if err != nil {
		core.Pf(errors.New("login failed"), core.CodeLoginFailed)
	}
	_, tokenString, err := api.App.TokenAuth.Encode(map[string]interface{}{"username": req.Username})
	core.Pf(err)
	res := new(loginResponseBody)
	res.Token = tokenString
	json.NewEncoder(w).Encode(&res)
}
