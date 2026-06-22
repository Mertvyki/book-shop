package core_http_middleware

import (
	"net/http"
)

const (
	corsAllowOrigin  = "Access-Control-Allow-Origin"
	corsAllowMethods = "Access-Control-Allow-Methods"
	corsAllowHeaders = "Access-Control-Allow-Headers"
	corsMaxAge       = "Access-Control-Max-Age"
)

func CORS() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(corsAllowOrigin, "*")
			w.Header().Set(corsAllowMethods, "GET, POST, PATCH, DELETE, OPTIONS")
			w.Header().Set(corsAllowHeaders, "Content-Type, Authorization, X-Request-ID")
			w.Header().Set(corsMaxAge, "86400")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
