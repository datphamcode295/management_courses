package errors

import "net/http"

var (
	ErrUnexpectedHeader = NewStringError("Unexpected headers", http.StatusUnauthorized)
	ErrInvalidChallenge = NewStringError("Header Challenge is invalid", http.StatusUnauthorized)
	ErrInvalidSignature = NewStringError("Header Signature is invalid", http.StatusUnauthorized)
	ErrExpiredChallenge = NewStringError("Challenge is expired", http.StatusForbidden)
	ErrInvalidAddress   = NewStringError("Invalid address", http.StatusBadRequest)
)
