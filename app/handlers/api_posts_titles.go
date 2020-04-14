package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	postModel "github.com/nickelghost/yuzu-cms/app/models/post"
)

func (hs Handlers) APIPostsTitles(c echo.Context) error {
	type ResponseItem struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}
	posts, err := postModel.GetAll(hs.DB, -1)
	if err != nil {
		return err
	}
	res := []ResponseItem{}
	for _, post := range *posts {
		resItem := ResponseItem{
			ID:    post.ID,
			Title: post.Title,
		}
		res = append(res, resItem)
	}
	return c.JSON(http.StatusOK, &res)
}
