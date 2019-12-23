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

// Conn represents the connection to the database
var Conn *sql.DB

// Init initializes the connection to our database
func Init(
	host string,
	port string,
	user string,
	password string,
	name string,
	ssl string,
) error {
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
	Conn, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	err = Conn.Ping()
	if err != nil {
		return err
	}
	return nil
}

// Migrate runs all migrations against our database
func Migrate() error {
	driver, err := postgres.WithInstance(Conn, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
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
