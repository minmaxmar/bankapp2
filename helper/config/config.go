package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel    string
	DatabaseURL string
	ServerPort  int `envconfig:"serverport" required:"true" default:"8080"`
}

func LoadConfig() *Config {

	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=UTC",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	port, err := strconv.Atoi(os.Getenv("SERVERPORT"))
	if err != nil {
		log.Fatalf("Error converting SERVERPORT to int: %v", err)
	}

	config := &Config{
		LogLevel:    os.Getenv("LOG_LEVEL"),
		DatabaseURL: dsn,
		ServerPort:  port,
	}

	if config.LogLevel == "" {
		config.LogLevel = "info"
	}

	return config
}
