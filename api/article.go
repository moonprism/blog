package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/moonprism/blog/core"
)

func bindArticleApi(app *core.App, r chi.Router) {
	api := articleApi{app}

	r.Get("/", api.list)
}

type articleApi struct {
	*core.App
}

func (api *articleApi) list(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
