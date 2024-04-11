package middleware

import (
	"encoding/json"
	"net/http"

	"runtime/debug"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/moonprism/blog/core"
)

// Recoverer 自定义go-chi框架的middleware.Recoverer函数
// 接收到自定义panic将响应客户端{"code": 1xx, "message": ""}
// Alternatively, look at https://github.com/go-chi/httplog middleware pkgs.
func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				if rvr == http.ErrAbortHandler {
					panic(rvr)
				}

				logEntry := middleware.GetLogEntry(r)
				if logEntry != nil {
					logEntry.Panic(rvr, debug.Stack())
				} else {
					middleware.PrintPrettyStack(rvr)
				}

				if r.Header.Get("Connection") != "Upgrade" {
					w.WriteHeader(http.StatusInternalServerError)
					if pe, ok := rvr.(core.PfError); ok {
						json.NewEncoder(w).Encode(&pe)
					}
				}
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
