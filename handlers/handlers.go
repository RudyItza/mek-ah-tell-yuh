package handlers

import (
	"fmt"
	"net/http"

	"github.com/RudyItza/mek-ah-tell-yuh/internal"
	"github.com/RudyItza/mek-ah-tell-yuh/internal/data"
	"github.com/RudyItza/mek-ah-tell-yuh/internal/validator"
)

// Home handler renders the home page
func Home(w http.ResponseWriter, r *http.Request) {
	// Just rendering a basic message for now
	w.Write([]byte("Welcome to Mek Ah Tell Yuh!"))
}

// CreateFeedback handler for handling feedback submissions
func CreateFeedback(app *internal.Application, w http.ResponseWriter, r *http.Request) {
	// Handle POST request for creating feedback
	if r.Method == http.MethodPost {
		feedback := &data.Feedback{
			FullName: r.FormValue("fullname"),
			Subject:  r.FormValue("subject"),
			Message:  r.FormValue("message"),
		}

		// Initialize the validator
		v := validator.New()

		// Perform validation checks
		data.ValidateFeedback(v, feedback)

		if !v.Valid() {
			// If validation fails, render error message
			http.Error(w, "Validation failed", http.StatusBadRequest)
			return
		}

		// Insert feedback into the database
		if err := app.Feedback.Insert(feedback); err != nil {
			// If database insertion fails, render error message
			http.Error(w, fmt.Sprintf("Unable to create feedback: %s", err), http.StatusInternalServerError)
			return
		}

		// Redirect to a success page or render a success template
		http.Redirect(w, r, "/feedback/success", http.StatusSeeOther)
		return
	}

	// Render feedback form (you could use a template here)
	http.ServeFile(w, r, "ui/templates/feedback_form.html")
}
