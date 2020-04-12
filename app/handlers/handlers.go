package handlers

import (
	"database/sql"

	"github.com/nickelghost/yuzu-cms/app/config"
)

// Handlers provides common data for our handlers
type Handlers struct {
	DB     *sql.DB
	Config config.Config
}
