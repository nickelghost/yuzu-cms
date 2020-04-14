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
	Slug           string
	IsDraft        bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// GetAll gets all posts that match a criteria
// isDraft doesn't look at the is_draft criteria if set to -1, only finds
// non-drafts if set to 0 and only rafts if set to 1
func GetAll(conn *sql.DB, isDraft int) (*[]Post, error) {
	var rows *sql.Rows
	var err error
	if isDraft < 0 {
		rows, err = conn.Query(`
            SELECT * FROM posts
            ORDER BY created_at DESC
        `)
	} else {
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
			&post.Slug,
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
			&post.Slug,
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
		&post.Slug,
		&post.IsDraft,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// GetBySlug fetches the post from the database by its URL slug
func GetBySlug(conn *sql.DB, slug string) (*Post, error) {
	post := Post{}
	err := conn.QueryRow(
		`SELECT * FROM posts WHERE slug = $1`,
		slug,
	).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Slug,
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
	slug string,
	isDraft bool,
) (Post, error) {
	post := Post{}
	err := conn.QueryRow(`
        INSERT INTO posts
            (title, content, slug, is_draft, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())
        RETURNING *
        `,
		title,
		content,
		slug,
		isDraft,
	).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Slug,
		&post.IsDraft,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

// Update the DB record for a post
func Update(
	conn *sql.DB,
	id int,
	title string,
	content string,
	slug string,
	isDraft bool,
) (Post, error) {
	post := Post{}
	err := conn.QueryRow(`
        UPDATE posts
        SET title = $2, content = $3, slug = $4, is_draft = $5, updated_at = NOW()
        WHERE id = $1
        RETURNING *
        `,
		id,
		title,
		content,
		slug,
		isDraft,
	).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Slug,
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

// GetContentPreview strips markdown from the content string, trims and returns it
func (p *Post) GetContentPreview(max int) string {
	preview := stripmd.Strip(p.Content)
	if len(preview) > max {
		preview = preview[:max]
	}
	return preview
}
