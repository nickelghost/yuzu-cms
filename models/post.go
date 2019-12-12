package models

import (
	"time"

	"gopkg.in/russross/blackfriday.v2"
)

// Post contains articles and pages used by the CMS
type Post struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetHTMLContent returns the post's markdown content as HTML
func (p Post) GetHTMLContent() string {
	return string(blackfriday.Run([]byte(p.Content)))
}
