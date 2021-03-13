package page

import (
	"time"

	"github.com/nickelghost/yuzu-cms/app/models/post"
)

type Page struct {
	ID           int
	PostID       int
	Index        int
	Slug         string
	InNavigation bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Post         post.Post
}
