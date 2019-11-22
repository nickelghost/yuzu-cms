package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nickelghost/cms/controllers"
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

func main() {
	// Load variables from .env
	err := godotenv.Load()
	// Check for errors and ignore if file not found
	if err != nil && err.Error() != "open .env: no such file or directory" {
		log.Fatal(err)
	}
	// Load view templates
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	// Init Echo
	e := echo.New()
	// Set our template renderer as the renderer for Echo
	e.Renderer = t
	// Add Echo middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("public"))
	// Define routes
	e.GET("/", controllers.Homepage)
	// Define where to serve
	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Println(
			"No APP_PORT env variable found. Using default port 3000...",
		)
		port = "3000"
	}
	connStr := fmt.Sprintf(":%s", port)
	// Start the server
	e.Logger.Fatal(e.Start(connStr))
}
