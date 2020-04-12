package seed

import (
	"database/sql"

	"github.com/nickelghost/yuzu-cms/app/models"
)

var posts = []models.Post{
	models.Post{
		Title: "My vacation",
		Content: `
I went to Maledives, and this is my story.

## The flight

The flight was quite nice. I few with Lufthansa.

## Hotel

I booked a five star hotel on the largest island.
        `,
		IsDraft: false,
	},
	models.Post{
		Title: "My draft",
		Content: `
This is just a test for the draft functionality.

## One heading

Some more content here.

## Second heading

And even more content.
        `,
		IsDraft: true,
	},
	models.Post{
		Title: "Some more stuff",
		Content: `
This is a preview of some stuff

# Big heading

Some content...

###### Small heading ;)

And even more content!
        `,
		IsDraft: false,
	},
}

func seedPosts(conn *sql.DB) error {
	for _, post := range posts {
		_, err := conn.Exec(
			`INSERT INTO
                posts (title, content, is_draft, created_at, updated_at)
            VALUES
                ($1, $2, $3, NOW(), NOW())
            `,
			post.Title,
			post.Content,
			post.IsDraft,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
