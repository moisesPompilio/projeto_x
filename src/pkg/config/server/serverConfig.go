package server

import (
	"os"

	"github.com/joho/godotenv"
)

func Load_Server() (string, error) {
	err := godotenv.Load("environment_variable/server/.env")
	if err != nil {
		return "", err
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port, nil
}
