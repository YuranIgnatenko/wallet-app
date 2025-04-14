package services

import (
	"testing"
	"wallet-app/config"
	"wallet-app/db"
	"wallet-app/models"

	"github.com/google/uuid"
)

func TestGetWalletAll(t *testing.T) {
	if err := config.LoadConfig("../config.env"); err != nil {
		t.Error(err)
	}
	pool, err := db.ConnectDatabase()
	if err != nil {
		t.Error(err)
	}
	walletService := NewWalletService(pool)
	_, err = walletService.GetWalletAll()
	if err != nil {
		t.Error(err)
	}
}

func TestCreateWallet(t *testing.T) {
	if err := config.LoadConfig("../config.env"); err != nil {
		t.Error(err)
	}
	pool, err := db.ConnectDatabase()
	if err != nil {
		t.Error(err)
	}
	walletService := NewWalletService(pool)
	_, err = walletService.CreateWallet()
	if err != nil {
		t.Error(err)
	}
}

func TestGetWallet(t *testing.T) {
	if err := config.LoadConfig("../config.env"); err != nil {
		t.Error(err)
	}
	pool, err := db.ConnectDatabase()
	if err != nil {
		t.Error(err)
	}
	walletService := NewWalletService(pool)
	wallet, err := walletService.CreateWallet()
	if err != nil {
		t.Error(err)
	}
	_, err = walletService.GetWallet(wallet.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateWalletFromData(t *testing.T) {
	if err := config.LoadConfig("../config.env"); err != nil {
		t.Error(err)
	}
	id := uuid.MustParse("6556baca-96f1-48a4-aff2-50508a2c8b01")
	pool, err := db.ConnectDatabase()
	if err != nil {
		t.Error(err)
	}
	walletService := NewWalletService(pool)
	err = walletService.CreateWalletFromData(&models.Wallet{ID: id, Balance: 0})
	if err != nil {
		t.Error(err)
	}
	_, err = walletService.GetWallet(id)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateWalletBalance(t *testing.T) {
	if err := config.LoadConfig("../config.env"); err != nil {
		t.Error(err)
	}
	id := uuid.MustParse("6556baca-96f1-48a4-aff2-50508a2c8b00")
	wallet := &models.Wallet{ID: id, Balance: 0}
	pool, err := db.ConnectDatabase()
	if err != nil {
		t.Error(err)
	}
	walletService := NewWalletService(pool)
	err = walletService.CreateWalletFromData(wallet)
	if err != nil {
		t.Error(err)
	}
	wallet.Deposit(0.1)
	walletService.UpdateWalletBalance(wallet)
	wallet_updated, err := walletService.GetWallet(id)
	if err != nil {
		t.Error(err)
	}
	if wallet_updated.Balance != 0.1 {
		t.Error("Error updating balance on wallet")
	}
}
