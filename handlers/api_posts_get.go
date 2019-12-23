package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nickelghost/cms/db"
	"github.com/nickelghost/cms/models"
)

// APIPostsGet fetches a single post by its id
func APIPostsGet(c echo.Context) error {
	sql, err := db.GetSQL("api_posts_get")
	if err != nil {
		return err
	}
	post := new(models.Post)
	err = db.Conn.QueryRow(sql, c.Param("id")).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, post)
}
