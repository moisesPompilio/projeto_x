package server

import (
	"os"

	"github.com/joho/godotenv"
)

func Load_Server() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "3005"
	}
	return port, nil
}
