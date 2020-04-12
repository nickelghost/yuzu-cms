package seed

import (
	"database/sql"
)

func Seed(conn *sql.DB) error {
	_, err := seedPosts(conn)
	if err != nil {
		return err
	}
	return nil
}
