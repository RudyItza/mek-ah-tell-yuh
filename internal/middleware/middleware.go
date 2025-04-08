package middleware

import (
	"net/http"
)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement rate limiting logic here
		next.ServeHTTP(w, r)
	})
}
