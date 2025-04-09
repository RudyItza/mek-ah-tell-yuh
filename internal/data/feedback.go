package data

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/RudyItza/mek-ah-tell-yuh/internal/validator" // Correct import path
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

// Insert inserts a new feedback into the database and returns the ID and created date.
func (m *FeedbackModel) Insert(feedback *Feedback) error {
	query := `
		INSERT INTO feedback (fullname, subject, message)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`
	err := m.DB.QueryRow(query, feedback.FullName, feedback.Subject, feedback.Message).
		Scan(&feedback.ID, &feedback.CreatedAt)
	if err != nil {
		return fmt.Errorf("unable to insert feedback: %w", err)
	}
	return nil
}

// ValidateFeedback validates the feedback input.
func ValidateFeedback(v *validator.Validator, feedback *Feedback) {
	v.Check(feedback.FullName != "", "fullname", "must be provided")
	v.Check(feedback.Subject != "", "subject", "must be provided")
	v.Check(feedback.Message != "", "message", "must be provided")
}
