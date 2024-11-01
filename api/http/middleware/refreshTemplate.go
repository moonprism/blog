package middleware

import (
	"net/http"

	"github.com/moonprism/blog/core"
)

func RefreshTemplate(app *core.App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			app.InitTmpl()
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
