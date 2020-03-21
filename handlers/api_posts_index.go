package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	stripmd "github.com/writeas/go-strip-markdown"
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
		resItem := APIPostsIndexResponseItem{}
		err := rows.Scan(
			&resItem.ID,
			&resItem.Title,
			&resItem.ContentPreview,
			&resItem.IsDraft,
			&resItem.CreatedAt,
			&resItem.UpdatedAt,
		)
		if err != nil {
			return err
		}
		resItem.ContentPreview = stripmd.Strip(resItem.ContentPreview)
		if len(resItem.ContentPreview) > 80 {
			resItem.ContentPreview = resItem.ContentPreview[:80]
		}
		res = append(res, resItem)
	}
	return c.JSON(http.StatusOK, res)
}
