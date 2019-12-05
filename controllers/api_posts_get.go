package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/nickelghost/cms/common"
	"github.com/nickelghost/cms/models"
)

type APIPostsGetResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func APIPostsGet(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	post := new(models.Post)
	db := c.(*common.CustomContext).DB
	err = db.Find(&post, id).Error
	if err != nil {
		return err
	}
	res := APIPostsGetResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	return c.JSON(http.StatusOK, res)
}
