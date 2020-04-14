package seed

import (
	"database/sql"

	pageModel "github.com/nickelghost/yuzu-cms/app/models/page"
)

func Seed(conn *sql.DB) error {
	posts, err := seedPosts(conn)
	if err != nil {
		return err
	}
	_, err = pageModel.Create(conn, (*posts)[0].ID, 1, "about-my-vacation", true)
	if err != nil {
		return err
	}
	return nil
}
