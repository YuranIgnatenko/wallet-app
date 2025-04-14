package models

import (
	"testing"
	"wallet-app/errors"

	"github.com/google/uuid"
)

func TestDeposit(t *testing.T) {
	wallet := NewWallet()
	wallet.Deposit(0.1)
	if wallet.Balance != 0.1 {
		t.Error("Error deposit func wallet")
	}
}

func TestWithdraw(t *testing.T) {
	wallet := NewWallet()
	wallet.Deposit(0.2)
	wallet.Withdraw(0.1)
	if wallet.Balance != 0.1 {
		t.Error("Error deposit func wallet")
	}
}

func TestSetOperationBalance(t *testing.T) {
	test_cases := []struct {
		operaionType    OperationWallet
		expectedBalance float64
	}{
		{OperationWallet{uuid.MustParse("6556baca-96f1-48a4-aff2-50508a2c8b11"), "DEPOSIT", 0.1}, 0.1},
		{OperationWallet{uuid.MustParse("6556baca-96f1-48a4-aff2-50508a2c8b11"), "WITHDRAW", 0.1}, -0.1},
		{OperationWallet{uuid.MustParse("6556baca-96f1-48a4-aff2-50508a2c8b11"), "errortype", 999}, 0.0},
	}
	for _, tc := range test_cases {
		wallet := &Wallet{
			ID:      tc.operaionType.WalletId,
			Balance: 0,
		}

		wallet.SetOperationBalance(tc.operaionType)
		if tc.expectedBalance != wallet.Balance {
			t.Error("Error: expected balance != wallet.Balance", tc.expectedBalance, wallet.Balance, tc.operaionType.OperationType)
		}
	}

}

func TestValidateAmountValue(t *testing.T) {
	test_cases := []struct {
		amount float64
		result error
	}{
		{1, nil},
		{0, errors.ErrorValidateAmountValue},
		{-999, errors.ErrorValidateAmountValue},
	}
	wallet := NewWallet()

	for _, tc := range test_cases {
		err := wallet.ValidateAmountValue(tc.amount)
		if err != tc.result {
			t.Error(err)
		}
	}
}
