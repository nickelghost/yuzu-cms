package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nickelghost/yuzu-cms/app/config"
	"github.com/nickelghost/yuzu-cms/app/db"
	"github.com/nickelghost/yuzu-cms/app/handlers"
	"github.com/nickelghost/yuzu-cms/app/template"
	"github.com/nickelghost/yuzu-cms/seed"
)

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func forwardWebpack(e *echo.Echo, targetURL string) {
	g := e.Group("/admin")
	webpackURL, err := url.Parse(targetURL)
	if err != nil {
		log.Fatal(err)
	}
	g.Use(middleware.Proxy(middleware.NewRandomBalancer(
		[]*middleware.ProxyTarget{{URL: webpackURL}},
	)))
}

func main() {
	config, err := config.GetFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	// Get the database connection
	dbConn, err := db.Init(config.GetDBConnString())
	defer dbConn.Close()
	if err != nil {
		log.Fatal(err)
	}
	if contains(os.Args, "-migrate") {
		err = db.Migrate(dbConn, "migrations")
		if err != nil {
			log.Fatal(err)
		}
	}
	if contains(os.Args, "-seed") {
		err := seed.Seed(dbConn)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Load view templates
	renderer := template.GetRenderer("themes/%s/views/*.html", config.AppTheme)
	// Init Echo
	e := echo.New()
	// Set our template renderer as the renderer for Echo
	e.Renderer = renderer
	// Add Echo middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/admin", "public")
	e.Static(
		fmt.Sprintf("/%s", config.AppTheme),
		fmt.Sprintf("themes/%s/public", config.AppTheme),
	)
	// Init the handlers object
	hs := handlers.Handlers{
		DB:     dbConn,
		Config: config,
	}
	// Public routes
	e.GET("/", hs.Homepage)
	e.GET("/:slug", hs.Page)
	e.GET("/posts/:slug", hs.Post)
	e.POST("/api/v1/login", hs.APILogin)
	// Auth-only v1 API routes
	v1auth := e.Group("/api/v1")
	v1auth.Use(middleware.JWT(config.AppSecret))
	v1auth.GET("/posts", hs.APIPostsIndex)
	v1auth.GET("/posts/titles", hs.APIPostsTitles)
	v1auth.GET("/posts/:id", hs.APIPostsGet)
	v1auth.POST("/posts", hs.APIPostsCreate)
	v1auth.PUT("/posts/:id", hs.APIPostsUpdate)
	v1auth.DELETE("/posts/:id", hs.APIPostsDelete)
	v1auth.GET("/pages", hs.APIPagesIndex)
	v1auth.POST("/pages", hs.APIPagesCreate)
	v1auth.PUT("/pages/:id", hs.APIPagesUpdate)
	v1auth.DELETE("/pages/:id", hs.APIPagesDelete)
	// This forwards webpack's web server for development
	if config.AppForwardWebpack {
		forwardWebpack(e, "http://localhost:3001")
	}
	connStr := fmt.Sprintf(":%d", config.AppPort)
	// Start the server
	e.Logger.Fatal(e.Start(connStr))
}
