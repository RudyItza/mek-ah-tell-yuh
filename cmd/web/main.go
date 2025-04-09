package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/RudyItza/mek-ah-tell-yuh/internal"
	"github.com/RudyItza/mek-ah-tell-yuh/internal/data"
	"github.com/RudyItza/mek-ah-tell-yuh/internal/render"
	"github.com/RudyItza/mek-ah-tell-yuh/routes"
	"github.com/gorilla/sessions"
)

func main() {
	// Logger initialization
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Database initialization
	dataSourceName := "postgres://mekuser:folklore@localhost/mekahtellyuh?sslmode=disable"
	db, err := data.InitializeDB(dataSourceName)
	if err != nil {
		logger.Error("Failed to connect to the database: " + err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// Template manager initialization
	tmplMgr, err := render.NewTemplateManager("ui/templates/*.tmpl")
	if err != nil {
		logger.Error("Failed to load templates: " + err.Error())
		os.Exit(1)
	}

	// Initialize session store using gorilla sessions
	sessionStore := sessions.NewCookieStore([]byte("your-secret-key"))

	// Application initialization
	app := &internal.Application{
		Logger:    logger,
		Feedback:  &data.FeedbackModel{DB: db},
		Templates: tmplMgr,
		Session:   &internal.Session{Store: sessionStore}, // Set session store here
		Stories:   &data.StoryModel{DB: db},
	}

	// Initialize routes
	mux := routes.Routes(app)

	logger.Info("Starting server on :4000")
	err = http.ListenAndServe(":4000", mux)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
