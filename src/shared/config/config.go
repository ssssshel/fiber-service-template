package config

import (
	"os"

	"github.com/ssssshel/fiber-service-template/src/shared/interfaces"
)

// DBConnectionConfig

func DBConnectionConfig() *interfaces.IDBConnConfig {
	values := &interfaces.IDBConnConfig{
		DB_NAME:     os.Getenv("POSTGRESQL_DB"),
		DB_PORT:     os.Getenv("POSTGRESQL_DB_PORT"),
		DB_HOST:     os.Getenv("POSTGRESQL_DB_HOST"),
		DB_PASSWORD: os.Getenv("POSTGRESQL_DB_PASSWORD"),
		DB_SCHEMA:   os.Getenv("POSTGRESQL_DB_SCHEMA"),
		DB_USER:     os.Getenv("POSTGRESQL_DB_USER"),
	}

	return values
}

// ATKey

func ATKey() string {
	key := os.Getenv("ACCESS_TOKEN_PRIVATE_KEY")
	return key
}

// APIVersion

type APIVersion uint

const (
	V1 APIVersion = 1
)

// Port

func Port() string {
	port := os.Getenv("PORT")
	return port
}
