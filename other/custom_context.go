package other

import (
	"database/sql"

	"github.com/labstack/echo"
)

type CustomContext struct {
	echo.Context
	DB *sql.DB
}
