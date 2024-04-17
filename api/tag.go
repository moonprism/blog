package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/model"
)

func bindTagApi(app *core.App, r chi.Router) {
	api := tagApi{app}

	r.Get("/", api.list)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(app.TokenAuth))
		r.Use(jwtauth.Authenticator(app.TokenAuth))

		r.Post("/", api.create)
		r.Put("/{id}", api.update)
		r.Delete("/{id}", api.delete)
	})
}

type tagApi struct {
	*core.App
}

func (api *tagApi) list(w http.ResponseWriter, r *http.Request) {
	tags := make([]*model.Tag, 0)
	err := api.O.Find(&tags).Error
	core.P(err)
	json.NewEncoder(w).Encode(tags)
}

func (api *tagApi) create(w http.ResponseWriter, r *http.Request) {
	tag := new(model.Tag)
	err := json.NewDecoder(r.Body).Decode(tag)
	core.P(err)
	err = api.O.Create(tag).Error
	core.P(err)
}

func (api *tagApi) update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	core.P(err)
	tag := new(model.Tag)
	err = json.NewDecoder(r.Body).Decode(tag)
	core.P(err)
	tag.ID = uint(id)
	err = api.O.Save(tag).Error
	core.P(err)
	json.NewDecoder(r.Body).Decode(tag)
}

func (api *tagApi) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	core.P(err)
	tag := new(model.Tag)
	tag.ID = uint(id)
	api.O.Select("Articles").Delete(&tag)
	core.P(err)
}
