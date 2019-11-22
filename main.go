package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nickelghost/cms/controllers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", controllers.Homepage)

	e.Logger.Fatal(e.Start(":3000"))
}
