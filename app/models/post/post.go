package post

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

// GetAll gets all posts that match a criteria
// drafts doesn't look at is_draft criteria if set to 0, finds only drafts if
// set to 1 and only non-drafts if set to -1
func GetAll(conn *sql.DB, drafts int) (*[]Post, error) {
	var rows *sql.Rows
	var err error
	if drafts == 0 {
		rows, err = conn.Query(`
            SELECT * FROM posts
            ORDER BY created_at DESC
        `)
	} else {
		var isDraft bool
		if drafts < 0 {
			isDraft = false
		} else if drafts > 0 {
			isDraft = true
		}
		rows, err = conn.Query(`
            SELECT * FROM posts
            WHERE is_draft = $1
            ORDER BY created_at DESC
        `, isDraft)
	}
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

func GetAllByIds(conn *sql.DB, ids []int) (*[]Post, error) {
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

func GetByID(conn *sql.DB, id int) (*Post, error) {
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

// Create inserts a new post to the database
func Create(
	conn *sql.DB,
	title string,
	content string,
	isDraft bool,
) (Post, error) {
	post := Post{}
	err := conn.QueryRow(`
        INSERT INTO posts (title, content, is_draft, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW())
        RETURNING *
        `,
		title,
		content,
		isDraft,
	).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.IsDraft,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return Post{}, err
	}
	return post, nil
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

func (p *Post) GetContentPreview(max int) string {
	preview := stripmd.Strip(p.Content)
	if len(preview) > max {
		preview = preview[:max]
	}
	return preview
}
