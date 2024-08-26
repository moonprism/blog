package api

import (
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
		Select("FROM_UNIXTIME(created, '%Y') AS year, COUNT(*) AS count").
		Group("year").
		Order("id DESC").
		Find(&result)
	api.JSON(w, result)
}
