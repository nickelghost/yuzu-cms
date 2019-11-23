package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nickelghost/cms/common"
	"github.com/nickelghost/cms/models"
)

func APIPostsIndex(c echo.Context) error {
	db := c.(*common.CustomContext).DB
	var posts []models.Post
	db.Find(&posts)
	return c.JSON(http.StatusOK, posts)
}
