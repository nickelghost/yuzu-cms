package template

import (
	"fmt"
	"html/template"
)

// GetRenderer returns the complete renderer for Echo to use
func GetRenderer(location string, theme string) *Template {
	renderer := &Template{
		templates: template.Must(template.ParseGlob(
			fmt.Sprintf(location, theme),
		)),
	}
	return renderer
}
