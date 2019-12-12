package controllers

import (
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/labstack/echo"
	"github.com/nickelghost/cms/models"
	"github.com/nickelghost/cms/other"
)

// APIPostsGet fetches a single post by its id
func APIPostsGet(c echo.Context) error {
	db := c.(*other.CustomContext).DB
	sql, _, err := sq.Select(
		"id", "title", "content", "created_at", "updated_at",
	).From("posts").Where("id = $1").ToSql()
	if err != nil {
		return err
	}
	post := new(models.Post)
	err = db.QueryRow(sql, c.Param("id")).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, post)
}
