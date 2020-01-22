package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestInit(t *testing.T) {
	godotenv.Load("../.env")
	conn, err := Init(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL"),
	)
	if err != nil {
		t.Errorf("Connection to the DB failed:\n%s", err.Error())
	}
	if err := conn.Ping(); err != nil {
		t.Errorf("Could not ping the DB:\n%s", err)
	}
}

func TestInitMock(t *testing.T) {
	godotenv.Load("../.env")
	_, err := Init(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		"nonexistingdatabase",
		os.Getenv("DB_SSL"),
	)
	if err == nil {
		t.Errorf("Didn't return an error on wrong database\n%s", err)
	}
}

func TestMigrate(t *testing.T) {
	godotenv.Load("../.env")
	conn, _ := Init(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL"),
	)
	err := Migrate(conn, "../migrations")
	if err != nil && err.Error() != "no change" {
		t.Errorf("Couldn't run migrations:\n%s", err)
	}
}
