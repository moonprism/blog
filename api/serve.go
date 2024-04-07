package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/moonprism/blog/core"
)

func Serve(app *core.App) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/articles", func(r chi.Router) { bindArticleApi(app, r) })
	http.ListenAndServe(":3000", r)
}
