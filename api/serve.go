package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/moonprism/blog/core"
	m "github.com/moonprism/blog/core/http/middleware"
)

func Serve(app *core.App) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(m.Recoverer)
	r.Use(m.JsonResponse)
	//r.Use(m.Delay)

	r.Route("/api", func(r chi.Router) {
		r.Route("/", func(r chi.Router) { bindAuthApi(app, r) })
		r.Route("/article", func(r chi.Router) { bindArticleApi(app, r) })
		r.Route("/tag", func(r chi.Router) { bindTagApi(app, r) })
		r.Route("/attachment", func(r chi.Router) { bindAttachmentApi(app, r) })
		r.Route("/gist", func(r chi.Router) { bindGistApi(app, r) })

		r.Route("/group", func(r chi.Router) { bindGroupApi(app, r) })
	})

	return http.ListenAndServe(app.Setting.Server.Addr, r)
}
