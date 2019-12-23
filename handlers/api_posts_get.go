package handlers

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/nickelghost/cms/db"
)

// APIPostsGetResponse represents a response of a requested post
type APIPostsGetResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// APIPostsGet fetches a single post by its id
func APIPostsGet(c echo.Context) error {
	sql, err := ioutil.ReadFile("queries/api_posts_get.sql")
	if err != nil {
		return err
	}
	res := new(APIPostsGetResponse)
	err = db.Conn.QueryRow(string(sql), c.Param("id")).Scan(
		&res.ID,
		&res.Title,
		&res.Content,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
