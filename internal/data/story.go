package data

import (
	"database/sql"
	"fmt"

	"github.com/RudyItza/mek-ah-tell-yuh/internal/validator" // Import your validator package
)

type Story struct {
	ID       string
	Title    string
	Content  string
	Language string
	Location string
	Category string
	UserID   string
}

type StoryModel struct {
	DB *sql.DB
}

// Insert a new story into the database
func (m *StoryModel) Insert(story *Story) error {
	query := `
        INSERT INTO stories (title, content, language, location, category, user_id)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`
	return m.DB.QueryRow(query, story.Title, story.Content, story.Language, story.Location, story.Category, story.UserID).Scan(&story.ID)
}

// Update an existing story in the database
func (m *StoryModel) Update(story *Story) error {
	query := `
        UPDATE stories
        SET title = $1, content = $2, language = $3, location = $4, category = $5
        WHERE id = $6 AND user_id = $7`
	result, err := m.DB.Exec(query, story.Title, story.Content, story.Language, story.Location, story.Category, story.ID, story.UserID)
	if err != nil {
		return fmt.Errorf("unable to update story: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("unable to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no story found to update")
	}
	return nil
}

// Get a story by its ID
func (m *StoryModel) GetByID(id string) (*Story, error) {
	query := `SELECT id, title, content, language, location, category, user_id FROM stories WHERE id = $1`
	story := &Story{}
	err := m.DB.QueryRow(query, id).Scan(&story.ID, &story.Title, &story.Content, &story.Language, &story.Location, &story.Category, &story.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("story not found")
		}
		return nil, fmt.Errorf("unable to retrieve story: %w", err)
	}
	return story, nil
}

// Delete a story from the database
func (m *StoryModel) Delete(id string) error {
	query := `DELETE FROM stories WHERE id = $1`
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("unable to delete story: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("unable to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no story found to delete")
	}
	return nil
}

// ValidateStory validates the fields of a Story
func ValidateStory(v *validator.Validator, story *Story) {
	// Check that the title is not empty
	v.Check(story.Title != "", "title", "Title cannot be blank")
	// Check that the content is not empty
	v.Check(story.Content != "", "content", "Content cannot be blank")
	// Check that the language is not empty
	v.Check(story.Language != "", "language", "Language cannot be blank")
	// Check that the location is not empty
	v.Check(story.Location != "", "location", "Location cannot be blank")
	// Check that the category is not empty
	v.Check(story.Category != "", "category", "Category cannot be blank")
}
