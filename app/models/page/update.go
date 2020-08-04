package page

import "database/sql"

// Update updates the page in the DB
func (p *Page) Update(
	conn *sql.DB,
	postID int,
	index int,
	slug string,
	inNavigation bool,
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
            updated_at = NOW()
        WHERE id = $1
        RETURNING
            id, post_id, index, slug,
            in_navigation, created_at, updated_at
        `,
		p.ID, postID, index, slug, inNavigation,
	).Scan(
		&p.ID,
		&p.PostID,
		&p.Index,
		&p.Slug,
		&p.InNavigation,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
