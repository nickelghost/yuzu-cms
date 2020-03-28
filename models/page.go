package models

import (
	"database/sql"
	"time"
)

type Page struct {
	ID           int
	PostID       int
	Index        int
	Slug         string
	InNavigation bool
	Heading      *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Post         Post
}

func (Page) GetAll(conn *sql.DB, withPosts bool) (*[]Page, error) {
	rows, err := conn.Query("SELECT * FROM pages ORDER BY index DESC")
	if err != nil {
		return nil, err
	}
	pages := []Page{}
	postIds := []int{}
	for rows.Next() {
		page := Page{}
		err := rows.Scan(
			&page.ID,
			&page.PostID,
			&page.Index,
			&page.Slug,
			&page.InNavigation,
			&page.Heading,
			&page.CreatedAt,
			&page.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)
		postIds = append(postIds, page.PostID)
	}
	if withPosts {
		posts, err := (Post{}).GetAllByIds(conn, postIds)
		if err != nil {
			return nil, err
		}
		pagesWithPosts := []Page{}
		for _, page := range pages {
			pageWithPost := page
			for _, post := range *posts {
				if page.PostID == int(post.ID) {
					pageWithPost.Post = post
				}
			}
			pagesWithPosts = append(pagesWithPosts, pageWithPost)
		}
		pages = pagesWithPosts
	}
	return &pages, nil
}

func (Page) GetById(conn *sql.DB, id int) (*Page, error) {
	page := Page{}
	err := conn.QueryRow("SELECT * FROM pages WHERE id = $1", id).Scan(
		&page.ID,
		&page.PostID,
		&page.Index,
		&page.Slug,
		&page.InNavigation,
		&page.Heading,
		&page.CreatedAt,
		&page.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &page, nil
}

func (Page) GetBySlug(conn *sql.DB, slug string) (*Page, error) {
	page := Page{}
	err := conn.QueryRow("SELECT * FROM pages WHERE slug = $1", slug).Scan(
		&page.ID,
		&page.PostID,
		&page.Index,
		&page.Slug,
		&page.InNavigation,
		&page.Heading,
		&page.CreatedAt,
		&page.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &page, nil
}

func (Page) Create(
	conn *sql.DB,
	postID int,
	index int,
	slug string,
	inNavigation bool,
	heading *string,
) (*Page, error) {
	page := Page{}
	err := conn.QueryRow(`
        INSERT INTO
        pages (post_id, index, slug, in_navigation, heading, created_at, updated_at)
        VALUES
        ($1, $2, $3, $4, $5, NOW(), NOW())
        RETURNING *
        `,
		postID,
		index,
		slug,
		inNavigation,
		heading,
	).Scan(
		&page.ID,
		&page.PostID,
		&page.Index,
		&page.Slug,
		&page.InNavigation,
		&page.Heading,
		&page.CreatedAt,
		&page.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &page, nil
}

func (p *Page) Update(
	conn *sql.DB,
	postID int,
	index int,
	slug string,
	inNavigation bool,
	heading *string,
) error {
	if p.Index != index {
		_, err := conn.Exec(`
            UPDATE pages
            SET index = index + $2
            WHERE index = $1
        `, index, p.Index-index)
		if err != nil {
			return err
		}
	}
	err := conn.QueryRow(`
        UPDATE pages
        SET
            post_id = $2,
            index = $3,
            slug = $4,
            in_navigation = $5,
            heading = $6,
            updated_at = NOW()
        WHERE id = $1
        RETURNING
            id, post_id, index, slug,
            in_navigation, heading, created_at, updated_at
        `,
		p.ID, postID, index, slug, inNavigation, heading,
	).Scan(
		&p.ID,
		&p.PostID,
		&p.Index,
		&p.Slug,
		&p.InNavigation,
		&p.Heading,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p Page) Delete(conn *sql.DB) error {
	_, err := conn.Exec("DELETE FROM pages WHERE id = $1", p.ID)
	if err != nil {
		return err
	}
	return nil
}
