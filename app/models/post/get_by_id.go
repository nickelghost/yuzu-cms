package post

import "database/sql"

// GetByID returns a single post with the ID specified
func GetByID(conn *sql.DB, id int) (*Post, error) {
	post := Post{}
	err := conn.QueryRow(
		`SELECT * FROM posts WHERE id = $1`,
		id,
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
