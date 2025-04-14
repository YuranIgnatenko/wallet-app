package db

import (
	"context"
	"wallet-app/config"
	"wallet-app/internal/logger"
	"wallet-app/internal/utils"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectDatabase() (*pgxpool.Pool, error) {
	database_url := config.AppConfig.DATABASE_URL
	pool, err := pgxpool.Connect(context.Background(), database_url)

	if err != nil {
		return nil, err
	}
	if err = pool.Ping(context.Background()); err != nil {
		return nil, err
	}
	logger.Log.Infof("Creating connection database: url=%v", config.AppConfig.DATABASE_URL)
	return pool, nil
}

func MigrationReWriteTableWallets(pool *pgxpool.Pool) error {
	query := utils.ReadQuery("table_drop_create.sql")
	_, err := pool.Exec(context.Background(), query)
	if err != nil {
		logger.Log.Fatalf("Error migration db, err:%s", err.Error())
	}
	logger.Log.Info("Completed re-creating table (DROP, CREATE, INSERT)")
	return nil
}
