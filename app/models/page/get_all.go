package page

import (
	"database/sql"

	"github.com/nickelghost/yuzu-cms/app/models/post"
)

// GetAll returns all pages from the database
func GetAll(conn *sql.DB, withPosts bool) (*[]Page, error) {
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
		posts, err := post.GetAllByIds(conn, postIds)
		if err != nil {
			return nil, err
		}
		pagesWithPosts := []Page{}
		for _, page := range pages {
			pageWithPost := page
			for _, post := range *posts {
				if page.PostID == post.ID {
					pageWithPost.Post = post
				}
			}
			pagesWithPosts = append(pagesWithPosts, pageWithPost)
		}
		pages = pagesWithPosts
	}
	return &pages, nil
}
