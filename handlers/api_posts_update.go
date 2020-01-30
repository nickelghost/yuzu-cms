package handlers

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
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
	req := new(APIPostsUpdateRequest)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	err = validator.New().Struct(req)
	if err != nil {
		return err
	}
	res := new(APIPostsUpdateResponse)
	err = hs.DB.QueryRow(
		hs.SQL["api_posts_update.sql"],
		c.Param("id"),
		req.Title,
		req.Content,
		req.IsDraft,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Content,
		&res.IsDraft,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
