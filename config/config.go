package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ServerAddr string `env:"SERVER_ADDR"`
	ServerPort string `env:"SERVER_PORT"`
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addr := os.Getenv("SERVER_ADDR")
	port := os.Getenv("SERVER_PORT")

	cfg := Config{
		ServerAddr: addr,
		ServerPort: port,
	}

	return &cfg
}
