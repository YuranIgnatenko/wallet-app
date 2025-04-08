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
	SERVER_HOST  string
	SERVER_PORT  string
	SERVER_URL   string
}

var AppConfig *Config

// Initialization variable AppConfig (type *Config)
// for access from all code app
func LoadConfig(namefile string) error {
	if _, err := os.Stat(namefile); os.IsNotExist(err) {
		return err
	}

	err := godotenv.Load(namefile)
	if err != nil {
		log.Fatal(err)
		return err
	}

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db_user, db_pass, db_host, db_port, db_name)
	server_host := os.Getenv("SERVER_HOST")
	server_port := os.Getenv("SERVER_PORT")
	server_url := fmt.Sprintf("%s:%s", server_host, server_port)

	AppConfig = &Config{
		DB_HOST:      db_host,
		DB_PORT:      db_port,
		DB_USER:      db_user,
		DB_PASS:      db_pass,
		DB_NAME:      db_name,
		DATABASE_URL: db_url,
		SERVER_HOST:  server_host,
		SERVER_PORT:  server_port,
		SERVER_URL:   server_url,
	}
	return nil
}
