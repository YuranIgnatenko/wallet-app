package v1

import (
	"wallet-app/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/wallets/:id", HandlerGetBalance(cfg))
	router.GET("/api/v1/wallets", HandlerGetWallets(cfg))
	router.POST("/api/v1/wallet", HandlePostAmount(cfg))
	return router
}
