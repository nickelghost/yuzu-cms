package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config represents the configuration of the CMS instance
type Config struct {
	AppPort           int
	AppSecret         []byte
	AppTheme          string
	AppForwardWebpack bool
	AppUserName       string
	AppUserPassword   []byte
	DBHost            string
	DBPort            int
	DBUser            string
	DBPassword        string
	DBName            string
	DBRequireSSL      bool
}

// GetDBConnString returns a connection string for our database
func (c *Config) GetDBConnString() string {
	var dbRequireSSL string
	if c.DBRequireSSL {
		dbRequireSSL = "require"
	} else {
		dbRequireSSL = "disable"
	}
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost,
		c.DBPort,
		c.DBUser,
		c.DBPassword,
		c.DBName,
		dbRequireSSL,
	)
	return connStr
}

// GetFromEnv reads environment varialbes to get configuration for the app
func GetFromEnv() (Config, error) {
	// Load variables from .env
	err := godotenv.Load()
	// Check for errors and ignore if file not found
	if err != nil && err.Error() != "open .env: no such file or directory" {
		log.Fatal(err)
	}
	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		return Config{}, errors.New("Invalid app port specified")
	}
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return Config{}, errors.New("Invalid db port specified")
	}
	appForwardWebpack, err := strconv.ParseBool(os.Getenv("APP_FORWARD_WEBPACK"))
	if err != nil {
		return Config{}, errors.New("Invalid app forward webpack option specified")
	}
	dbSSL, err := strconv.ParseBool(os.Getenv("DB_REQUIRE_SSL"))
	if err != nil {
		return Config{}, errors.New("Invalid DB SSL option specified")
	}
	config := Config{
		AppPort:           appPort,
		AppSecret:         []byte(os.Getenv("APP_SECRET")),
		AppTheme:          os.Getenv("APP_THEME"),
		AppForwardWebpack: appForwardWebpack,
		AppUserName:       os.Getenv("APP_USER_NAME"),
		AppUserPassword:   []byte(os.Getenv("APP_USER_PASS")),
		DBHost:            os.Getenv("DB_HOST"),
		DBPort:            dbPort,
		DBUser:            os.Getenv("DB_USER"),
		DBPassword:        os.Getenv("DB_PASS"),
		DBName:            os.Getenv("DB_NAME"),
		DBRequireSSL:      dbSSL,
	}
	return config, nil
}
