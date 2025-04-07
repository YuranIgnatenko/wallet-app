// ./curl_deposit_wallet.bat {walletId} {amount} {operationType}
// ./curl_deposit_wallet.bat eb8a9ae5-4df0-4682-929a-01a9c529c7de 190.30 DEPOSIT

curl -X POST http://127.0.0.1:8080/api/v1/wallet --data '{"walletId":"%1", "amount":%2, "operationType":"%3"}' -H "Content-Type: application/json" 