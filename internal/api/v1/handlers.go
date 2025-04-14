package v1

import (
	"net/http"
	"wallet-app/internal/errors"
	"wallet-app/internal/logger"
	"wallet-app/internal/services"
	"wallet-app/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type WalletHandler struct {
	walletService *services.WalletService
}

func NewWalletHandler(poolConnection *pgxpool.Pool) *WalletHandler {
	return &WalletHandler{
		walletService: services.NewWalletService(poolConnection),
	}
}

func (walletHandler *WalletHandler) OperationWallet() gin.HandlerFunc {
	logger.Log.Info("using handler OperationWallet")
	return func(c *gin.Context) {
		var operation models.OperationWallet
		var wallet *models.Wallet
		var err error

		if err = c.BindJSON(&operation); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors.ErrorParsingOperationWalletJSON))
			logger.Log.Fatalf("Error convert json to struct, err:%s", err.Error())
			return
		}
		if wallet, err = walletHandler.walletService.GetWallet(operation.WalletId); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors.ErrorGetWalletFromDatabase))
			logger.Log.Fatalf("Error get wallet, err:%s, id=%v", err.Error(), operation.WalletId)
			return
		}
		if err = wallet.SetOperationBalance(operation); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(err))
			logger.Log.Fatalf("Error set operation balance, err:%s, operation=%#v", err.Error(), operation)
			return
		}
		if err = walletHandler.walletService.UpdateWalletBalance(wallet); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors.ErrorUpdateWalletBalanceFromDatabase))
			logger.Log.Fatalf("Error updating wallet balance, err:%s", err.Error())
			return
		}
		c.JSON(http.StatusOK, models.NewResponseData(wallet))
	}
}

func (walletHandler *WalletHandler) GetWallets() gin.HandlerFunc {
	logger.Log.Info("using handler GetWallets")
	return func(c *gin.Context) {
		var wallets []*models.Wallet
		var err error

		if wallets, err = walletHandler.walletService.GetWalletAll(); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(err))
			logger.Log.Fatalf("Error get wallets all, err:%s", err.Error())
			return
		}
		c.JSON(http.StatusOK, models.NewResponseData(wallets))
	}
}

func (walletHandler *WalletHandler) GetBalance() gin.HandlerFunc {
	logger.Log.Info("using handler GetBalance")
	return func(c *gin.Context) {
		var wallet *models.Wallet
		var id_string string
		var id uuid.UUID
		var err error

		id_string = c.Param("id")

		if id, err = uuid.Parse(id_string); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors.ErrorIdWallet))
			logger.Log.Fatalf("Error uuid parsing, err:%s, id=%s", err.Error(), id_string)
			return
		}

		if wallet, err = walletHandler.walletService.GetWallet(id); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors.ErrorGetWalletFromDatabase))
			logger.Log.Fatalf("Errorget wallet, err:%s, id=%s", err.Error(), id_string)
			return
		}
		c.JSON(http.StatusOK, models.NewResponseData(wallet))
	}
}
