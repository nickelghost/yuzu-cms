package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Admin(c echo.Context) error {
	return c.Render(http.StatusOK, "admin.html", nil)
}
