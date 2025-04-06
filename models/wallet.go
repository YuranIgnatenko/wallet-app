package models

import (
	"errors"

	"github.com/google/uuid"
)

type Wallet struct {
	ID      uuid.UUID `json:"id"`
	Balance float64   `json:"balance"`
}

func NewWallet() *Wallet {
	return &Wallet{
		ID:      uuid.New(),
		Balance: 0,
	}
}

func (wallet *Wallet) Deposit(amount float64) error {
	if amount > 0 {
		return errors.New("amoubnrt")
	}
	wallet.Balance += amount
	return nil
}

func (wallet *Wallet) Withdraw(amount float64) error {
	if wallet.Balance < amount {
		return errors.New("amoubnrt")
	}
	wallet.Balance -= amount
	return nil
}
