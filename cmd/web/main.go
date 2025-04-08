// cmd/main.go
package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/RudyItza/mek-ah-tell-yuh/internal"
	"github.com/RudyItza/mek-ah-tell-yuh/internal/data"
	"github.com/RudyItza/mek-ah-tell-yuh/routes"
)

func main() {
	// Set up the logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Initialize the Application struct with the logger and feedback model
	app := &internal.Application{
		Logger:   logger,
		Feedback: &data.FeedbackModel{}, // Adjust according to your actual data model
	}

	// Set up routes and pass the application struct
	mux := routes.Routes(app) // Passing app to the Routes function
	logger.Info("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
