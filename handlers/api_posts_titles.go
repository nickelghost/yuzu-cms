package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nickelghost/yuzu-cms/models"
)

func (hs Handlers) APIPostsTitles(c echo.Context) error {
	type ResponseItem struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}
	posts, err := (models.Post{}).GetAll(hs.DB)
	if err != nil {
		return err
	}
	res := []ResponseItem{}
	for _, post := range *posts {
		resItem := ResponseItem{
			ID:    int(post.ID),
			Title: post.Title,
		}
		res = append(res, resItem)
	}
	return c.JSON(http.StatusOK, &res)
}
