package boot

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

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
