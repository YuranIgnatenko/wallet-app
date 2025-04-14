package main

import (
	"log"
	v1 "wallet-app/api/v1"
	"wallet-app/config"
	"wallet-app/db"

	"github.com/jackc/pgx/v4/pgxpool"
)

var poolConnection *pgxpool.Pool

func main() {
	if err := config.LoadConfig("config.env"); err != nil {
		panic(err)
	}

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
