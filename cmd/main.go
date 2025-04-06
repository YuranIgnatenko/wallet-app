package main

import (
	"log"
	v1 "wallet-app/api/v1"
	"wallet-app/config"
	"wallet-app/db"
)

func main() {
	config.LoadConfig("config.env")

	// for create table and fill test-data
	if conn, err := db.ConnectDatabase(); err == nil {
		defer conn.Close()
		db.MigrationReWriteTableWallets(conn)
	}

	router := v1.SetupRouter()
	log.Fatal(router.Run())
}
