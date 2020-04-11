package models

import (
	"database/sql"
	"fmt"
	"html/template"
	"strings"
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

func (Post) GetAll(conn *sql.DB) (*[]Post, error) {
	rows, err := conn.Query("SELECT * FROM posts ORDER BY created_at DESC")
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

func (Post) GetAllByIds(conn *sql.DB, ids []int) (*[]Post, error) {
	idsString := strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")
	param := "{" + idsString + "}"
	rows, err := conn.Query("SELECT * FROM posts WHERE id = ANY($1::int[])", param)
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

func (Post) GetByID(conn *sql.DB, id int) (*Post, error) {
	post := Post{}
	err := conn.QueryRow(
		`SELECT * FROM posts WHERE id = $1`,
		id,
	).Scan(
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
	return &post, nil
}

// GetHTMLContent returns the post's markdown content as HTML
func (p *Post) GetHTMLContent() template.HTML {
	str := string(blackfriday.Run([]byte(p.Content)))
	return template.HTML(str)
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

func (p Post) GetContentPreview(max int) string {
	preview := stripmd.Strip(p.Content)
	if len(preview) > max {
		preview = preview[:max]
	}
	return preview
}
