package database

import (
	"os"
	"time"
)

// config represents the config's attributes
type config struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string

	ctxTimeout time.Duration
}

// newConfigMongoDB will create a new instance of config for MongoDB
func newConfigMongoDB() *config {
	return &config{
		host:       os.Getenv("MONGODB_HOST"),
		database:   os.Getenv("MONGODB_DATABASE"),
		password:   os.Getenv("MONGODB_ROOT_PASSWORD"),
		user:       os.Getenv("MONGODB_ROOT_USER"),
		ctxTimeout: 60 * time.Second,
	}
}

// newConfigPostgres will create a new instance of config for PostgreSQL
func newConfigPostgres() *config {
	return &config{
		host:     os.Getenv("POSTGRES_HOST"),
		database: os.Getenv("POSTGRES_DATABASE"),
		port:     os.Getenv("POSTGRES_PORT"),
		driver:   os.Getenv("POSTGRES_DRIVER"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
	}
}
