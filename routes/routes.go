// routes.go

package routes

import (
	"net/http"

	"github.com/RudyItza/mek-ah-tell-yuh/handlers"
	"github.com/RudyItza/mek-ah-tell-yuh/internal"
	"github.com/RudyItza/mek-ah-tell-yuh/internal/middleware"
	"github.com/gorilla/mux"
)

// Routes function sets up the application routes
func Routes(app *internal.Application) http.Handler {
	muxRouter := mux.NewRouter()

	// Home page route
	muxRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}).Methods(http.MethodGet)

	muxRouter.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		handlers.Home(app, w, r)
	}).Methods(http.MethodGet)

	// Feedback submission route
	muxRouter.HandleFunc("/feedback/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateFeedback(app, w, r)
	}).Methods(http.MethodGet, http.MethodPost)

	// Feedback success route
	muxRouter.HandleFunc("/feedback/success", func(w http.ResponseWriter, r *http.Request) {
		// Render feedback success page
		err := app.Templates.Render(w, "feedback_success.tmpl", nil)
		if err != nil {
			http.Error(w, "Could not render success page", http.StatusInternalServerError)
		}
	}).Methods(http.MethodGet)

	// Login page route
	muxRouter.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.Login(app, w, r)
	}).Methods(http.MethodGet, http.MethodPost)

	// Apply authentication middleware to the story creation route
	muxRouter.Handle("/stories/new", middleware.IsAuthenticated(app, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateStory(app, w, r)
	}))).Methods(http.MethodGet, http.MethodPost)

	// Apply authentication middleware to other routes (e.g., edit, delete)
	muxRouter.Handle("/stories/{id}/edit", middleware.IsAuthenticated(app, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.EditStory(app, w, r)
	}))).Methods(http.MethodGet, http.MethodPost)

	muxRouter.Handle("/stories/{id}/delete", middleware.IsAuthenticated(app, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteStory(app, w, r)
	}))).Methods(http.MethodPost)

	// Other routes...

	return muxRouter
}
