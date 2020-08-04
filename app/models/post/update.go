package post

import "database/sql"

// Update the DB record for a post
func Update(
	conn *sql.DB,
	id int,
	title string,
	content string,
	slug string,
	isDraft bool,
) (Post, error) {
	post := Post{}
	err := conn.QueryRow(`
        UPDATE posts
        SET title = $2, content = $3, slug = $4, is_draft = $5, updated_at = NOW()
        WHERE id = $1
        RETURNING *
        `,
		id,
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
