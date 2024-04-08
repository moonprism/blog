package api

import (
	"encoding/json"
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
	Name     string `json:"name"`
	Password string `json:"password"`
}

type loginResponseBody struct {
	Token string `json:"token"`
}

func (api *authApi) login(w http.ResponseWriter, r *http.Request) {
	var req loginRequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	core.Pf(err)
	if req.Name != api.App.Setting.Account.Name {
		w.WriteHeader(401)
		return
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(api.App.Setting.Account.Pass),
		[]byte(req.Password),
	)
	if err != nil {
		w.WriteHeader(401)
		return
	}
	_, tokenString, err := api.App.TokenAuth.Encode(map[string]interface{}{"user_id": req.Name})
	core.Pf(err)
	res := new(loginResponseBody)
	res.Token = tokenString
	json.NewEncoder(w).Encode(&res)
}
