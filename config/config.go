package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URL string
}

func NewConfig(namefile string) *Config {
	err := godotenv.Load(namefile)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	database_url := os.Getenv("DATABASE_URL")

	return &Config{
		DATABASE_URL: database_url,
	}
}
