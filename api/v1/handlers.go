package v1

import (
	"net/http"
	"wallet-app/config"
	"wallet-app/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandlePostAmount(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data JsonDataRequestPostWallet
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		wallet, err := services.GetWallet(cfg, data.WalletId)
		if data.OperationType == "DEPOSIT" {
			wallet.Deposit(data.Amount)
		}
		if data.OperationType == "WITHDRAW" {
			wallet.Withdraw(data.Amount)
		}
		services.UpdateWallet(cfg, wallet)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, wallet)
	}
}

func HandlerGetWallets(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		wallets, err := services.GetWalletAll(cfg)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, wallets)
	}
}

func HandlerGetBalance(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		id_string := c.Param("id")
		id, err := uuid.Parse(id_string)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		wallet, err := services.GetWallet(cfg, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, wallet.Balance)
	}
}

// func postWallet() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		body, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			panic(err)
// 		}
// 		var post_data_json JsonDataRequestPostWallet

// 		err = json.Unmarshal(body, &post_data_json)
// 		if err != nil {
// 			w.WriteHeader(http.StatusNotFound)
// 			json.NewEncoder(w).Encode(NewErrorResponse(err))
// 			return
// 		}

// 		err = isValidatePostData(post_data_json)
// 		if err != nil {
// 			w.WriteHeader(http.StatusNotFound)
// 			json.NewEncoder(w).Encode(NewErrorResponse(err))
// 			return
// 		}

// 		wallet := Wallet{
// 			UUID: post_data_json.WalletID,
// 		}

// 		if post_data_json.OperationType == "DEPOSIT" {
// 			err = db.QueryRow("UPDATE wallets SET balance = balance + $1 WHERE wallets.uuid = $2 AND wallets.balance > 0 RETURNING (SELECT balance FROM wallets WHERE uuid = $2);", post_data_json.Amount, wallet.UUID).Scan(&wallet.Balance)
// 			if err != nil {
// 				fmt.Println(err)
// 				w.WriteHeader(http.StatusNotFound)
// 				json.NewEncoder(w).Encode(NewErrorResponse(err))
// 				return
// 			}
// 		}
// 		if post_data_json.OperationType == "WITHDRAW" {
// 			err = db.QueryRow("UPDATE wallets SET balance = balance - $1 WHERE wallets.uuid = $2 AND wallets.balance > $1 RETURNING (SELECT balance FROM wallets WHERE uuid = $2);", post_data_json.Amount, wallet.UUID).Scan(&wallet.Balance)
// 			if err != nil {
// 				fmt.Println(err)
// 				w.WriteHeader(http.StatusNotFound)
// 				json.NewEncoder(w).Encode(NewErrorResponse(err))
// 				return
// 			}
// 		}

// 		err = db.QueryRow("SELECT balance FROM wallets WHERE uuid = $1", wallet.UUID).Scan(&wallet.Balance)
// 		if err != nil {
// 			w.WriteHeader(http.StatusNotFound)
// 			json.NewEncoder(w).Encode(NewErrorResponse(ErrorNotFoundWallet))
// 			return
// 		}

// 		json.NewEncoder(w).Encode(wallet)
// 	}
// }
