package page

import "database/sql"

// GetByID retrieves a page with the id specified from the DB
func GetByID(conn *sql.DB, id int) (*Page, error) {
	page := Page{}
	err := conn.QueryRow("SELECT * FROM pages WHERE id = $1", id).Scan(
		&page.ID,
		&page.PostID,
		&page.Index,
		&page.Slug,
		&page.InNavigation,
		&page.CreatedAt,
		&page.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &page, nil
}
