package page

import "database/sql"

// Delete removes a page from the DB
func (p Page) Delete(conn *sql.DB) error {
	_, err := conn.Exec("DELETE FROM pages WHERE id = $1", p.ID)
	if err != nil {
		return err
	}
	return nil
}
