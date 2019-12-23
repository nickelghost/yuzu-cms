package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/nickelghost/cms/db"
	"github.com/nickelghost/cms/models"
)

// APIPostsIndexResponseItem represents an array item for the response
type APIPostsIndexResponseItem struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// APIPostsIndex fetches all posts that were added
func APIPostsIndex(c echo.Context) error {
	sql, err := db.GetSQL("api_posts_index")
	if err != nil {
		return err
	}
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return err
	}
	var res []APIPostsIndexResponseItem
	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return err
		}
		resItem := APIPostsIndexResponseItem{
			ID:        post.ID,
			Title:     post.Title,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}
		res = append(res, resItem)
	}
	return c.JSON(http.StatusOK, res)
}
