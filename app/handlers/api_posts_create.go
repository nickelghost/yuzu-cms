package handlers

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	postModel "github.com/nickelghost/yuzu-cms/app/models/post"
)

// APIPostsCreate adds a new post for the website
func (hs Handlers) APIPostsCreate(c echo.Context) error {
	type Request struct {
		Title   string `json:"title" validate:"required,min=1"`
		Content string `json:"content" validate:"min=10"`
		Slug    string `json:"slug" validate:"required,min=1"`
		IsDraft bool   `json:"is_draft"`
	}
	type Response struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		IsDraft   bool      `json:"is_draft"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	req := new(Request)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	err = validator.New().Struct(req)
	if err != nil {
		return err
	}
	post, err := postModel.Create(
		hs.DB,
		req.Title,
		req.Content,
		req.Slug,
		req.IsDraft,
	)
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
