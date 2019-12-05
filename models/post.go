package models

import (
	"time"

	"gopkg.in/russross/blackfriday.v2"
)

type Post struct {
	ID        uint `gorm:"primary_key"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p Post) GetHTMLContent() string {
	return string(blackfriday.Run([]byte(p.Content)))
}
