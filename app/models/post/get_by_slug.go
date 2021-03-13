package post

import "database/sql"

// GetBySlug fetches the post from the database by its URL slug
func GetBySlug(conn *sql.DB, slug string) (*Post, error) {
	post := Post{}
	err := conn.QueryRow(
		`SELECT * FROM posts WHERE slug = $1`,
		slug,
	).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Slug,
		&post.IsDraft,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
