package db

import (
	"context"
	"wallet-app/config"

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
	return pool, nil
}

func MigrationReWriteTableWallets(pool *pgxpool.Pool) error {
	var err error
	ctx := context.Background()
	if _, err = pool.Exec(ctx, "DROP TABLE $1;", config.AppConfig.DB_NAME); err != nil {
		return err
	}
	if _, err = pool.Exec(ctx, "CREATE TABLE $1 (id UUID NOT NULL, balance NUMERIC(10,2) NOT NULL);", config.AppConfig.DB_NAME); err != nil {
		return err
	}
	if _, err = pool.Exec(ctx, "INSERT INTO $1 (id, balance) VALUES (gen_random_uuid (),50.215), (gen_random_uuid (),221.214);", config.AppConfig.DB_NAME); err != nil {
		return err
	}
	return nil
}
