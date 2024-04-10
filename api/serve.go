package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/moonprism/blog/core"
)

func Serve(app *core.App) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Route("/", func(r chi.Router) { bindAuthApi(app, r) })
		r.Route("/article", func(r chi.Router) { bindArticleApi(app, r) })
	})

	return http.ListenAndServe(app.Setting.Server.Addr, r)
}
