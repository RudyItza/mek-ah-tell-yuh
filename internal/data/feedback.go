package data

import (
	"database/sql"
	"time"

	"github.com/RudyItza/mek-ah-tell-yuh/internal/validator"
)

type Feedback struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"fullname"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type FeedbackModel struct {
	DB *sql.DB
}

// Insert a new feedback
func (m *FeedbackModel) Insert(feedback *Feedback) error {
	query := `
		INSERT INTO feedback (fullname, subject, message)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`
	return m.DB.QueryRow(query, feedback.FullName, feedback.Subject, feedback.Message).
		Scan(&feedback.ID, &feedback.CreatedAt)
}

// Validation function
func ValidateFeedback(v *validator.Validator, feedback *Feedback) {
	v.Check(feedback.FullName != "", "fullname", "must be provided")
	v.Check(feedback.Subject != "", "subject", "must be provided")
	v.Check(feedback.Message != "", "message", "must be provided")
}
