package v1

import (
	"fmt"
	"net/http"
	"wallet-app/config"
	"wallet-app/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandlePostAmount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data JsonDataRequestPostWallet
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		wallet, err := services.GetWallet(data.WalletId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		fmt.Println(wallet, wallet.Balance, data.Amount)
		if data.OperationType == "DEPOSIT" {
			wallet.Deposit(data.Amount)
		}
		if data.OperationType == "WITHDRAW" {
			wallet.Withdraw(data.Amount)
		}
		services.UpdateWalletBalance(config.AppConfig, wallet)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, wallet)
	}
}

func HandlerGetWallets() gin.HandlerFunc {
	return func(c *gin.Context) {
		wallets, err := services.GetWalletAll()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, wallets)
	}
}

func HandlerGetBalance() gin.HandlerFunc {
	return func(c *gin.Context) {
		id_string := c.Param("id")
		id, err := uuid.Parse(id_string)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		wallet, err := services.GetWallet(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, wallet.Balance)
	}
}
