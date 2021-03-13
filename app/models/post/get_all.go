package post

import "database/sql"

// GetAll gets all posts that match a criteria
// isDraft doesn't look at the is_draft criteria if set to -1, only finds
// non-drafts if set to 0 and only rafts if set to 1
func GetAll(conn *sql.DB, isDraft int) (*[]Post, error) {
	var rows *sql.Rows
	var err error
	if isDraft < 0 {
		rows, err = conn.Query(`
            SELECT * FROM posts
            ORDER BY created_at DESC
        `)
	} else {
		rows, err = conn.Query(`
            SELECT * FROM posts
            WHERE is_draft = $1
            ORDER BY created_at DESC
        `, isDraft)
	}
	if err != nil {
		return nil, err
	}
	posts := []Post{}
	for rows.Next() {
		post := Post{}
		err := rows.Scan(
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
		posts = append(posts, post)
	}
	return &posts, nil
}
