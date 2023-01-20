package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadDB_PostgresSQL() (string, error) {
	err := godotenv.Load("environment_variable/DB/.env")
	if err != nil {
		return "", err
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}

	database := os.Getenv("DB_DATABASE")
	if database == "" {
		database = "db-golang"
	}

	username := os.Getenv("DB_USERNAME")
	if username == "" {
		username = "root"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "root"
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	CONNECTION := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4,utf8&readTimeout=30s&writeTimeout=30s&parseTime=true", username, password, host, port, database)
	return CONNECTION, nil
}
