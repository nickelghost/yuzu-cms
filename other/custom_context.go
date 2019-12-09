package other

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type CustomContext struct {
	echo.Context
	DB *gorm.DB
}
