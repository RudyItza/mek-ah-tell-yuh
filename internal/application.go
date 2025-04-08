// internal/application.go
package internal

import (
	"log/slog"

	"github.com/RudyItza/mek-ah-tell-yuh/internal/data"
)

// Application struct that holds the logger and feedback models
type Application struct {
	Logger   *slog.Logger        // Capitalized field to export it
	Feedback *data.FeedbackModel // Capitalized field to export it
}
