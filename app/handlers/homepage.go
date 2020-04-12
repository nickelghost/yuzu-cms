package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nickelghost/yuzu-cms/app/models"
	postModel "github.com/nickelghost/yuzu-cms/app/models/post"
)

// Homepage renders the index page
func (hs Handlers) Homepage(c echo.Context) error {
	posts, err := postModel.GetAll(hs.DB, -1)
	if err != nil {
		return err
	}
	pages, err := (models.Page{}).GetAll(hs.DB, true)
	if err != nil {
		return err
	}
	return c.Render(
		http.StatusOK,
		"homepage.html",
		map[string]interface{}{"posts": posts, "pages": pages},
	)
}
