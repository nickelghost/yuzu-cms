package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/nickelghost/cms/models"
)

// Homepage renders the index page
func (hs Handlers) Homepage(c echo.Context) error {
	sql, err := ioutil.ReadFile("queries/homepage.sql")
	if err != nil {
		return err
	}
	rows, err := hs.DB.Query(string(sql))
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
