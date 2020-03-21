package models

import (
	"time"

	stripmd "github.com/writeas/go-strip-markdown"
	"gopkg.in/russross/blackfriday.v2"
)

// Post contains articles and pages used by the CMS
type Post struct {
	ID             uint
	Title          string
	Content        string
	ContentPreview string
	IsDraft        bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// GetHTMLContent returns the post's markdown content as HTML
func (p *Post) GetHTMLContent() string {
	return string(blackfriday.Run([]byte(p.Content)))
}

// PopulateContentPreview strips markdown from the content string, trims it
// and saves ContentPreview
func (p *Post) PopulateContentPreview() {
	preview := stripmd.Strip(p.Content)
	if len(preview) > 80 {
		preview = preview[:80]
	}
	p.ContentPreview = preview
}
