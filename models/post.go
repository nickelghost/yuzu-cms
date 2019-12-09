package models

import (
	"time"

	"gopkg.in/russross/blackfriday.v2"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p Post) GetHTMLContent() string {
	return string(blackfriday.Run([]byte(p.Content)))
}
