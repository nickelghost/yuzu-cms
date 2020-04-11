package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// APIPostsGetResponse represents a response of a requested post
type APIPostsGetResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	IsDraft   bool      `json:"is_draft"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// APIPostsGet fetches a single post by its id
func (hs Handlers) APIPostsGet(c echo.Context) error {
	res := APIPostsGetResponse{}
	err := hs.DB.QueryRow(hs.SQL["api_posts_get.sql"], c.Param("id")).Scan(
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
