package routes

import (
	"net/http"

	"github.com/RudyItza/mek-ah-tell-yuh/handlers"
	"github.com/RudyItza/mek-ah-tell-yuh/internal"
)

// Routes function sets up the application routes
func Routes(app *internal.Application) http.Handler {
	mux := http.NewServeMux()

	// Home page route
	mux.HandleFunc("/", handlers.Home)

	// Feedback submission route
	mux.HandleFunc("/feedback/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateFeedback(app, w, r) // Passing app to the handler
	})

	return mux
}
