package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nickelghost/yuzu-cms/src/models"
)

// Homepage renders the index page
func (hs Handlers) Homepage(c echo.Context) error {
	rows, err := hs.DB.Query(hs.SQL["homepage.sql"])
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
	pages, err := (models.Page{}).GetAll(hs.DB, true)
	if err != nil {
		return err
	}
	return c.Render(
		http.StatusOK,
		"homepage.html",
		map[string]interface{}{"posts": posts, "pages": pages},
	)
}
