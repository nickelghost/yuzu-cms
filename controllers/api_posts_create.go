package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/nickelghost/cms/common"
	"github.com/nickelghost/cms/models"
)

type APIPostsCreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type APIPostsCreateResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func APIPostsCreate(c echo.Context) error {
	req := new(APIPostsCreateRequest)
	err := c.Bind(req)
	if err != nil {
		log.Println("yup")
		return err
	}
	db := c.(*common.CustomContext).DB
	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
	}
	err = db.Create(&post).Error
	if err != nil {
		return err
	}
	res := APIPostsCreateResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	return c.JSON(http.StatusOK, res)
}
