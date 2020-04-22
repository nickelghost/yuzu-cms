package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	postModel "github.com/nickelghost/yuzu-cms/app/models/post"
)

func (hs Handlers) APIPostsDelete(c echo.Context) error {
	type Response struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		IsDraft   bool      `json:"is_draft"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	post, err := postModel.Delete(hs.DB, id)
	if err != nil {
		return err
	}
	res := Response{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		IsDraft:   post.IsDraft,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	return c.JSON(http.StatusOK, res)
}
