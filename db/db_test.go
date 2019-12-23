package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kylelemons/godebug/diff"
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

func TestGetSQL(t *testing.T) {
	filePath := "./queries/test_query.sql"
	sqlString := "SELECT * FROM tests;\n"
	f, _ := os.Create(filePath)
	defer f.Close()
	defer os.Remove(filePath)
	f.Write([]byte(sqlString))
	sql, err := GetSQL("test_query")
	if err != nil {
		t.Errorf("GetSQL threw an error:\n%s", err.Error())
	}
	if sql != sqlString {
		t.Errorf(
			"Did not get the correct SQL:\n%v",
			diff.Diff(sql, sqlString),
		)
	}
}
