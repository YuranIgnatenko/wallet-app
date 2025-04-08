package main

import (
	"log"
	v1 "wallet-app/api/v1"
	"wallet-app/config"
	"wallet-app/db"
)

func main() {
	if err := config.LoadConfig("config.env"); err != nil {
		panic(err)
	}

	// for create table and fill test-data
	if conn, err := db.ConnectDatabase(); err == nil {
		defer conn.Close()
		err = db.MigrationReWriteTableWallets(conn)
		if err != nil{
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}

	router := v1.SetupRouter()
	log.Fatal(router.Run(config.AppConfig.SERVER_URL))
}
