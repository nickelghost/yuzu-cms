package controllers

import (
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/labstack/echo"
	"github.com/nickelghost/cms/models"
	"github.com/nickelghost/cms/other"
)

// APIPostsIndex fetches all posts that were added
func APIPostsIndex(c echo.Context) error {
	db := c.(*other.CustomContext).DB
	sql, _, err := sq.Select(
		"id", "title", "created_at", "updated_at",
	).From("posts").ToSql()
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
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return err
		}
		posts = append(posts, post)
	}
	return c.JSON(http.StatusOK, posts)
}
