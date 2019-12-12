package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nickelghost/cms/database"
	"github.com/nickelghost/cms/models"
	"github.com/nickelghost/cms/other"
)

// Homepage renders the index page
func Homepage(c echo.Context) error {
	db := c.(*other.CustomContext).DB
	sql, err := database.GetSQL("homepage")
	if err != nil {
		return err
	}
	rows, err := db.Query(sql)
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
