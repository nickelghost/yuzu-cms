package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	postsModel "github.com/nickelghost/yuzu-cms/app/models/post"
)

// APIPostsGetResponse represents a response of a requested post
type APIPostsGetResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	IsDraft   bool      `json:"is_draft"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// APIPostsGet fetches a single post by its id
func (hs Handlers) APIPostsGet(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	post, err := postsModel.GetByID(hs.DB, id)
	if err != nil {
		return err
	}
	res := APIPostsGetResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		IsDraft:   post.IsDraft,
		Slug:      post.Slug,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
