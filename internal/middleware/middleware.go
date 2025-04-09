package middleware

import (
	"net/http"

	"github.com/RudyItza/mek-ah-tell-yuh/internal" // Import internal package
)

// RateLimit is an example of a rate-limiting middleware
func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement rate limiting logic here (if necessary)
		next.ServeHTTP(w, r)
	})
}

// IsAuthenticated is a middleware that checks if a user is logged in
func IsAuthenticated(app *internal.Application, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use the GetUserID method to check if the user is logged in
		userID, err := app.Session.GetUserID(r)
		if err != nil || userID == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
