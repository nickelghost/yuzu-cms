package page

import "database/sql"

// Create inserts a new page to the DB
func Create(
	conn *sql.DB,
	postID int,
	index int,
	slug string,
	inNavigation bool,
) (*Page, error) {
	page := Page{}
	err := conn.QueryRow(`
        INSERT INTO
        pages (post_id, index, slug, in_navigation, created_at, updated_at)
        VALUES
        ($1, $2, $3, $4, NOW(), NOW())
        RETURNING *
        `,
		postID,
		index,
		slug,
		inNavigation,
	).Scan(
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
