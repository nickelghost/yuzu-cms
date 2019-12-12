package controllers

import (
	"net/http"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/labstack/echo"
	"github.com/nickelghost/cms/models"
	"github.com/nickelghost/cms/other"
)

type requestAPIPostsCreate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// APIPostsCreate adds a new post for the website
func APIPostsCreate(c echo.Context) error {
	req := new(requestAPIPostsCreate)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	db := c.(*other.CustomContext).DB
	post := models.Post{}
	sql, _, err := sq.
		Insert("posts").
		Columns("title", "content", "created_at", "updated_at").
		Values("$1", "$2", "$3", "$4").
		Suffix("RETURNING id, title, content, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	err = db.QueryRow(
		sql,
		req.Title,
		req.Content,
		time.Now(),
		time.Now(),
	).Scan(
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
