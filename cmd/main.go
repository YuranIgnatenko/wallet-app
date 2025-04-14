package main

import (
	"log"
	"wallet-app/config"
	"wallet-app/db"
	v1 "wallet-app/internal/api/v1"
	"wallet-app/internal/logger"

	"github.com/jackc/pgx/v4/pgxpool"
)

var poolConnection *pgxpool.Pool

func main() {
	if err := config.LoadConfig("config.env"); err != nil {
		panic(err)
	}

	logger.InitLogger()
	logger.Log.Info("Starting app")

	poolConnection, err := db.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer poolConnection.Close()

	err = db.MigrationReWriteTableWallets(poolConnection)
	if err != nil {
		log.Fatal(err)
	}

	router := v1.SetupRouter(poolConnection)
	log.Fatal(router.Run(config.AppConfig.SERVER_URL))
}
