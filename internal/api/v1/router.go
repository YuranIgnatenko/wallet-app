package v1

import (
	"wallet-app/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SetupRouter(poolConnection *pgxpool.Pool) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.MiddlewareContentType)
	walletHandler := NewWalletHandler(poolConnection)
	router.GET("/api/v1/wallets/:id", walletHandler.GetBalance())
	router.GET("/api/v1/wallets", walletHandler.GetWallets())
	router.POST("/api/v1/wallet", walletHandler.OperationWallet())
	return router
}
