package models

import (
	"testing"

	"github.com/kylelemons/godebug/diff"
)

func TestGetHTMLContent(t *testing.T) {
	post := Post{Content: "# Title\n\nSome content here"}
	html := post.GetHTMLContent()
	expected := "<h1>Title</h1>\n\n<p>Some content here</p>\n"
	if html != expected {
		t.Fatalf(
			"Did not return correct HTML converted string\n%v",
			diff.Diff(html, expected),
		)
	}
}
