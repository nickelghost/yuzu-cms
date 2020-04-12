package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/nickelghost/yuzu-cms/app/models"
)

func (hs Handlers) APIPagesUpdate(c echo.Context) error {
	type Request struct {
		PostID         int     `json:"post_id"`
		PositionChange int     `json:"position_change"`
		Slug           string  `json:"slug"`
		InNavigation   bool    `json:"in_navigation"`
		Heading        *string `json:"heading"`
	}
	req := Request{}
	c.Bind(&req)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	page, err := (models.Page{}).GetById(hs.DB, id)
	if err != nil {
		return err
	}
	err = page.Update(
		hs.DB,
		req.PostID,
		page.Index+req.PositionChange,
		req.Slug,
		req.InNavigation,
		req.Heading,
	)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &page)
}
