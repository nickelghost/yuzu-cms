package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"text/template"

	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nickelghost/cms/db"
	"github.com/nickelghost/cms/handlers"
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
	// Get the database connection
	dbConn, err := db.Init(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL"),
	)
	defer dbConn.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrate(dbConn, "migrations")
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	theme := os.Getenv("APP_THEME")
	if theme == "" {
		theme = "default"
	}
	// Load view templates
	t := &Template{
		templates: template.Must(template.ParseGlob(
			fmt.Sprintf("themes/%s/views/*.html", theme),
		)),
	}
	// Init Echo
	e := echo.New()
	// Set our template renderer as the renderer for Echo
	e.Renderer = t
	// Add Echo middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/admin", "public")
	e.Static(
		fmt.Sprintf("/%s", theme),
		fmt.Sprintf("themes/%s/public", theme),
	)
	// Init the handlers object
	hs := handlers.Handlers{DB: dbConn}
	// Define routes
	e.GET("/", hs.Homepage)
	e.GET("/api/v1/posts", hs.APIPostsIndex)
	e.GET("/api/v1/posts/:id", hs.APIPostsGet)
	e.POST("/api/v1/posts", hs.APIPostsCreate)
	e.PUT("/api/v1/posts/:id", hs.APIPostsUpdate)
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
