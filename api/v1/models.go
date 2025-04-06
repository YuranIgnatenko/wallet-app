package v1

import "github.com/google/uuid"

type JsonDataRequestPostWallet struct {
	WalletId      uuid.UUID `json:"walletId"`
	OperationType string    `json:"operationType"`
	Amount        float64   `json:"amount"`
}
