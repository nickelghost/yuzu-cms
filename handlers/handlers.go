package handlers

import "database/sql"

// Handlers provides common data for our handlers
type Handlers struct {
	DB  *sql.DB
	SQL map[string]string
}
