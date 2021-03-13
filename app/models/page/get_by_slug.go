package page

import "database/sql"

// GetBySlug gets the page with the slug specified from the DB
func GetBySlug(conn *sql.DB, slug string) (*Page, error) {
	page := Page{}
	err := conn.QueryRow("SELECT * FROM pages WHERE slug = $1", slug).Scan(
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
