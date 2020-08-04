package post

import "database/sql"

// Delete removes a post from the database and returns it
func Delete(conn *sql.DB, id int) (Post, error) {
	post := Post{}
	err := conn.QueryRow(`DELETE FROM posts WHERE id = $1 RETURNING *`, id).Scan(
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
	return post, err
}
