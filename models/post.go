package models

import (
	"time"
)

type Post struct {
	ID             uint `gorm:"primary_key"`
	Title          string
	Content        string
	ContentPreview string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
