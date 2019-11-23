package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/nickelghost/cms/common"
	"github.com/nickelghost/cms/models"
)

type APIPostsIndexResponseItem struct {
	ID        uint
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func APIPostsIndex(c echo.Context) error {
	db := c.(*common.CustomContext).DB
	var posts []models.Post
	db.Select([]string{
		"id",
		"title",
		"created_at",
		"updated_at",
	}).Find(&posts)
	var res []APIPostsIndexResponseItem
	for _, post := range posts {
		resPost := APIPostsIndexResponseItem{
			ID:        post.ID,
			Title:     post.Title,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}
		res = append(res, resPost)
	}
	return c.JSON(http.StatusOK, res)
}
