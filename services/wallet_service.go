package services

import (
	"context"
	"wallet-app/config"
	"wallet-app/db"
	"wallet-app/models"

	"github.com/google/uuid"
)

func GetWalletAll(cfg *config.Config) ([]models.Wallet, error) {
	conn, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(context.Background(), "SELECT * FROM wallets")
	if err != nil {
		return nil, err
	}
	var wallets []models.Wallet
	for rows.Next() {
		var wallet models.Wallet
		if err := rows.Scan(&wallet.ID, &wallet.Balance); err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet)
	}
	return wallets, nil
}

func GetWallet(cfg *config.Config, id uuid.UUID) (*models.Wallet, error) {
	conn, err := db.ConnectDatabase(cfg)
	defer conn.Close()
	if err != nil {
		return nil, err
	}
	var wallet models.Wallet
	err = conn.QueryRow(context.Background(), "SELECT id, balance FROM wallets WHERE id = $1", id).Scan(&wallet.ID, &wallet.Balance)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func CreateWallet(cfg *config.Config) (*models.Wallet, error) {
	conn, err := db.ConnectDatabase(cfg)
	defer conn.Close()
	if err != nil {
		return nil, err
	}
	wallet := models.NewWallet()
	_, err = conn.Exec(context.Background(), "INSERT INTO wallets (id, balance) VALUES ($1, $2)", wallet.ID, wallet.Balance)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func UpdateWallet(cfg *config.Config, wallet *models.Wallet) error {
	conn, err := db.ConnectDatabase(cfg)
	defer conn.Close()
	if err != nil {
		return err
	}
	_, err = conn.Exec(context.Background(), "UPDATE wallets SET balance = $1 WHERE id = $2", wallet.Balance, wallet.ID)
	return err
}
