package middleware

import (
	"math/rand"
	"net/http"
	"time"
)

// for test
func Delay(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(3500)
		if t > 3000 {
			panic("server error")
		}
		time.Sleep(time.Duration(t * int(time.Millisecond)))
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
