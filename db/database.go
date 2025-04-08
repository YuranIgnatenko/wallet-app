package db

import (
	"context"
	"fmt"
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
	var query_drop_table = fmt.Sprintf("DROP TABLE %s", config.AppConfig.DB_NAME)
	var query_create_table = fmt.Sprintf("CREATE TABLE %s (id UUID NOT NULL, balance NUMERIC(10,2) NOT NULL);", config.AppConfig.DB_NAME)
	var query_insert_to_table = fmt.Sprintf("INSERT INTO %s (id, balance) VALUES (gen_random_uuid (),50.215), (gen_random_uuid (),221.214);", config.AppConfig.DB_NAME)

	var err error
	ctx := context.Background()

	pool.Exec(ctx, query_drop_table)

	if _, err = pool.Exec(ctx, query_create_table); err != nil {
		return err
	}
	if _, err = pool.Exec(ctx, query_insert_to_table); err != nil {
		return err
	}
	return nil
}
