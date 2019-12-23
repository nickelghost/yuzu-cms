package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nickelghost/cms/db"
	"github.com/nickelghost/cms/models"
)

// Homepage renders the index page
func Homepage(c echo.Context) error {
	sql, err := db.GetSQL("homepage")
	if err != nil {
		return err
	}
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return err
	}
	var posts []models.Post
	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return err
		}
		posts = append(posts, post)
	}
	return c.Render(
		http.StatusOK,
		"homepage.html",
		map[string]interface{}{"posts": posts},
	)
}
