package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/models"
)

func bindGroupApi(app *core.App, r chi.Router) {
	api := groupApi{app}

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(app.TokenAuth))
		r.Use(jwtauth.Authenticator(app.TokenAuth))
		r.Get("/year", api.byYear)
		r.Get("/gist-lang", api.byGistLang)
	})
}

type groupApi struct {
	*core.App
}

type GroupByYear struct {
	Year  int `json:"year"`
	Count int `json:"count"`
}

func (api *groupApi) byYear(w http.ResponseWriter, r *http.Request) {
	// q := r.URL.Query().Get("model")
	var result []*GroupByYear
	api.O.Model(&models.Attachment{}).
		Select(fmt.Sprintf("%s AS year, COUNT(*) AS count", api.O.DateFormatField("created", "%Y"))).
		Group("year").
		Order("id DESC").
		Find(&result)
	api.JSON(w, result)
}

type GistGroupByLang struct {
	Lang  string `json:"lang"`
	Count int    `json:"count"`
}

func (api *groupApi) byGistLang(w http.ResponseWriter, r *http.Request) {
	var result []*GistGroupByLang
	api.O.Model(&models.Gist{}).
		Select("lang, COUNT(*) AS count").
		Group("lang").
		Order("count desc").
		Find(&result)
	api.JSON(w, result)
}
