package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/nickelghost/yuzu-cms/src/models"
)

func (hs Handlers) APIPagesIndex(c echo.Context) error {
	type ResponseItemPost struct {
		ID             uint      `json:"id"`
		Title          string    `json:"title"`
		ContentPreview string    `json:"content_preview"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}
	type ResponseItem struct {
		ID           int              `json:"id"`
		PostID       int              `json:"post_id"`
		Index        int              `json:"index"`
		Slug         string           `json:"slug"`
		InNavigation bool             `json:"in_navigation"`
		Heading      *string          `json:"heading"`
		CreatedAt    time.Time        `json:"created_at"`
		UpdatedAt    time.Time        `json:"updated_at"`
		Post         ResponseItemPost `json:"post"`
	}
	pages, err := (models.Page{}).GetAll(hs.DB, true)
	if err != nil {
		return err
	}
	res := []ResponseItem{}
	for _, page := range *pages {
		page.Post.PopulateContentPreview()
		item := ResponseItem{
			ID:           page.ID,
			PostID:       page.PostID,
			Index:        page.Index,
			Slug:         page.Slug,
			InNavigation: page.InNavigation,
			Heading:      page.Heading,
			CreatedAt:    page.CreatedAt,
			UpdatedAt:    page.UpdatedAt,
			Post: ResponseItemPost{
				ID:             page.Post.ID,
				Title:          page.Post.Title,
				ContentPreview: page.Post.ContentPreview,
				CreatedAt:      page.Post.CreatedAt,
				UpdatedAt:      page.Post.UpdatedAt,
			},
		}
		res = append(res, item)
	}
	return c.JSON(http.StatusOK, res)
}
