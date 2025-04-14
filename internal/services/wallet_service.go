package services

import (
	"context"
	"wallet-app/internal/utils"
	"wallet-app/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type WalletService struct {
	poolConnection *pgxpool.Pool
}

func NewWalletService(poolConnection *pgxpool.Pool) *WalletService {
	return &WalletService{
		poolConnection: poolConnection,
	}
}

func (walletService *WalletService) GetWalletAll() ([]*models.Wallet, error) {
	rows, err := walletService.poolConnection.Query(
		context.Background(), utils.ReadQuery("wallet_get_all.sql"))
	if err != nil {
		return nil, err
	}
	var wallets []*models.Wallet
	for rows.Next() {
		var wallet models.Wallet
		if err := rows.Scan(&wallet.ID, &wallet.Balance); err != nil {
			return nil, err
		}
		wallets = append(wallets, &wallet)
	}
	return wallets, nil
}

func (walletService *WalletService) GetWallet(id uuid.UUID) (*models.Wallet, error) {
	var wallet models.Wallet
	err := walletService.poolConnection.QueryRow(context.Background(), utils.ReadQuery("wallet_get_by_id.sql"), id).Scan(&wallet.ID, &wallet.Balance)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (walletService *WalletService) CreateWallet() (*models.Wallet, error) {
	wallet := models.NewWallet()
	_, err := walletService.poolConnection.Exec(context.Background(), utils.ReadQuery("wallet_create.sql"), wallet.ID, wallet.Balance)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (walletService *WalletService) CreateWalletFromData(wallet *models.Wallet) error {
	_, err := walletService.poolConnection.Exec(context.Background(), utils.ReadQuery("wallet_create.sql"), wallet.ID, wallet.Balance)
	return err
}

func (walletService *WalletService) UpdateWalletBalance(wallet *models.Wallet) error {
	_, err := walletService.poolConnection.Exec(
		context.Background(), utils.ReadQuery("internal/services/queries/wallet_update.sql"), &wallet.Balance, &wallet.ID)
	return err
}
