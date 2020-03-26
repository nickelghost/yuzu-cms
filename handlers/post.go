package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/nickelghost/yuzu-cms/models"
)

func (hs Handlers) Post(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	post, err := (models.Post{}).GetByID(hs.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.String(http.StatusNotFound, "Not found")
		}
		return err
	}
	pages, err := (models.Page{}).GetAll(hs.DB, true)
	if err != nil {
		return err
	}
	return c.Render(
		http.StatusOK,
		"post.html",
		map[string]interface{}{"post": post, "pages": pages},
	)
}
