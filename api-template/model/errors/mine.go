package errors

import "net/http"

var (
	ErrWorkingNotFound = NewStringError("Working not found", http.StatusBadRequest)
)
