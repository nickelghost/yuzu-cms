package other

import (
	"database/sql"

	"github.com/labstack/echo"
)

// CustomContext provides structure for an additional context for echo
// that contains data needed across requests
type CustomContext struct {
	echo.Context
	DB *sql.DB
}
