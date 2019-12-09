package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/nickelghost/cms/models"
	"github.com/nickelghost/cms/other"
)

type APIPostsIndexResponseItem struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func APIPostsIndex(c echo.Context) error {
	db := c.(*other.CustomContext).DB
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
