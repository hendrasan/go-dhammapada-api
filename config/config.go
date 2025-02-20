package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Port       string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		Port:       os.Getenv("APP_PORT"),
	}

	if config.Port == "" {
		config.Port = "8080" // Default port if not specified
	}

	// Debug config values
	log.Printf("Loaded config - Host: %s, Port: %s, User: %s, DBName: %s, Server Port: %s",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBName,
		config.Port,
	)

	return config, nil
}
