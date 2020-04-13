package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	pageModel "github.com/nickelghost/yuzu-cms/app/models/page"
)

func (hs Handlers) APIPagesDelete(c echo.Context) error {
	type Response struct {
		ID           int       `json:"id"`
		Index        int       `json:"index"`
		Slug         string    `json:"slug"`
		InNavigation bool      `json:"in_navigation"`
		Heading      *string   `json:"heading"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	page, err := pageModel.GetById(hs.DB, id)
	if err != nil {
		return err
	}
	err = page.Delete(hs.DB)
	if err != nil {
		return err
	}
	res := Response{
		ID:           page.ID,
		Index:        page.Index,
		Slug:         page.Slug,
		InNavigation: page.InNavigation,
		Heading:      page.Heading,
		CreatedAt:    page.CreatedAt,
		UpdatedAt:    page.UpdatedAt,
	}
	return c.JSON(http.StatusOK, res)
}
