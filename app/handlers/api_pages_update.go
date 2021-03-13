package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	pageModel "github.com/nickelghost/yuzu-cms/app/models/page"
)

// APIPagesUpdate updates a page
func (hs Handlers) APIPagesUpdate(c echo.Context) error {
	type Request struct {
		PostID         int    `json:"post_id" validate:"required,min=1"`
		PositionChange int    `json:"position_change"`
		Slug           string `json:"slug" validate:"required,min=1"`
		InNavigation   bool   `json:"in_navigation"`
	}
	type Response struct {
		ID           int       `json:"id"`
		Index        int       `json:"index"`
		Slug         string    `json:"slug"`
		InNavigation bool      `json:"in_navigation"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
	req := Request{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = validator.New().Struct(req)
	if err != nil {
		return err
	}
	page, err := pageModel.GetByID(hs.DB, id)
	if err != nil {
		return err
	}
	err = page.Update(
		hs.DB,
		req.PostID,
		page.Index+req.PositionChange,
		req.Slug,
		req.InNavigation,
	)
	if err != nil {
		return err
	}
	res := Response{
		ID:           page.ID,
		Index:        page.Index,
		Slug:         page.Slug,
		InNavigation: page.InNavigation,
		CreatedAt:    page.CreatedAt,
		UpdatedAt:    page.UpdatedAt,
	}
	return c.JSON(http.StatusOK, &res)
}
