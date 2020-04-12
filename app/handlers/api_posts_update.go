package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	postModel "github.com/nickelghost/yuzu-cms/app/models/post"
)

// APIPostsUpdateRequest represents a request for updating a post
type APIPostsUpdateRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"min=10"`
	IsDraft bool   `json:"is_draft"`
}

// APIPostsUpdateResponse represents a response of an updated post
type APIPostsUpdateResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	IsDraft   bool      `json:"is_draft"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// APIPostsUpdate updates a post over JSON
func (hs Handlers) APIPostsUpdate(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	req := new(APIPostsUpdateRequest)
	err = c.Bind(req)
	if err != nil {
		return err
	}
	err = validator.New().Struct(req)
	if err != nil {
		return err
	}
	post, err := postModel.Update(
		hs.DB,
		id,
		req.Title,
		req.Content,
		req.IsDraft,
	)
	res := APIPostsUpdateResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		IsDraft:   post.IsDraft,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
