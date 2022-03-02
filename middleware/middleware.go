package middleware

import "net/http"

func JSONResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("content-type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
