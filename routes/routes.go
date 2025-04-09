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

	// Login page route
	mux.HandleFunc("/login", handlers.Login)

	// Create story route
	mux.HandleFunc("/stories/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateStory(app, w, r) // Passing app to the handler
	})

	// Edit story route
	mux.HandleFunc("/stories/edit", func(w http.ResponseWriter, r *http.Request) {
		handlers.EditStory(app, w, r) // Passing app to the handler
	})

	// Delete story route
	mux.HandleFunc("/stories/delete", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteStory(app, w, r) // Passing app to the handler
	})

	return mux
}
