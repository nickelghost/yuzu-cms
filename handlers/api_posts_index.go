package handlers

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/nickelghost/cms/db"
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
	sql, err := ioutil.ReadFile("queries/api_posts_index.sql")
	if err != nil {
		return err
	}
	rows, err := db.Conn.Query(string(sql))
	if err != nil {
		return err
	}
	var res []APIPostsIndexResponseItem
	for rows.Next() {
		resItem := APIPostsIndexResponseItem{}
		err := rows.Scan(
			&resItem.ID,
			&resItem.Title,
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
