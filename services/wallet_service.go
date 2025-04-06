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

	rows, err := conn.Query(context.Background(), QueryGetWalletAll())

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
	err = conn.QueryRow(context.Background(), QueryGetWallet(), id).Scan(&wallet.ID, &wallet.Balance)
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
	_, err = conn.Exec(context.Background(), QueryCreateWallet(), wallet.ID, wallet.Balance)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func UpdateWalletBalance(cfg *config.Config, wallet *models.Wallet) error {
	conn, err := db.ConnectDatabase()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(context.Background(), QueryUpdateWalletBalance(), &wallet.Balance, &wallet.ID)
	return err
}
