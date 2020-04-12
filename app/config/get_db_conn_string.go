package config

import "fmt"

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
