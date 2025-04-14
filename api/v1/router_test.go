package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"wallet-app/config"
	"wallet-app/db"
	"wallet-app/models"
	"wallet-app/services"

	"github.com/google/uuid"
)

var _ = config.LoadConfig("../../config.env")

func TestRouteGetBalance(t *testing.T) {

	testCases := []struct {
		wallet             *models.Wallet
		method             string
		path               string
		expectedStatusCode int
	}{
		{&models.Wallet{
			ID:      uuid.MustParse("6b030132-0f2f-44c9-9759-931a63486dd6"),
			Balance: 0,
		}, "GET", "/api/v1/wallets/", 200},

		{&models.Wallet{
			ID:      uuid.MustParse("6b030132-0f2f-44c9-9759-931a00000000"),
			Balance: 0,
		}, "GET", "/api/v1/wallets/", 200},
	}
	pool, err := db.ConnectDatabase()
	if err != nil {
		t.Error(err)
	}
	router := SetupRouter(pool)
	walletService := services.NewWalletService(pool)

	for _, tcase := range testCases {
		err := walletService.CreateWalletFromData(tcase.wallet)
		if err != nil {
			t.Error(err)
		}

		r, _ := http.NewRequest(tcase.method, tcase.path+tcase.wallet.ID.String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		if tcase.expectedStatusCode != w.Code {
			t.Errorf("Waiting: %v\n Result: %v", tcase.expectedStatusCode, w.Code)
		}
	}
}

func TestRouteGetWallets(t *testing.T) {
	pool, err := db.ConnectDatabase()
	if err != nil {
		t.Error(err)
	}
	router := SetupRouter(pool)
	walletService := services.NewWalletService(pool)

	r, err := http.NewRequest("GET", "/api/v1/wallets", nil)
	if err != nil {
		t.Error(err)
	}
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Error(err, w.Code)
	}

	_, err = walletService.GetWalletAll()
	if err != nil {
		t.Error(err)
	}

	if http.StatusOK != w.Code {
		t.Errorf("Waiting: %v\n Result: %v", http.StatusOK, w.Code)
	}
}

func TestRouteOperationWallet(t *testing.T) {
	pool, err := db.ConnectDatabase()
	if err != nil {
		t.Error(err)
	}
	router := SetupRouter(pool)
	walletService := services.NewWalletService(pool)

	tested_amount_deposit := 0.1
	tested_wallet, err := walletService.CreateWallet()
	if err != nil {
		t.Error(err)
	}

	request_data := models.OperationWallet{
		WalletId:      tested_wallet.ID,
		OperationType: models.OperationTypeDeposit,
		Amount:        tested_amount_deposit,
	}

	request_data_json, err := json.Marshal(request_data)
	if err != nil {
		t.Error(err)
	}

	r, err := http.NewRequest("POST", "/api/v1/wallet", bytes.NewBuffer(request_data_json))
	if err != nil {
		t.Error(err)
	}
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Error(err, w.Code)
	}

	// rewrite variable for
	// updating balance (get wallet from database)
	tested_wallet, err = walletService.GetWallet(tested_wallet.ID)
	if err != nil {
		t.Error(err)
	}

	if http.StatusOK != w.Code {
		t.Errorf("Waiting: %v\n Result: %v", http.StatusOK, w.Code)
	}

}
