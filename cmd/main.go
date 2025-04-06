package main

import (
	"log"
	v1 "wallet-app/api/v1"
	"wallet-app/config"
	"wallet-app/db"
)

func main() {
	cfg := config.NewConfig("config.env")

	// for create table and fill test-data
	if conn, err := db.ConnectDatabase(cfg); err == nil {
		db.MigrationReWriteTableWallets(conn)
	}

	router := v1.SetupRouter(cfg)
	log.Fatal(router.Run())
}
