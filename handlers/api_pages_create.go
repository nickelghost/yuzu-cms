package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/nickelghost/yuzu-cms/models"
)

func (hs Handlers) APIPagesCreate(c echo.Context) error {
	type Request struct {
		PostID       int     `json:"post_id"`
		Index        int     `json:"index"`
		Slug         string  `json:"slug"`
		InNavigation bool    `json:"in_navigation"`
		Heading      *string `json:"heading"`
	}
	type Response struct {
		ID           int       `json:"id"`
		Index        int       `json:"index"`
		Slug         string    `json:"slug"`
		InNavigation bool      `json:"in_navigation"`
		Heading      *string   `json:"heading"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
	req := Request{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	page, err := (models.Page{}).Create(
		hs.DB,
		req.PostID,
		req.Index,
		req.Slug,
		req.InNavigation,
		req.Heading,
	)
	res := Response{
		ID:           page.ID,
		Index:        page.Index,
		Slug:         page.Slug,
		InNavigation: page.InNavigation,
		Heading:      page.Heading,
		CreatedAt:    page.CreatedAt,
		UpdatedAt:    page.UpdatedAt,
	}
	return c.JSON(http.StatusOK, &res)
}
