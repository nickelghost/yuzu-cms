package template

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo"
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

// Template is a struct required for rendering templating views
type Template struct {
	templates *template.Template
}

// Render actually handles the rendering of the templates
func (t *Template) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	return t.templates.ExecuteTemplate(w, name, data)
}