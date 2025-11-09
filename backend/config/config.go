package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

// DBURL builds the database connection string
func DBURL() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Config("DB_HOST"),
		Config("DB_PORT"),
		Config("DB_USER"),
		Config("DB_PASSWORD"),
		Config("DB_NAME"),
	)
}
