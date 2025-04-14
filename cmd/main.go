package main

import (
	"wallet-app/config"
	"wallet-app/db"
	v1 "wallet-app/internal/api/v1"
	"wallet-app/internal/logger"
)

func main() {
	if err := config.LoadConfig("config.env"); err != nil {
		panic(err)
	}

	logger.InitLogger()
	logger.Log.Info("Starting wallet-app")

	poolConnection, err := db.ConnectDatabase()
	if err != nil {
		logger.Log.Fatalf("Error connecting db, err:%s", err.Error())
		panic(err)
	}
	defer poolConnection.Close()

	err = db.MigrationReWriteTableWallets(poolConnection)
	if err != nil {
		logger.Log.Fatalf("Error migration db, err:%s", err.Error())
	}

	router := v1.SetupRouter(poolConnection)
	logger.Log.Fatal(router.Run(config.AppConfig.SERVER_URL))
}
