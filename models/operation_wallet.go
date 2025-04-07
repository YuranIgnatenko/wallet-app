package models

import "github.com/google/uuid"

const OperationTypeDeposit = "DEPOSIT"
const OperationTypeWithdraw = "WITHDRAW"

type OperationWallet struct {
	WalletId      uuid.UUID `json:"walletId"`
	OperationType string    `json:"operationType"`
	Amount        float64   `json:"amount"`
}
