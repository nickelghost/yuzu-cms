package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/nickelghost/cms/db"
	"github.com/nickelghost/cms/models"
)

type requestAPIPostsCreate struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"min=10"`
}

// APIPostsCreate adds a new post for the website
func APIPostsCreate(c echo.Context) error {
	req := new(requestAPIPostsCreate)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	err = validator.New().Struct(req)
	if err != nil {
		return err
	}
	post := models.Post{}
	sql, err := ioutil.ReadFile("queries/api_posts_create.sql")
	if err != nil {
		return err
	}
	err = db.Conn.QueryRow(
		string(sql),
		req.Title,
		req.Content,
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
