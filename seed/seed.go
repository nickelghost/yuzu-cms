package seed

import (
	"database/sql"
)

func Seed(conn *sql.DB) error {
	err := seedPosts(conn)
	if err != nil {
		return err
	}
	return nil
}
