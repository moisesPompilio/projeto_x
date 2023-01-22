package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func LoadDB_PostgreSQL() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	database := os.Getenv("DB_DATABASE")
	if database == "" {
		database = "db-golang"
	}

	username := os.Getenv("DB_USERNAME")
	if username == "" {
		username = "postgres"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}

	CONNECTION := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", username, password, host, port, database, sslmode)
	return CONNECTION, nil
}

func GetReaderPostgreSQL() *sqlx.DB {
	DB_CONNECTION, err := LoadDB_PostgreSQL()
	if err != nil {
		panic(err)
	}

	reader := sqlx.MustConnect("pg", DB_CONNECTION)

	return reader
}

func GetWriterPostgreSQL() *sqlx.DB {
	DB_CONNECTION, err := LoadDB_PostgreSQL()
	if err != nil {
		panic(err)
	}
	writer := sqlx.MustConnect("pg", DB_CONNECTION)

	return writer
}
