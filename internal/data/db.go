package data

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// InitializeDB initializes the database connection
func InitializeDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Check if the connection is established
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
