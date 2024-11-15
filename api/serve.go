package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	m "github.com/moonprism/blog/api/http/middleware"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/ui"
)

func Serve(app *core.App) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(m.Recoverer)
	// https://go-chi.io/#/pages/middleware?id=realip
	r.Use(middleware.RealIP)
	//r.Use(m.Delay)

	vanillaFileServer := http.FileServer(http.FS(ui.GetVanillaEmbedFS()))
	r.Handle("/v/*", http.StripPrefix("/v", vanillaFileServer))

	r.Route("/api", func(r chi.Router) {
		r.Use(m.JsonResponse)
		r.Route("/", func(r chi.Router) { bindAuthApi(app, r) })
		r.Route("/article", func(r chi.Router) { bindArticleApi(app, r) })
		r.Route("/tag", func(r chi.Router) { bindTagApi(app, r) })
		r.Route("/attachment", func(r chi.Router) { bindAttachmentApi(app, r) })
		r.Route("/gist", func(r chi.Router) { bindGistApi(app, r) })
		r.Route("/comment", func(r chi.Router) { bindCommentApi(app, r) })

		r.Route("/group", func(r chi.Router) { bindGroupApi(app, r) })
		r.Route("/settings", func(r chi.Router) { bindSettingsApi(app, r) })
	})

	r.Route("/", func(r chi.Router) {
		r.Get("/", articlePageListRoute(app))
		r.Get("/posts", articlePageListRoute(app))
		r.Get("/posts/tag/{id}", articlePageListRoute(app))
		r.Get("/post/{id}", articlePageDetailRoute(app))
		r.Get("/comments/{articleID}", commentPageListRoute(app))
		r.Get("/gists", gistsPageRoute(app))
		r.Get("/gists/search", gistsSearchRoute(app))
		r.Get("/links", articlePageLinksRoute(app))
		r.Get("/about", articlePageAboutRoute(app))
	})

	return http.ListenAndServe(app.Settings.Server.Addr, r)
}
