package boot

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadEnv gets our .env file or throws a fatal
func LoadEnv() {
	// Load variables from .env
	err := godotenv.Load()
	// Check for errors and ignore if file not found
	if err != nil && err.Error() != "open .env: no such file or directory" {
		log.Fatal(err)
	}
}

// GetRenderer returns the complete renderer for Echo to use
func GetRenderer(location string, theme string) *Template {
	renderer := &Template{
		templates: template.Must(template.ParseGlob(
			fmt.Sprintf(location, theme),
		)),
	}
	return renderer
}

// GetTheme gets the theme from the environment or sets default
func GetTheme() string {
	theme := os.Getenv("APP_THEME")
	if theme == "" {
		theme = "default"
	}
	return theme
}

// GetPort gets port from the environment or falls back to 3000
func GetPort() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Println(
			"No APP_PORT env variable found. Using default port 3000...",
		)
		port = "3000"
	}
	return port
}

// GetPostgresConnString creates a connection string for our database from
// environment variables
func GetPostgresConnString() string {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL"),
	)
	return connStr
}

// GetSQL reads SQL files in a given location and returns a map with .sql file
// names and their content
func GetSQL(root string) map[string]string {
	sqlFiles := make(map[string]string)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".sql" {
			sql, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			sqlFiles[info.Name()] = string(sql)
			return nil
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Reading SQL files failed:\n%s", err)
	}
	return sqlFiles
}
