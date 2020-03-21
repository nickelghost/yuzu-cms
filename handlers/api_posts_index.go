package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/nickelghost/yuzu-cms/models"
)

// APIPostsIndexResponseItem represents an array item for the response
type APIPostsIndexResponseItem struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	ContentPreview string    `json:"content_preview"`
	IsDraft        bool      `json:"is_draft"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// APIPostsIndex fetches all posts that were added
func (hs Handlers) APIPostsIndex(c echo.Context) error {
	rows, err := hs.DB.Query(hs.SQL["api_posts_index.sql"])
	if err != nil {
		return err
	}
	res := make([]APIPostsIndexResponseItem, 0)
	for rows.Next() {
		// TODO: Scan model instead and move preview creation to the model
		post := models.Post{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.IsDraft,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return err
		}
		post.PopulateContentPreview()
		resItem := APIPostsIndexResponseItem{
			ID:             post.ID,
			Title:          post.Title,
			ContentPreview: post.ContentPreview,
			IsDraft:        post.IsDraft,
			CreatedAt:      post.CreatedAt,
			UpdatedAt:      post.UpdatedAt,
		}
		res = append(res, resItem)
	}
	return c.JSON(http.StatusOK, res)
}
