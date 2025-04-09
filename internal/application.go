package internal

import (
	"log/slog"

	"github.com/RudyItza/mek-ah-tell-yuh/internal/data"
	"github.com/RudyItza/mek-ah-tell-yuh/internal/render"
)

// Application holds application-wide dependencies.
type Application struct {
	Logger    *slog.Logger
	Feedback  *data.FeedbackModel
	Templates *render.TemplateManager
	Session   *Session // Use your custom Session type here, not *sessions.CookieStore
	Stories   *data.StoryModel
	Users     *data.UserModel
}
