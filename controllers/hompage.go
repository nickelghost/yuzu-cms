package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Homepage(c echo.Context) error {
	return c.String(http.StatusOK, "Homepage")
}
