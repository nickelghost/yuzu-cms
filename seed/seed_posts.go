package seed

import (
	"database/sql"

	postModel "github.com/nickelghost/yuzu-cms/app/models/post"
)

var postBlueprints = []postModel.Post{
	postModel.Post{
		Title: "My vacation",
		Slug:  "my-vacation",
		Content: `
I went to Maledives, and this is my story.

## The flight

The flight was quite nice. I few with Lufthansa.

## Hotel

I booked a five star hotel on the largest island.
        `,
		IsDraft: false,
	},
	postModel.Post{
		Title: "My draft",
		Slug:  "some-custom-slug",
		Content: `
This is just a test for the draft functionality.

## One heading

Some more content here.

## Second heading

And even more content.
        `,
		IsDraft: true,
	},
	postModel.Post{
		Title: "Some more stuff",
		Slug:  "some-more-stuff",
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

func seedPosts(conn *sql.DB) (*[]postModel.Post, error) {
	posts := []postModel.Post{}
	for _, postBlueprint := range postBlueprints {
		post, err := postModel.Create(
			conn,
			postBlueprint.Title,
			postBlueprint.Content,
			postBlueprint.Slug,
			postBlueprint.IsDraft,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return &posts, nil
}
