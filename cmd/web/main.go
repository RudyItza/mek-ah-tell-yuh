// File: cmd/main.go
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

	// Database connection string
	dataSourceName := "postgres://mekuser:folklore@localhost/mekahtellyuh?sslmode=disable" // Replace with your actual DSN

	// Initialize the database
	db, err := data.InitializeDB(dataSourceName)
	if err != nil {
		logger.Error("Failed to connect to the database: " + err.Error())
		os.Exit(1)
	}
	defer db.Close() // Ensure the database connection is closed when the application exits

	// Initialize the Application struct with the logger and feedback model
	app := &internal.Application{
		Logger:   logger,
		Feedback: &data.FeedbackModel{DB: db}, // Pass the database connection to the FeedbackModel
	}

	// Set up routes and pass the application struct
	mux := routes.Routes(app) // Passing app to the Routes function
	logger.Info("Starting server on :4000")
	err = http.ListenAndServe(":4000", mux)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
