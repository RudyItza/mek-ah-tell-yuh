package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	Passwordhash string    `json:"-"` // Do not expose password in JSON
	CreatedAt    time.Time `json:"created_at"`
}

type UserModel struct {
	DB *sql.DB
}

// Insert hashes the user's password and stores it in the database
func (m *UserModel) Insert(user *User) error {
	// Hash the password before inserting
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Passwordhash), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("could not hash password: %w", err)
	}

	// Insert query using hashed password
	query := `
        INSERT INTO users (username, password_hash)
        VALUES ($1, $2)
        RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Execute the query and scan the result
	return m.DB.QueryRowContext(
		ctx,
		query,
		user.Username,
		hashedPassword, // Store the hashed password
	).Scan(&user.ID, &user.CreatedAt)
}

// GetByEmail retrieves a user by their email from the database
func (m *UserModel) GetByEmail(email string) (*User, error) {
	query := `
        SELECT id, username, password_hash, created_at
        FROM users
        WHERE username = $1` // Use 'username' or 'email' if applicable
	row := m.DB.QueryRow(query, email)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Passwordhash, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, fmt.Errorf("could not get user by email: %w", err)
	}

	return &user, nil
}
