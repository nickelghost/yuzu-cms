package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestInit(t *testing.T) {
	godotenv.Load("../.env")
	err := Init(
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
}

func TestInitMock(t *testing.T) {
	godotenv.Load("../.env")
	err := Init(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		"nonexistingdatabase",
		os.Getenv("DB_SSL"),
	)
	if err == nil {
		t.Errorf("Didn't return an error on wrong password")
	}
}
