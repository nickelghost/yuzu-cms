package handlers

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	postModel "github.com/nickelghost/yuzu-cms/app/models/post"
)

// APIPostsCreateRequest represents a request for creating a post
type APIPostsCreateRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"min=10"`
	IsDraft bool   `json:"is_draft"`
}

// APIPostsCreateResponse represents a response of a created post
type APIPostsCreateResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	IsDraft   bool      `json:"is_draft"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// APIPostsCreate adds a new post for the website
func (hs Handlers) APIPostsCreate(c echo.Context) error {
	req := new(APIPostsCreateRequest)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	err = validator.New().Struct(req)
	if err != nil {
		return err
	}
	post, err := postModel.Create(hs.DB, req.Title, req.Content, req.IsDraft)
	if err != nil {
		return err
	}
	res := APIPostsCreateResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		IsDraft:   post.IsDraft,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	return c.JSON(http.StatusOK, res)
}
