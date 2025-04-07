package v1

import (
	"net/http"
	"wallet-app/errors_app"
	"wallet-app/models"
	"wallet-app/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandlerOperationWallet() gin.HandlerFunc {
	return func(c *gin.Context) {
		var operation models.OperationWallet
		var wallet *models.Wallet
		var err error

		if err = c.BindJSON(&operation); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors_app.ErrorParsingOperationWalletJSON))
			return
		}
		if wallet, err = services.GetWallet(operation.WalletId); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors_app.ErrorGetWalletFromDatabase))
			return
		}
		if err = wallet.SetOperationBalance(operation); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(err))
			return
		}
		if err = services.UpdateWalletBalance(wallet); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors_app.ErrorUpdateWalletBalanceFromDatabase))
			return
		}
		c.JSON(http.StatusOK, models.NewResponseData(wallet))
	}
}

func HandlerGetWallets() gin.HandlerFunc {
	return func(c *gin.Context) {
		var wallets []models.Wallet
		var err error

		if wallets, err = services.GetWalletAll(); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors_app.ErrorGetWalletFromDatabase))
			return
		}
		c.JSON(http.StatusOK, models.NewResponseData(wallets))
	}
}

func HandlerGetBalance() gin.HandlerFunc {
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

		if wallet, err = services.GetWallet(id); err != nil {
			c.JSON(http.StatusNotFound, models.NewResponseError(errors_app.ErrorGetWalletFromDatabase))
			return
		}
		c.JSON(http.StatusOK, models.NewResponseData(wallet))
	}
}
