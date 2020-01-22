package db

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// Driver for migrations from sql files
	_ "github.com/golang-migrate/migrate/v4/source/file"

	// PostgreSQL driver
	_ "github.com/lib/pq"
)

// Init initializes the connection to our database and tests it
func Init(
	host string,
	port string,
	user string,
	password string,
	name string,
	ssl string,
) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		user,
		password,
		name,
		ssl,
	)
	var err error
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Migrate runs all migrations against our database
func Migrate(conn *sql.DB, path string) error {
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+path,
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil {
		return err
	}
	return nil
}
