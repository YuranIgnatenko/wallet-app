package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST      string
	DB_PORT      string
	DB_USER      string
	DB_PASS      string
	DB_NAME      string
	DATABASE_URL string
}

var AppConfig *Config

// Initialization variable AppConfig (type *Config)
// for access from all code app
func LoadConfig(namefile string) {
	err := godotenv.Load(namefile)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")

	db_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		db_host, db_port, db_user, db_pass, db_name)

	AppConfig = &Config{
		DB_HOST:      db_host,
		DB_PORT:      db_port,
		DB_USER:      db_user,
		DB_PASS:      db_pass,
		DB_NAME:      db_name,
		DATABASE_URL: db_url,
	}
}
