package models

import (
	"database/sql"
	"strconv"
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

func (_ Post) GetAllByIds(conn *sql.DB, ids []int) (*[]Post, error) {
	idsString := ""
	for i, id := range ids {
		idsString = idsString + strconv.Itoa(id)
		if i != len(ids)-1 {
			idsString = idsString + ", "
		}
	}
	// TODO: Implement a solution that doesn't involve string concat
	rows, err := conn.Query("SELECT * FROM posts WHERE id IN (" + idsString + ")")
	if err != nil {
		return nil, err
	}
	posts := []Post{}
	for rows.Next() {
		post := Post{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.IsDraft,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return &posts, nil
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
