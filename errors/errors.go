package errors

import "errors"

var (
	ErrorAmountValue        = errors.New("Incorrected amount value")
	ErrorNotFoundWallet     = errors.New("Not found wallet (incorrected uuid)")
	ErrorValueAmount        = errors.New("Error Value amount (0 < amount < 99999999)")
	ErrorValueOperationType = errors.New("Error Value operationType (DEPOSIT | WITHDRAW)")
	ErrorSyntaxPost         = errors.New("Error body request")
)

type ErrorResponse struct {
	Detail string
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{Detail: err.Error()}
}
