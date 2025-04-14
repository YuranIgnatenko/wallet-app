package errors

import "errors"

var (
	ErrorParsingOperationWalletJSON      = errors.New("Ошибка разбора json, пример: {'walletId':uuid, 'amount':float64, 'operationType':'DEPOSIT' | 'WITHDRAW'}")
	ErrorGetWalletFromDatabase           = errors.New("Ошибка walletId, праверьте ID (uuid type)")
	ErrorValidateAmountValue             = errors.New("Ошибка amount, ожидается ( 0 < amount < 999999 )")
	ErrorOperationTypeValue              = errors.New("Ошибка operationType, ожидается ( 'DEPOSIT' | 'WITHDRAW' )")
	ErrorUpdateWalletBalanceFromDatabase = errors.New("Ошибка обновления баланса")
	ErrorIdWallet                        = errors.New("Ошибка id wallet, используйте корректный uuid")
)
