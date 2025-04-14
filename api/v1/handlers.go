package v1

import (
	"net/http"
	"wallet-app/errors_app"
	"wallet-app/models"
	"wallet-app/services"

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
	return func(c *gin.Context) {
		var operation models.OperationWallet
		var wallet *models.Wallet
		var err error

		if err = c.BindJSON(&operation); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors_app.ErrorParsingOperationWalletJSON))
			return
		}
		if wallet, err = walletHandler.walletService.GetWallet(operation.WalletId); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors_app.ErrorGetWalletFromDatabase))
			return
		}
		if err = wallet.SetOperationBalance(operation); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(err))
			return
		}
		if err = walletHandler.walletService.UpdateWalletBalance(wallet); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors_app.ErrorUpdateWalletBalanceFromDatabase))
			return
		}
		c.JSON(http.StatusOK, models.NewResponseData(wallet))
	}
}

func (walletHandler *WalletHandler) GetWallets() gin.HandlerFunc {
	return func(c *gin.Context) {
		var wallets []*models.Wallet
		var err error

		if wallets, err = walletHandler.walletService.GetWalletAll(); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors_app.ErrorGetWalletFromDatabase))
			return
		}
		c.JSON(http.StatusOK, models.NewResponseData(wallets))
	}
}

func (walletHandler *WalletHandler) GetBalance() gin.HandlerFunc {
	return func(c *gin.Context) {
		var wallet *models.Wallet
		var id_string string
		var id uuid.UUID
		var err error

		id_string = c.Param("id")

		if id, err = uuid.Parse(id_string); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors_app.ErrorIdWallet))
			return
		}

		if wallet, err = walletHandler.walletService.GetWallet(id); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors_app.ErrorGetWalletFromDatabase))
			return
		}
		c.JSON(http.StatusOK, models.NewResponseData(wallet))
	}
}
