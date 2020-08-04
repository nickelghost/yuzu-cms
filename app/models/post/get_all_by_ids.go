package post

import (
	"database/sql"
	"fmt"
	"strings"
)

// GetAllByIds gets only the posts with IDs specified
func GetAllByIds(conn *sql.DB, ids []int) (*[]Post, error) {
	// Create a SQL param string from the IDs provided
	idsString := strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")
	param := "{" + idsString + "}"
	rows, err := conn.Query("SELECT * FROM posts WHERE id = ANY($1::int[])", param)
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
