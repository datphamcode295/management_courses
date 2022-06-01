package errors

import "net/http"

var (
	ErrInvalidTxHash     = NewStringError("Invalid Tx hash", http.StatusUnauthorized)
	ErrNotOwnerTx        = NewStringError("Invalid owner of transaction", http.StatusUnauthorized)
	ErrInvalidBuyType    = NewStringError("Invalid Buy type", http.StatusBadRequest)
	ErrInvalidMethodName = NewStringError("Invalid method name", http.StatusUnauthorized)
	ErrInvalidContract   = NewStringError("Invalid contract", http.StatusUnauthorized)
	ErrInvalidInputData  = NewStringError("Invalid input data", http.StatusUnauthorized)
	ErrInvalidReceiver   = NewStringError("Wrong receiver address", http.StatusBadRequest)
	ErrInvalidNftAddress = NewStringError("Invalid nft address", http.StatusBadRequest)
)
