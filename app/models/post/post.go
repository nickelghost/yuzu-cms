package post

import (
	"html/template"
	"time"

	stripmd "github.com/writeas/go-strip-markdown"
	"gopkg.in/russross/blackfriday.v2"
)

// Post contains articles and pages used by the CMS
type Post struct {
	ID             int
	Title          string
	Content        string
	ContentPreview string
	Slug           string
	IsDraft        bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// GetHTMLContent returns the post's markdown content as HTML
func (p *Post) GetHTMLContent() template.HTML {
	str := string(blackfriday.Run([]byte(p.Content)))
	return template.HTML(str)
}

// GetContentPreview strips markdown from the content string, trims and returns it
func (p *Post) GetContentPreview(max int) string {
	preview := stripmd.Strip(p.Content)
	if len(preview) > max {
		preview = preview[:max]
	}
	return preview
}
