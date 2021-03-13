package post

import "database/sql"

// Create inserts a new post to the database
func Create(
	conn *sql.DB,
	title string,
	content string,
	slug string,
	isDraft bool,
) (Post, error) {
	post := Post{}
	err := conn.QueryRow(`
        INSERT INTO posts
            (title, content, slug, is_draft, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())
        RETURNING *
        `,
		title,
		content,
		slug,
		isDraft,
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
		return Post{}, err
	}
	return post, nil
}
