package services

import (
	"fmt"
	"wallet-app/config"
)

// short func for readble
var f = fmt.Sprintf

// output string not contain params
func QueryGetWalletAll() string {
	return f("SELECT * FROM %s ;",
		config.AppConfig.DB_NAME)
}

// output string contain param
// $1 - id (uuid.UUID)
func QueryGetWallet() string {
	return f("SELECT * FROM %s WHERE id = $1 ;",
		config.AppConfig.DB_NAME)
}

// output string contain params
// $1 - id (uuid.UUID)
// $2 - balance (float64)
func QueryCreateWallet() string {
	return f("INSERT INTO %s (id, balance) VALUES ($1, $2) ;",
		config.AppConfig.DB_NAME)
}

// output string contain params
// $1 - balance (float64)
// $2 - id (uuid.UUID)
func QueryUpdateWalletBalance() string {
	return f("UPDATE %s SET balance = $1 WHERE id = $2 ;", config.AppConfig.DB_NAME)
}
