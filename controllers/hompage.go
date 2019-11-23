package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nickelghost/cms/common"
	"github.com/nickelghost/cms/models"
)

func Homepage(c echo.Context) error {
	db := c.(*common.CustomContext).DB
	var posts []models.Post
	db.Select([]string{
		"id",
		"title",
		"content_preview",
		"created_at",
		"updated_at",
	}).Find(&posts)
	return c.Render(
		http.StatusOK,
		"homepage.html",
		map[string]interface{}{"posts": posts},
	)
}
