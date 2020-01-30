package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// APIPostsIndexResponseItem represents an array item for the response
type APIPostsIndexResponseItem struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	IsDraft   bool      `json:"is_draft"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// APIPostsIndex fetches all posts that were added
func (hs Handlers) APIPostsIndex(c echo.Context) error {
	rows, err := hs.DB.Query(hs.SQL["api_posts_index.sql"])
	if err != nil {
		return err
	}
	res := make([]APIPostsIndexResponseItem, 0)
	for rows.Next() {
		resItem := APIPostsIndexResponseItem{}
		err := rows.Scan(
			&resItem.ID,
			&resItem.Title,
			&resItem.IsDraft,
			&resItem.CreatedAt,
			&resItem.UpdatedAt,
		)
		if err != nil {
			return err
		}
		res = append(res, resItem)
	}
	return c.JSON(http.StatusOK, res)
}
