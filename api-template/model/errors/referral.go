package errors

import "net/http"

var (
	ErrNotRewardClaim = NewStringError("Do not have reward to claim", http.StatusBadRequest)
	ErrInvalidAmount  = NewStringError("Invalid amount", http.StatusBadRequest)
	ErrInvalidClaimId = NewStringError("Invalid claim ID", http.StatusBadRequest)
)
