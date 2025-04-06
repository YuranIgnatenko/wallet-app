package services

import (
	"context"
	"wallet-app/config"
	"wallet-app/db"
	"wallet-app/models"

	"github.com/google/uuid"
)

func GetWalletAll() ([]models.Wallet, error) {
	conn, err := db.ConnectDatabase()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(context.Background(), "SELECT * FROM $1", config.AppConfig.DB_NAME)
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

func GetWallet(id uuid.UUID) (*models.Wallet, error) {
	conn, err := db.ConnectDatabase()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var wallet models.Wallet
	err = conn.QueryRow(context.Background(), "SELECT id, balance FROM $1 WHERE id = $2", config.AppConfig.DB_NAME, id).Scan(&wallet.ID, &wallet.Balance)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func CreateWallet() (*models.Wallet, error) {
	conn, err := db.ConnectDatabase()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	wallet := models.NewWallet()
	_, err = conn.Exec(context.Background(), "INSERT INTO $1 (id, balance) VALUES ($2, $3)", config.AppConfig.DB_NAME, wallet.ID, wallet.Balance)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func UpdateWallet(cfg *config.Config, wallet *models.Wallet) error {
	conn, err := db.ConnectDatabase()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(context.Background(), "UPDATE $1 SET balance = $2 WHERE id = $3", config.AppConfig.DB_NAME, wallet.Balance, wallet.ID)
	return err
}
