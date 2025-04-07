package models

import (
	"wallet-app/errors_app"

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

func (wallet *Wallet) Deposit(amount float64) {
	wallet.Balance += amount
}

func (wallet *Wallet) Withdraw(amount float64) {
	wallet.Balance -= amount
}

func (wallet *Wallet) ValidateAmountValue(amount float64) error {
	if amount <= 0 {
		return errors_app.ErrorValidateAmountValue
	}
	return nil
}

func (wallet *Wallet) SetOperationBalance(operation_wallet OperationWallet) error {
	if err := wallet.ValidateAmountValue(operation_wallet.Amount); err != nil {
		return err
	}
	switch operation_wallet.OperationType {
	case OperationTypeDeposit:
		wallet.Deposit(operation_wallet.Amount)
	case OperationTypeWithdraw:
		wallet.Withdraw(operation_wallet.Amount)
	default:
		return errors_app.ErrorOperationTypeValue
	}
	return nil
}
