package db

import (
	"fmt"
	"os"
	"testing"

	"github.com/nickelghost/yuzu-cms/boot"

	"github.com/golang-migrate/migrate/v4"

	"github.com/joho/godotenv"
)

func TestInit(t *testing.T) {
	godotenv.Load("../.env")
	conn, err := Init(boot.GetPostgresConnString())
	if err != nil {
		t.Errorf("Connection to the DB failed:\n%s", err.Error())
	}
	if err := conn.Ping(); err != nil {
		t.Errorf("Could not ping the DB:\n%s", err)
	}
}

func TestInitMock(t *testing.T) {
	godotenv.Load("../.env")
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		"nonexistingdatabase",
		os.Getenv("DB_SSL"),
	)
	fmt.Println(connStr)
	_, err := Init(connStr)
	if err == nil {
		t.Error("Didn't return an error on wrong database")
	}
}

func TestMigrate(t *testing.T) {
	// TODO: Separate cases of when the DB is migrated and when isn't migrated
	godotenv.Load("../.env")
	conn, _ := Init(boot.GetPostgresConnString())
	err := Migrate(conn, "../migrations")
	if err != nil && err != migrate.ErrNoChange {
		t.Errorf("Couldn't run migrations:\n%s", err.Error())
	}
}
