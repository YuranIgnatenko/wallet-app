package db

import (
	"context"
	"fmt"
	"wallet-app/config"

	"github.com/jackc/pgx/v4/pgxpool"
)


func ConnectDatabase(cfg *config.Config) (*pgxpool.Pool, error) {
	database_url := cfg.DATABASE_URL
	fmt.Println(database_url)
	
	pool, err := pgxpool.Connect(context.Background(), database_url)

	if err != nil {
		return nil, err
	}
	if err = pool.Ping(context.Background()); err != nil{
		return nil, err
	}
	// defer pool.Close()
	return pool, nil
}

func  MigrationReWriteTableWallets(pool *pgxpool.Pool) error {
	ctx := context.Background()
	var err error
	_, err = pool.Exec(ctx, "DROP TABLE wallets;")
	if err != nil {
		return err
	}
	_, err = pool.Exec(ctx, "CREATE TABLE wallets (id UUID NOT NULL, balance NUMERIC(10,2) NOT NULL);")
	if err != nil {
		return err
	}
	_, err = pool.Exec(ctx, "INSERT INTO wallets (id, balance) VALUES (gen_random_uuid (),50.215), (gen_random_uuid (),221.214);")
	if err != nil {
		return err
	}

	// _, err = db.Exec("UPDATE wallets SET balance = balance + 999 WHERE wallets.uuid = '25e65230-a3ef-4b8e-a7e4-788379548fc6' AND wallets.balance > 0;")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return nil
}
