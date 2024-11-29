package api

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	m "github.com/moonprism/blog/api/http/middleware"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/ui"
)

func Serve(app *core.App) error {
	r := chi.NewRouter()
	if app.IsDev() {
		r.Use(middleware.Logger)
	}
	r.Use(m.Recoverer)
	// https://go-chi.io/#/pages/middleware?id=realip
	r.Use(middleware.RealIP)
	//r.Use(m.Delay)

	if app.IsDev() {
		vanillaStaticFS := ui.GetVanillaDistFS()
		if vanillaStaticFS != nil {
			// 单个可执行文件专用
			vanillaFileServer := http.FileServer(http.FS(vanillaStaticFS))
			r.Handle("/v/*", http.StripPrefix("/v", vanillaFileServer))
			r.Get("/404/*", func(w http.ResponseWriter, r *http.Request) {
				filePath := chi.URLParam(r, "*")
				data, err := ui.ReadAdminDistFile(filePath)
				if err != nil {
					// 如果文件不存在，返回自定义 404 页面
					data, err = ui.ReadAdminDistFile("404.html")
					core.P(err)
				}
				ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(filePath), "."))
				switch ext {
				case "css":
					w.Header().Set("Content-Type", "text/css")
				case "js":
					w.Header().Set("Content-Type", "application/javascript")
				default:
					w.Header().Set("Content-Type", "text/html")
				}
				w.Write(data)
			})
		} else {
			// 直接使用当前目录静态文件
			staticDir := "./www"
			r.Handle("/v/*", http.StripPrefix("/v", http.FileServer(http.Dir(filepath.Join(staticDir, "/v")))))
			r.Get("/404/*", func(w http.ResponseWriter, r *http.Request) {
				filePath := filepath.Join(staticDir, r.URL.Path)
				if _, err := os.Stat(filePath); os.IsNotExist(err) {
					data, err := os.ReadFile(filepath.Join(staticDir, "404/404.html"))
					core.P(err)
					w.Header().Set("Content-Type", "text/html")
					w.Write(data)
					return
				}
				http.ServeFile(w, r, filePath)
			})
		}
	}

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
