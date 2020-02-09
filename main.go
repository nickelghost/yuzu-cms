package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nickelghost/yuzu-cms/boot"

	"github.com/golang-migrate/migrate/v4"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nickelghost/yuzu-cms/db"
	"github.com/nickelghost/yuzu-cms/handlers"
)

func main() {
	boot.LoadEnv()
	// Get the database connection
	dbConn, err := db.Init(boot.GetPostgresConnString())
	defer dbConn.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrate(dbConn, "migrations")
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	theme := boot.GetTheme()
	// Load view templates
	renderer := boot.GetRenderer("themes/%s/views/*.html", theme)
	// Init Echo
	e := echo.New()
	// Set our template renderer as the renderer for Echo
	e.Renderer = renderer
	// Add Echo middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/admin", "public")
	e.Static(
		fmt.Sprintf("/%s", theme),
		fmt.Sprintf("themes/%s/public", theme),
	)
	// Init the handlers object
	hs := handlers.Handlers{DB: dbConn, SQL: boot.GetSQL("queries/")}
	// Public routes
	e.GET("/", hs.Homepage)
	e.POST("/api/v1/login", hs.APILogin)
	// Auth-only v1 API routes
	v1auth := e.Group("/api/v1")
	v1auth.Use(middleware.JWT([]byte(os.Getenv("APP_SECRET"))))
	v1auth.GET("/posts", hs.APIPostsIndex)
	v1auth.GET("/posts/:id", hs.APIPostsGet)
	v1auth.POST("/posts", hs.APIPostsCreate)
	v1auth.PUT("/posts/:id", hs.APIPostsUpdate)
	// This forwards webpack's web server for development
	if os.Getenv("APP_WEBPACK_FORWARD") == "true" {
		boot.ForwardWebpack(e, "http://localhost:3001")
	}
	// Determine where to serve
	port := boot.GetPort()
	connStr := fmt.Sprintf(":%s", port)
	// Start the server
	e.Logger.Fatal(e.Start(connStr))
}
