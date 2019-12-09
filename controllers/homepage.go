package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nickelghost/cms/models"
	"github.com/nickelghost/cms/other"
)

func Homepage(c echo.Context) error {
	db := c.(*other.CustomContext).DB
	var posts []models.Post
	db.Select([]string{
		"id",
		"title",
		"content",
		"created_at",
		"updated_at",
	}).Find(&posts)
	return c.Render(
		http.StatusOK,
		"homepage.html",
		map[string]interface{}{"posts": posts},
	)
}
