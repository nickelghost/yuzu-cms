package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	postModel "github.com/nickelghost/yuzu-cms/app/models/post"
)

// APIPostsIndexResponseItem represents an array item for the response
type APIPostsIndexResponseItem struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	ContentPreview string    `json:"content_preview"`
	IsDraft        bool      `json:"is_draft"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// APIPostsIndex fetches all posts that were added
func (hs Handlers) APIPostsIndex(c echo.Context) error {
	draftOnly, err := strconv.ParseBool(c.QueryParam("draft"))
	if err != nil {
		return err
	}
	var drafts int
	if draftOnly {
		drafts = 1
	} else {
		drafts = -1
	}
	posts, err := postModel.GetAll(hs.DB, drafts)
	if err != nil {
		return err
	}
	res := make([]APIPostsIndexResponseItem, 0)
	for _, post := range *posts {
		post.ContentPreview = post.GetContentPreview(80)
		resItem := APIPostsIndexResponseItem{
			ID:             post.ID,
			Title:          post.Title,
			ContentPreview: post.ContentPreview,
			IsDraft:        post.IsDraft,
			CreatedAt:      post.CreatedAt,
			UpdatedAt:      post.UpdatedAt,
		}
		res = append(res, resItem)
	}
	return c.JSON(http.StatusOK, res)
}
