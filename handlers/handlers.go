package handlers

import (
	"fmt"
	"net/http"

	"github.com/RudyItza/mek-ah-tell-yuh/internal"
	"github.com/RudyItza/mek-ah-tell-yuh/internal/data"
	"github.com/RudyItza/mek-ah-tell-yuh/internal/validator"

	"github.com/go-chi/chi"
)

// Home renders the home page.
func Home(app *internal.Application, w http.ResponseWriter, r *http.Request) {
	err := app.Templates.Render(w, "home.tmpl", nil)
	if err != nil {
		app.Logger.Error("Could not render home page", err)
		http.Error(w, "Could not render home page", http.StatusInternalServerError)
	}
}

// Login handles user login.
func Login(app *internal.Application, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle login logic here
		// Example: Validate user credentials and set session values
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Render login page
	err := app.Templates.Render(w, "login.tmpl", nil)
	if err != nil {
		app.Logger.Error("Could not render login page", err)
		http.Error(w, "Could not render login page", http.StatusInternalServerError)
	}
}

// CreateFeedback handles the feedback form submission.
func CreateFeedback(app *internal.Application, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		feedback := &data.Feedback{
			FullName: r.FormValue("fullname"),
			Subject:  r.FormValue("subject"),
			Message:  r.FormValue("message"),
		}

		v := validator.New()
		// Assuming data.ValidateFeedback is a custom function for validation
		data.ValidateFeedback(v, feedback)

		if !v.Valid() {
			http.Error(w, "Validation failed", http.StatusBadRequest)
			return
		}

		if err := app.Feedback.Insert(feedback); err != nil {
			app.Logger.Error("Unable to create feedback", err)
			http.Error(w, fmt.Sprintf("Unable to create feedback: %s", err), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/feedback/success", http.StatusSeeOther)
		return
	}

	err := app.Templates.Render(w, "feedback_form.tmpl", nil)
	if err != nil {
		app.Logger.Error("Could not render feedback form", err)
		http.Error(w, "Could not render feedback form", http.StatusInternalServerError)
	}
}

// CreateStory handles the story creation form submission.
func CreateStory(app *internal.Application, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		userID, err := app.Session.GetUserID(r) // Correctly using GetUserID on app.Session
		if err != nil {
			app.Logger.Error("Unable to retrieve user ID from session", err)
			http.Error(w, "Unable to retrieve user ID", http.StatusUnauthorized)
			return
		}

		story := &data.Story{
			Title:    r.FormValue("title"),
			Content:  r.FormValue("content"),
			Language: r.FormValue("language"),
			Location: r.FormValue("location"),
			Category: r.FormValue("category"),
			UserID:   userID, // Use the retrieved user ID
		}

		v := validator.New()
		// Assuming data.ValidateStory is a custom function for validation
		data.ValidateStory(v, story)

		if !v.Valid() {
			http.Error(w, "Validation failed", http.StatusBadRequest)
			return
		}

		if err := app.Stories.Insert(story); err != nil {
			app.Logger.Error("Unable to create story", err)
			http.Error(w, fmt.Sprintf("Unable to create story: %s", err), http.StatusInternalServerError)
			return
		}

		// Ensure the story has been inserted correctly before using `story.ID`
		http.Redirect(w, r, fmt.Sprintf("/stories/%s", story.ID), http.StatusSeeOther)
		return
	}

	err := app.Templates.Render(w, "create_story.tmpl", nil)
	if err != nil {
		app.Logger.Error("Could not render create story page", err)
		http.Error(w, "Could not render create story page", http.StatusInternalServerError)
	}
}

// EditStory handles story editing.
func EditStory(app *internal.Application, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		storyID := r.FormValue("id")
		userID, err := app.Session.GetUserID(r) // Correctly using GetUserID on app.Session
		if err != nil {
			app.Logger.Error("Unable to retrieve user ID from session", err)
			http.Error(w, "Unable to retrieve user ID", http.StatusUnauthorized)
			return
		}

		story := &data.Story{
			ID:       storyID,
			Title:    r.FormValue("title"),
			Content:  r.FormValue("content"),
			Language: r.FormValue("language"),
			Location: r.FormValue("location"),
			Category: r.FormValue("category"),
			UserID:   userID, // Use the retrieved user ID
		}

		v := validator.New()
		data.ValidateStory(v, story)

		if !v.Valid() {
			http.Error(w, "Validation failed", http.StatusBadRequest)
			return
		}

		if err := app.Stories.Update(story); err != nil {
			app.Logger.Error("Unable to update story", err)
			http.Error(w, fmt.Sprintf("Unable to update story: %s", err), http.StatusInternalServerError)
			return
		}

		// Corrected redirect after successful story update
		http.Redirect(w, r, fmt.Sprintf("/stories/%s", storyID), http.StatusSeeOther)
		return
	}

	storyID := chi.URLParam(r, "id")
	story, err := app.Stories.GetByID(storyID)
	if err != nil {
		app.Logger.Error("Unable to find story", err)
		http.Error(w, fmt.Sprintf("Unable to find story: %s", err), http.StatusNotFound)
		return
	}

	userID, err := app.Session.GetUserID(r) // Correctly using GetUserID on app.Session
	if err != nil {
		app.Logger.Error("Unable to retrieve user ID from session", err)
		http.Error(w, "Unable to retrieve user ID", http.StatusUnauthorized)
		return
	}

	if story.UserID != userID {
		http.Error(w, "You are not authorized to edit this story", http.StatusForbidden)
		return
	}

	err = app.Templates.Render(w, "edit_story.tmpl", map[string]interface{}{"Story": story})
	if err != nil {
		app.Logger.Error("Could not render edit story page", err)
		http.Error(w, "Could not render edit story page", http.StatusInternalServerError)
	}
}

// DeleteStory handles story deletion.
func DeleteStory(app *internal.Application, w http.ResponseWriter, r *http.Request) {
	storyID := chi.URLParam(r, "id")
	story, err := app.Stories.GetByID(storyID)
	if err != nil {
		app.Logger.Error("Unable to find story", err)
		http.Error(w, fmt.Sprintf("Unable to find story: %s", err), http.StatusNotFound)
		return
	}

	userID, err := app.Session.GetUserID(r) // Correctly using GetUserID on app.Session
	if err != nil {
		app.Logger.Error("Unable to retrieve user ID from session", err)
		http.Error(w, "Unable to retrieve user ID", http.StatusUnauthorized)
		return
	}

	if story.UserID != userID {
		http.Error(w, "You are not authorized to delete this story", http.StatusForbidden)
		return
	}

	if err := app.Stories.Delete(storyID); err != nil {
		app.Logger.Error("Unable to delete story", err)
		http.Error(w, fmt.Sprintf("Unable to delete story: %s", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/stories", http.StatusSeeOther)
}
