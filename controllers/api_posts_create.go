package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/nickelghost/cms/database"
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
	sql, err := database.GetSQL("api_posts_create")
	if err != nil {
		return err
	}
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
