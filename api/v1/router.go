package v1

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/wallets/:id", HandlerGetBalance())
	router.GET("/api/v1/wallets", HandlerGetWallets())
	router.POST("/api/v1/wallet", HandlePostAmount())
	return router
}
