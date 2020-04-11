package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	"github.com/nickelghost/yuzu-cms/src/models"
)

func (hs Handlers) Page(c echo.Context) error {
	page, err := (models.Page{}).GetBySlug(hs.DB, c.Param("slug"))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.String(http.StatusNotFound, "Not found")
		}
		return err
	}
	post, err := (models.Post{}).GetByID(hs.DB, page.PostID)
	if err != nil {
		return err
	}
	pages, err := (models.Page{}).GetAll(hs.DB, true)
	if err != nil {
		return err
	}
	return c.Render(
		http.StatusOK,
		"page.html",
		map[string]interface{}{"post": post, "pages": pages},
	)
}
